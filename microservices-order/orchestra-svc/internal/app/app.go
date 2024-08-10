package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"orchestra-svc/internal/delivery/messaging"
	"orchestra-svc/pkg"
	"orchestra-svc/pkg/consumer"
	"orchestra-svc/pkg/producer"
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
	msg    *messaging.MessageHandler
}

func NewApp(db *sql.DB, gin *gin.Engine, config *pkg.Config) *App {
	return &App{db: db, gin: gin, config: config}
}

func (a *App) Run() {

	userProducer, err := producer.NewKafkaProducer(
		[]string{a.config.KafkaBroker},
		a.config.UserTopic,
	)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}
	defer userProducer.Close()

	productProducer, err := producer.NewKafkaProducer(
		[]string{a.config.KafkaBroker},
		a.config.ProductTopic,
	)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}
	defer productProducer.Close()

	if err := a.startService(userProducer); err != nil {
		log.Fatalf("Error starting service: %v", err)
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", a.config.Port),
		Handler: a.gin,
	}

	consumer, err := consumer.NewKafkaConsumer(
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

	go func() {
		if err := consumer.Consume(ctxCancel); err != nil {
			log.Fatalf("Error consuming Kafka messages: %v", err)
		}
	}()

	go func() {
		log.Printf("Starting server on port %s", a.config.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}

	log.Println("Server exiting")
}
