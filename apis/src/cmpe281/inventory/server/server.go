package server

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Cassandra *gocql.Session
}

type Config struct {
	Ip   string
	Port string
	Wait time.Duration
}

func (srv *Server) Run(config Config) {
	router := mux.NewRouter()

	// Health Check Handler
	router.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
		return
	})

	// Set up API Handlers
	// TODO(bbamsch): Implement API Handlers

	httpSrv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", config.Ip, config.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("HTTP Server binding to %s", httpSrv.Addr)
		if err := httpSrv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Block until SIGINT (Ctrl+C) is received, then begin shutdown.
	<-c

	log.Printf("HTTP Server beginning shutdown")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), config.Wait)
	defer cancel()

	// Wait until all connections are finished or until timeout
	httpSrv.Shutdown(ctx)

	log.Println("HTTP Server shut down successfully")
}
