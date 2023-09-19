//go:build contract

package handlers

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/pact-foundation/pact-go/v2/matchers"
	"github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/LordMoMA/Intelli-Mall/baskets/internal/domain"
	"github.com/LordMoMA/Intelli-Mall/internal/ddd"
	"github.com/LordMoMA/Intelli-Mall/internal/registry"
	"github.com/LordMoMA/Intelli-Mall/internal/registry/serdes"
	"github.com/LordMoMA/Intelli-Mall/stores/storespb"
)

type String = matchers.String
type Map = matchers.Map

var Like = matchers.Like

func TestStoresConsumer(t *testing.T) {
	type mocks struct {
		stores   *domain.MockStoreCacheRepository
		products *domain.MockProductCacheRepository
	}

	type rawEvent struct {
		Name    string
		Payload map[string]any
	}

	reg := registry.New()
	err := storespb.RegistrationsWithSerde(serdes.NewJsonSerde(reg))
	if err != nil {
		t.Fatal(err)
	}

	pact, err := v4.NewAsynchronousPact(v4.Config{
		Provider: "stores-pub",
		Consumer: "baskets-sub",
		PactDir:  "./pacts",
	})
	if err != nil {
		t.Fatal(err)
	}

	tests := map[string]struct {
		given    []models.ProviderState
		metadata map[string]string
		content  Map
		on       func(m mocks)
	}{
		"a StoreCreated message": {
			metadata: map[string]string{
				"subject": storespb.StoreAggregateChannel,
			},
			content: Map{
				"Name": String(storespb.StoreCreatedEvent),
				"Payload": Like(Map{
					"id":   String("store-id"),
					"name": String("NewStore"),
				}),
			},
			on: func(m mocks) {
				m.stores.On("Add", mock.Anything, "store-id", "NewStore").Return(nil)
			},
		},
		"a StoreRebranded message": {
			metadata: map[string]string{
				"subject": storespb.StoreAggregateChannel,
			},
			content: Map{
				"Name": String(storespb.StoreRebrandedEvent),
				"Payload": Like(Map{
					"id":   String("store-id"),
					"name": String("RebrandedStore"),
				}),
			},
			on: func(m mocks) {
				m.stores.On("Rename", mock.Anything, "store-id", "RebrandedStore").Return(nil)
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := mocks{
				stores:   domain.NewMockStoreCacheRepository(t),
				products: domain.NewMockProductCacheRepository(t),
			}
			if tc.on != nil {
				tc.on(m)
			}
			handlers := NewIntegrationEventHandlers(m.stores, m.products)
			msgConsumerFn := func(contents v4.MessageContents) error {
				event := contents.Content.(*rawEvent)

				data, err := json.Marshal(event.Payload)
				if err != nil {
					return err
				}
				payload := reg.MustDeserialize(event.Name, data)

				return handlers.HandleEvent(
					context.Background(),
					ddd.NewEvent(event.Name, payload),
				)
			}

			message := pact.AddAsynchronousMessage()
			for _, given := range tc.given {
				message = message.GivenWithParameter(given)
			}
			assert.NoError(t, message.
				ExpectsToReceive(name).
				WithMetadata(tc.metadata).
				WithJSONContent(tc.content).
				AsType(&rawEvent{}).
				ConsumedBy(msgConsumerFn).
				Verify(t),
			)
		})
	}
}
