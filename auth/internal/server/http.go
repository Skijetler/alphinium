package server

import (
	"github.com/Skijetler/alphinium/auth/internal/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewHTTPServer(conf *config.Config, log *logrus.Logger) *http.ServeMux {
	log.Info("New HTTP Server")
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	return mux
}
