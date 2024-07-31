package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"war_ticket/internal/provider/db"
)

func main() {

	//driver := os.Getenv("DRIVER")
	//dsn := os.Getenv("DSN")
	driver := "postgres"
	dsn := "postgresql://root:root@localhost:5432/warticket?sslmode=disable"

	log.Println("DRIVER :: ", driver)
	log.Println("DSN :: ", dsn)

	dbConn := db.NewDB(driver, dsn)

	inject := initHandler(dbConn)

	//router := initRouter(eventHandler, ticketHandler, orderHandler, userRepository)

	router := initRouterGin(inject.geh, inject.gth, inject.goh, inject.ur)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
