//go:build wireinject
// +build wireinject

package main

import (
	conf "github.com/Skijetler/alphinium/gateway/internal/config"
	"github.com/Skijetler/alphinium/gateway/internal/interceptors"
	"github.com/Skijetler/alphinium/gateway/internal/registers"
)

func wireApp(*conf.Config, log.Logger) (*http.Server, func(), error) {
	panic(wire.Build(
		interceptors.ProviderSet,
		registers.ProviderSet,
		customHandler,
		newGatewayServer),
	)
}
