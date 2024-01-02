package main

import (
	"context"
	"golang-restful-api-technical-test/internal/config"
	"golang-restful-api-technical-test/internal/delivery/messaging"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	logger.Info("Starting worker service")

	ctx, cancel := context.WithCancel(context.Background())

	logger.Info("Setup user consumer")
	userConsumer := config.NewKafkaConsumer(viperConfig, logger)
	userHandler := messaging.NewUserConsumer(logger)
	go messaging.ConsumeTopic(ctx, userConsumer, "users", logger, userHandler.Consume)

	logger.Info("setup credit card consumer")
	creditcardConsumer := config.NewKafkaConsumer(viperConfig, logger)
	creditcardHandler := messaging.NewCreditcardConsumer(logger)
	go messaging.ConsumeTopic(ctx, creditcardConsumer, "creditcards", logger, creditcardHandler.Consume)

	logger.Info("worker is running")

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case s := <-terminateSignals:
			logger.Info("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s)
			cancel()
			stop = true
		}
	}

	time.Sleep(5 * time.Second) // wait for all consumers to finish processing
}
