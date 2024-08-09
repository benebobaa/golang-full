package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"orchestra-svc/internal/delivery/kafka"
	"orchestra-svc/pkg"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	db     *sql.DB
	gin    *gin.Engine
	config *pkg.Config
	msg    *kafka.MessageHandler
}

func NewApp(db *sql.DB, gin *gin.Engine, config *pkg.Config) *App {
	return &App{db: db, gin: gin, config: config}
}

func (a *App) Run() {

	if err := a.startService(); err != nil {
		log.Fatalf("Error starting service: %v", err)
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", a.config.Port),
		Handler: a.gin,
	}

	consumer, err := kafka.NewKafkaConsumer(
		[]string{a.config.KafkaBroker},
		a.config.GroupID,
		[]string{a.config.OrchestraTopic},
		a.msg,
	)
	defer consumer.Close()

	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}

	ctxCancel, cancel2 := context.WithCancel(context.Background())
	defer cancel2()

	readyChan := make(chan int)

	go func() {
		defer close(readyChan)
		if err := consumer.Consume(ctxCancel); err != nil {
			log.Fatalf("Error consuming Kafka messages: %v", err)
		}
		log.Println("Kafka consumer started")
		readyChan <- 1
	}()

	<-readyChan
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
