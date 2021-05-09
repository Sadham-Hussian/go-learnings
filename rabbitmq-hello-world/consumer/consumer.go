package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func logError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT")
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")

func main() {
	rabbitMQConn, err := amqp.Dial("amqp://" + rabbit_user + ":" + rabbit_password + "@" + rabbit_host + ":" + rabbit_port + "/")
	logError(err, "Failed to connect to RabbitMQ")
	log.Println("Connection created")
	defer rabbitMQConn.Close()

	ch, err := rabbitMQConn.Channel()
	logError(err, "Failed to open a channel")
	log.Println("Channel opened")
	defer ch.Close()

	msgs, err := ch.Consume(
		"hello-world", // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	logError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
