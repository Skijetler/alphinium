package main

import (
	"flag"
	conf "github.com/Skijetler/alphinium/gateway/internal/config"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ochttp"
	"log"
	"net/http"
	"os"
)

var (
	// Name is the name of the compiled software.
	Name string = "Gateway"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	// prefixs is the config environment variable prefix
	prefixs = []string{"ALPHINIUM_GATEWAY_", "GATEWAY_"}

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newGatewayServer(conf *conf.Config, oc *ochttp.Handler) *http.Server {
	return &http.Server{
		Addr:    conf.Rest.Addr,
		Handler: oc,
	}
}

func customHandler(mux *runtime.ServeMux, logger *logrus.Logger) (http.Handler, error) {
	if err := mux.HandlePath(
		"GET",
		"/health",
		func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			if _, err := w.Write([]byte("Gateway v" + Version + " serving")); err != nil {
				log.NewHelper(logger).Errorf("Health check error: %v", err)
			}
		},
	); err != nil {
		return nil, err
	}

	return wsproxy.WebsocketProxy(mux), nil
}
