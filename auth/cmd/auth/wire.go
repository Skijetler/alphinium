//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/Skijetler/alphinium/auth/internal/app"
	"github.com/Skijetler/alphinium/auth/internal/config"
	"github.com/Skijetler/alphinium/auth/internal/pkg/hash"
	"github.com/Skijetler/alphinium/auth/internal/pkg/paseto"
	"github.com/Skijetler/alphinium/auth/internal/repo"
	"github.com/Skijetler/alphinium/auth/internal/server"
	"github.com/Skijetler/alphinium/auth/internal/service"
	"github.com/Skijetler/alphinium/auth/internal/usecase"
	"github.com/google/wire"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sirupsen/logrus"
)

// wireApp init application.
func wireApp(*config.Config, *grpc_prometheus.ServerMetrics, *logrus.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		repo.ProviderSet,
		usecase.ProviderSet,
		service.ProviderSet,
		hash.ProviderSet,
		paseto.ProviderSet,
		newApp),
	)
}
