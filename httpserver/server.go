package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Config struct {
	HttpMetricAddr string
}

type HttpServer struct {
	Config Config
	Router *mux.Router
}

func New(config Config) *HttpServer {
	s := &HttpServer{
		Config: config,
		Router: mux.NewRouter(),
	}
	s.Router.Handle("/metrics", promhttp.Handler())
	return s
}

func (s *HttpServer) Start(errChan chan error) {
	errChan <- http.ListenAndServe(s.Config.HttpMetricAddr, s.Router)
}
