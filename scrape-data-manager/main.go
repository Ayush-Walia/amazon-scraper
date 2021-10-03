package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ayush-Walia/amazon-scraper/scrape-data-manager/handlers"
	"github.com/Ayush-Walia/amazon-scraper/scrape-data-manager/utils"
	"github.com/gorilla/mux"
)

func main() {
	Router := mux.NewRouter()

	// Initialize routes
	Router.HandleFunc("/health_check", handlers.HandleHealthCheck).Methods("GET")
	Router.HandleFunc("/save_product", handlers.HandleSaveProduct).Methods("POST")

	// Get port from env or use 8080 as default value.
	port := utils.GetEnv("PORT", "8080")

	s := &http.Server{
		Addr:    ":" + port,
		Handler: Router,
	}

	log.Println("Scrap data manager listening on port " + port)

	// Starting the server
	go func() {
		err := s.ListenAndServe()
		utils.CheckError(err)
	}()

	// trap sigterm / interrupt and gracefully shutdown the server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-done
	log.Println("Server Stopped, Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := s.Shutdown(ctx)
	utils.CheckError(err)
}
