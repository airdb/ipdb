package bootstrap

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"golang.org/x/exp/slog"
)

type serveDeps struct {
	fx.In

	Logger *slog.Logger
}

type Serve struct {
	deps *serveDeps

	mux    *chi.Mux
	server *http.Server
}

func NewRest(deps serveDeps) *Serve {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(render.SetContentType(render.ContentTypeHTML))

	return &Serve{deps: &deps, mux: mux}
}

func (s *Serve) Start() error {
	s.mux.Route("/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome to ipdb\n"))
		})
		r.Handle("/metrics", promhttp.Handler())
	})

	s.server = &http.Server{Addr: ":8080", Handler: s.mux}

	log.Println("Starting server on port ", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Serve) Stop() error {
	return s.server.Shutdown(context.TODO())
}
