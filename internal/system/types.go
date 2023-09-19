package system

import (
	"context"
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/LordMoMA/Intelli-Mall/internal/config"
	"github.com/LordMoMA/Intelli-Mall/internal/waiter"
)

type Service interface {
	Config() config.AppConfig
	DB() *sql.DB
	JS() nats.JetStreamContext
	Mux() *chi.Mux
	RPC() *grpc.Server
	Waiter() waiter.Waiter
	Logger() zerolog.Logger
}

type Module interface {
	Startup(context.Context, Service) error
}
