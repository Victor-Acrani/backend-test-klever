package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	srv *http.Server
}

func NewServer(addr string) *Server {
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	return &Server{srv: server}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) Start() error {
	// create router
	r := mux.NewRouter()
	// set middlewares
	r.Use(
		middlewareRequestID,
		middlwareLogger,
		middlewareRecover,
	)

	// create subrouter
	apiV2Router := r.PathPrefix("/api/v2").Subrouter()

	// set routes
	apiV2Router.HandleFunc("", s.HandlerHealthCheck).Methods(http.MethodGet)
	apiV2Router.HandleFunc("/details/{address}", s.HandlerDetails).Methods(http.MethodGet)
	apiV2Router.HandleFunc("/balance/{address}", s.HanlderBalance).Methods(http.MethodGet)
	apiV2Router.HandleFunc("/send", s.HandlerSend).Methods(http.MethodPost)

	// define no route
	r.PathPrefix("/").HandlerFunc(s.HandlerNoRoute)

	// assing router to server
	s.srv.Handler = r

	go func() {
		log.Println("listenning and serving on port " + s.srv.Addr)
		err := s.srv.ListenAndServe()
		if err != nil {
			log.Println("ListenAndServe() ->", err.Error())
		}
	}()

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	// Create a deadline to wait for.
	const timeout = 15 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// shutdown server
	return s.srv.Shutdown(ctx)
}
