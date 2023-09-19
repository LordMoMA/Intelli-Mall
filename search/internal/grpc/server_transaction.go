package grpc

import (
	"context"
	"database/sql"

	"google.golang.org/grpc"

	"github.com/LordMoMA/Intelli-Mall/internal/di"
	"github.com/LordMoMA/Intelli-Mall/search/internal/application"
	"github.com/LordMoMA/Intelli-Mall/search/internal/constants"
	"github.com/LordMoMA/Intelli-Mall/search/searchpb"
)

type serverTx struct {
	c di.Container
	searchpb.UnimplementedSearchServiceServer
}

var _ searchpb.SearchServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	searchpb.RegisterSearchServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) SearchOrders(ctx context.Context, request *searchpb.SearchOrdersRequest) (resp *searchpb.SearchOrdersResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.Application)}

	return next.SearchOrders(ctx, request)
}

func (s serverTx) GetOrder(ctx context.Context, request *searchpb.GetOrderRequest) (resp *searchpb.GetOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.Application)}

	return next.GetOrder(ctx, request)
}

func (s serverTx) closeTx(tx *sql.Tx, err error) error {
	if p := recover(); p != nil {
		_ = tx.Rollback()
		panic(p)
	} else if err != nil {
		_ = tx.Rollback()
		return err
	} else {
		return tx.Commit()
	}
}
