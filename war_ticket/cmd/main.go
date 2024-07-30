package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"war_ticket/internal/provider/db"
)

func main() {

	driver := os.Getenv("DRIVER")
	dsn := os.Getenv("DSN")

	log.Println("DRIVER :: ", driver)
	log.Println("DSN :: ", dsn)

	db := db.NewDB(driver, dsn)

	eventHandler, ticketHandler, orderHandler, userRepository := initHandler(db)

	router := initRouter(eventHandler, ticketHandler, orderHandler, userRepository)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Waiting signal send to chan quit
	// Blocking channel
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Println("Server exiting")
}
