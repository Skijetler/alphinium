package server

import (
	v1 "github.com/Skijetler/alphinium/auth/api/v1"
	"github.com/Skijetler/alphinium/auth/internal/config"
	"github.com/Skijetler/alphinium/auth/internal/service"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(conf *config.Config, auth *service.AuthService, metrics *grpc_prometheus.ServerMetrics, logger *logrus.Logger) *grpc.Server {
	logger.Info("New gRPC Server")
	log := logrus.NewEntry(logger)

	var opts = []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_validator.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(log),
			metrics.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(log),
			metrics.UnaryServerInterceptor(),
		)),
	}

	srv := grpc.NewServer(opts...)
	v1.RegisterAuthServer(srv, auth)
	grpc_prometheus.Register(srv)

	return srv
}
