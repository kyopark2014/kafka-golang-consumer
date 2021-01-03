package server

import (
	"context"
	"fmt"
	"kafka-golang-consumer/internal/config"
	"time"

	"github.com/segmentio/kafka-go"
)

// KafkaConsumer is a list of service
type KafkaConsumer struct{}

// Init is to start KafkaConsumer Service
func (p *KafkaConsumer) Init(conf *config.AppConfig) error {

	return nil
}

// Start is to run Profile Server
func (p *KafkaConsumer) Start() error {

	fmt.Println("Okay...")

	startKafka()

	fmt.Println("Kafka has been started...")

	time.Sleep(10 * time.Minute)

	return nil
}

/*
// OnTerminate is to close the servcie
func (p *KafkaConsumer) OnTerminate() error {
	log.I("KafkaConsumer Service was terminated")

	// To-Do: add codes for error cases if requires
	return nil
} */

func startKafka() {
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "eytopic",
		GroupID:  "g1",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Println("Some error occured ", err)
			continue
		}

		fmt.Println("Message is:", string(m.Value))
	}

}
