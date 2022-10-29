package main

import (
	"flag"
	"github.com/Skijetler/alphinium/auth/internal/app"
	"github.com/Skijetler/alphinium/auth/internal/config"
	"github.com/Skijetler/alphinium/auth/internal/pkg/logger"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "Auth"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	// prefixs is the config environment variable prefix
	prefixs = []string{"ALPHINIUM_AUTH_", "AUTH_"}
)

func newApp(conf *config.Config, logger *logrus.Logger, grpcSrv *grpc.Server, httpSrv *http.Server) *app.App {
	return app.New(
		conf,
		logger,
		grpcSrv,
		httpSrv,
	)
}

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	conf := config.GetConfig(flagconf)

	logger := logger.NewLogger(conf)
	metrics := grpc_prometheus.NewServerMetrics()

	app, cleanup, err := wireApp(conf, metrics, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
