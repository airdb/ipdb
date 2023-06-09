package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"ipdb/modules/ipipmod"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/ip2location/ip2location-go/v9"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
)

type serveDeps struct {
	fx.In

	// Logger *slog.Logger
	IP2LocationDB *ip2location.DB
	IPIPDB        *ipipmod.DB
}

type Serve struct {
	deps *serveDeps

	mux    *chi.Mux
	server *http.Server
}

func NewServe(deps serveDeps) *Serve {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(render.SetContentType(render.ContentTypeHTML))

	return &Serve{deps: &deps, mux: mux}
}

var helpMsg = `
==== Welcome to ipdb service ====

Usage: visit or curl the url directly, then you will get the ip information.

IP2Location: https://www.ip2location.com/
  GET /v1/ip2location/
  GET /v1/ip2location/{ip}

IPIP: https://www.ipip.net/
  GET /v1/ipip/
  GET /v1/ipip/{ip}

MaxMind: https://www.maxmind.com/
  GET /v1/maxmind/
  GET /v1/maxmind/{ip}

IPPlus360: https://www.ipplus360.com/
  GET /v1/ipplus360/
  GET /v1/ipplus360/{ip}


Thank you for following us!!
https://github.com/airdb
`

func (s *Serve) Stop() error {
	return s.server.Shutdown(context.TODO())
}

func (s *Serve) Start() error {
	s.mux.Route("/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(helpMsg))
		})
		r.Handle("/metrics", promhttp.Handler())
	})

	s.mux.Route("/v1/ipip/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
				return
			}

			log.Println("ping", ip)
			record, _ := s.deps.IPIPDB.DB.Find(ip, "CN")
			render.JSON(w, r, record)
		})
		r.HandleFunc("/{ip}", func(w http.ResponseWriter, r *http.Request) {
			ip := chi.URLParam(r, "ip")
			log.Println("ping", ip)
			record, _ := s.deps.IPIPDB.DB.Find(ip, "CN")

			render.JSON(w, r, record)
		})
	})

	s.mux.Route("/v1/ip2location/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
				return
			}

			log.Println("ping", ip)
			record, _ := s.deps.IP2LocationDB.Get_all(ip)
			render.JSON(w, r, record)
		})
		r.HandleFunc("/{ip}", func(w http.ResponseWriter, r *http.Request) {
			ip := chi.URLParam(r, "ip")
			log.Println("ping", ip)
			record, _ := s.deps.IP2LocationDB.Get_all(ip)

			render.JSON(w, r, record)
		})
	})

	s.mux.Route("/v1/maxmind/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
				return
			}

			log.Println("ping", ip)
			record, _ := s.deps.IP2LocationDB.Get_all(ip)
			render.JSON(w, r, record)
		})
		r.HandleFunc("/{ip}", func(w http.ResponseWriter, r *http.Request) {
			ip := chi.URLParam(r, "ip")
			log.Println("ping", ip)
			record, _ := s.deps.IP2LocationDB.Get_all(ip)

			render.JSON(w, r, record)
		})
	})

	s.mux.Route("/v1/ipplus360/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
				return
			}

			log.Println("ping", ip)
			record, _ := s.deps.IP2LocationDB.Get_all(ip)
			render.JSON(w, r, record)
		})
		r.HandleFunc("/{ip}", func(w http.ResponseWriter, r *http.Request) {
			ip := chi.URLParam(r, "ip")
			log.Println("ping", ip)
			record, _ := s.deps.IP2LocationDB.Get_all(ip)

			render.JSON(w, r, record)
		})
	})

	s.server = &http.Server{Addr: ":8080", Handler: s.mux}

	log.Println("Starting server on port ", s.server.Addr)
	return s.server.ListenAndServe()
}
