package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/LordMoMA/Intelli-Mall/internal/rpc"
	"github.com/LordMoMA/Intelli-Mall/search/internal/application"
	"github.com/LordMoMA/Intelli-Mall/search/internal/models"
	"github.com/LordMoMA/Intelli-Mall/stores/storespb"
)

type StoreRepository struct {
	endpoint string
}

var _ application.StoreRepository = (*StoreRepository)(nil)

func NewStoreRepository(endpoint string) StoreRepository {
	return StoreRepository{
		endpoint: endpoint,
	}
}

func (r StoreRepository) Find(ctx context.Context, storeID string) (store *models.Store, err error) {
	var conn *grpc.ClientConn
	conn, err = r.dial(ctx)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := storespb.NewStoresServiceClient(conn).GetStore(ctx, &storespb.GetStoreRequest{Id: storeID})
	if err != nil {
		return nil, err
	}

	return r.storeToDomain(resp.Store), nil
}

func (r StoreRepository) storeToDomain(store *storespb.Store) *models.Store {
	return &models.Store{
		ID:   store.GetId(),
		Name: store.GetName(),
	}
}

func (r StoreRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return rpc.Dial(ctx, r.endpoint)
}
