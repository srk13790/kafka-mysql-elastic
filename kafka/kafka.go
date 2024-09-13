package kafka

import (
	"fmt"
	db "kafka-new/database"
	"kafka-new/elasticsearch"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/IBM/sarama"
)

func KafkaExec(data []byte) {
	broker := "localhost:9092"
	topic := "ganoderma"

	// Run the producer in a separate goroutine
	go produceMessages(broker, topic, string(data))

	// Run the consumer in the main goroutine
	consumeMessages(broker, topic)
}

// Producer function
func produceMessages(broker string, topic string, message string) {
	producer, err := newProducer(broker)
	if err != nil {
		log.Fatalf("Failed to start producer: %v", err)
	}

	// for i := 0; i < 10; i++ {
		// message := fmt.Sprintf("Message #%d", i)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}

		_, _, err = producer.SendMessage(msg)
		if err != nil {
			log.Printf("Failed to produce message: %v", err)
		} else {
			log.Printf("Produced message: %s", message)
		}
		time.Sleep(1 * time.Second)
	// }

	producer.Close()
}

// Consumer function
func consumeMessages(broker string, topic string) {
	consumer, err := sarama.NewConsumer([]string{broker}, nil)
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("Failed to get the partitions: %v", err)
	}

	// Set up a signal to listen for interruption (Ctrl+C) to gracefully shut down
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Failed to start consuming partition: %v", err)
		}

		defer pc.Close()

		go func(pc sarama.PartitionConsumer) {
			for {
				select {
				case msg := <-pc.Messages():
					fmt.Printf("Consumed message: %s\n", string(msg.Value))
					// db.InsertData(string(msg.Value)) // Insert Data into Redis DB

					// Connet to MySQL DB
					conn := db.Connect()
					db.InsertMySqlData(conn, string(msg.Value))

					// Add data to elastic
					elasticsearch.StartElastic(string(msg.Value))

				case <-signals:
					return
				}
			}
		}(pc)
	}

	// Wait for interrupt signal to gracefully shut down the consumer
	<-signals
	fmt.Println("Interrupt is detected, shutting down.")
	consumer.Close()
}

// Helper function to create a producer
func newProducer(broker string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	return producer, err
}
