package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"server/controller"

	"github.com/gorilla/mux"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*120, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	fmt.Println("server started")
	router := mux.NewRouter()
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	router.HandleFunc("/animals", controller.GetAnimals).Methods(http.MethodGet)
	router.HandleFunc("/animals/{id}", controller.GetAnimal).Methods("GET").Name("test")
	router.HandleFunc("/animals", controller.CreateAnimal).Methods("POST")
	router.HandleFunc("/animals/{id}", controller.UpdateAnimal).Methods("PUT")
	router.HandleFunc("/animals/{id}", controller.DeleteAnimal).Methods("DELETE")
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	url, err := router.Get("test").URL("id", "23")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("sss")

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	time.Sleep(2 * time.Second)
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
