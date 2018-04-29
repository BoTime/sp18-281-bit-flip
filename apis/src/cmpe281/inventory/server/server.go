package server

import (
	"cmpe281/common/output"
	"cmpe281/inventory/handler"
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"cmpe281/common"
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
	rootRouter := mux.NewRouter()

	handlerContext := handler.RequestContext{
		Cassandra: srv.Cassandra,
	}

	// Health Check Handler
	rootRouter.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
		return
	})

	// Root API Handler
	rootRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Starbcuks Inventory API")
		return
	})

	// Set up API Handlers
	{
		rootRouter.HandleFunc("/stores", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				handlerContext.ListStores(w, r)
			default:
				output.WriteErrorMessage(w, http.StatusMethodNotAllowed, "Method Not Supported")
			}
		})
		rootRouter.HandleFunc("/stores/{store_id}", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				handlerContext.GetStore(w, r)
			default:
				output.WriteErrorMessage(w, http.StatusMethodNotAllowed, "Method Not Supported")
			}
		})
		rootRouter.HandleFunc("/stores/{store_id}/inventory", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				handlerContext.GetInventory(w, r)
			default:
				output.WriteErrorMessage(w, http.StatusMethodNotAllowed, "Method Not Supported")
			}
		})
		{
			allocRouter := rootRouter.PathPrefix("/stores/{store_id}/allocations").Subrouter()
			allocRouter.Use(common.AuthMiddleware(true))
			allocRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "GET":
					handlerContext.ListAllocations(w, r)
				case "POST":
					handlerContext.CreateAllocation(w, r)
				default:
					output.WriteErrorMessage(w, http.StatusMethodNotAllowed, "Method Not Supported")
				}
			})
			allocRouter.HandleFunc("/{allocation_id}", func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "GET":
					handlerContext.GetAllocation(w, r)
				case "POST":
					handlerContext.ConfirmAllocation(w, r)
				case "DELETE":
					handlerContext.ExpireAllocation(w, r)
				default:
					output.WriteErrorMessage(w, http.StatusMethodNotAllowed, "Method Not Supported")
				}
			})
		}
	}

	httpSrv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", config.Ip, config.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      rootRouter, // Pass our instance of gorilla/mux in.
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
