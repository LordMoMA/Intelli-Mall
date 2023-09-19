//go:build contract

package handlers

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/v2/message"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"

	"github.com/LordMoMA/Intelli-Mall/internal/am"
	"github.com/LordMoMA/Intelli-Mall/internal/ddd"
	"github.com/LordMoMA/Intelli-Mall/internal/registry"
	"github.com/LordMoMA/Intelli-Mall/internal/registry/serdes"
	"github.com/LordMoMA/Intelli-Mall/stores/internal/application"
	"github.com/LordMoMA/Intelli-Mall/stores/internal/application/commands"
	"github.com/LordMoMA/Intelli-Mall/stores/internal/domain"
	"github.com/LordMoMA/Intelli-Mall/stores/storespb"
)

var pactBrokerURL string
var pactUser string
var pactPass string
var pactToken string

func init() {
	getEnv := func(key, fallback string) string {
		if value, ok := os.LookupEnv(key); ok {
			return value
		}
		return fallback
	}

	pactBrokerURL = getEnv("PACT_URL", "http://127.0.0.1:9292")
	pactUser = getEnv("PACT_USER", "pactuser")
	pactPass = getEnv("PACT_PASS", "pactpass")
	pactToken = getEnv("PACT_TOKEN", "")
}

func TestStoresProducer(t *testing.T) {
	var err error

	stores := domain.NewFakeStoreRepository()
	products := domain.NewFakeProductRepository()
	mall := domain.NewFakeMallRepository()
	catalog := domain.NewFakeCatalogRepository()

	type rawEvent struct {
		Name    string
		Payload json.RawMessage
	}

	reg := registry.New()
	err = storespb.RegistrationsWithSerde(serdes.NewJsonSerde(reg))
	if err != nil {
		t.Fatal(err)
	}

	verifier := message.Verifier{}
	err = verifier.Verify(t, message.VerifyMessageRequest{
		VerifyRequest: provider.VerifyRequest{
			Provider:                   "stores-pub",
			ProviderVersion:            "1.0.0",
			BrokerURL:                  pactBrokerURL,
			BrokerUsername:             pactUser,
			BrokerPassword:             pactPass,
			BrokerToken:                pactToken,
			PublishVerificationResults: true,
			AfterEach: func() error {
				stores.Reset()
				products.Reset()
				return nil
			},
		},
		MessageHandlers: map[string]message.Handler{
			"a StoreCreated message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				// Assign
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				app := application.New(stores, products, catalog, mall, dispatcher)
				publisher := am.NewFakeEventPublisher()
				handler := NewDomainEventHandlers(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				// Act
				err := app.CreateStore(context.Background(), commands.CreateStore{
					ID:       "store-id",
					Name:     "NewStore",
					Location: "NewLocation",
				})
				if err != nil {
					return nil, nil, err
				}

				// Assert
				subject, event, err := publisher.Last()
				if err != nil {
					return nil, nil, err
				}

				return rawEvent{
						Name:    event.EventName(),
						Payload: reg.MustSerialize(event.EventName(), event.Payload()),
					}, map[string]any{
						"subject": subject,
					}, nil
			},
			"a StoreRebranded message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				app := application.New(stores, products, catalog, mall, dispatcher)
				publisher := am.NewFakeEventPublisher()
				handler := NewDomainEventHandlers(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				store := domain.NewStore("store-id")
				store.Name = "NewStore"
				store.Location = "NewLocation"
				stores.Reset(store)

				err := app.RebrandStore(context.Background(), commands.RebrandStore{
					ID:   "store-id",
					Name: "RebrandedStore",
				})
				if err != nil {
					return nil, nil, err
				}

				subject, event, err := publisher.Last()
				if err != nil {
					return nil, nil, err
				}

				return rawEvent{
						Name:    event.EventName(),
						Payload: reg.MustSerialize(event.EventName(), event.Payload()),
					}, map[string]any{
						"subject": subject,
					}, nil
			},
		},
	})

	if err != nil {
		t.Error(err)
	}
}
