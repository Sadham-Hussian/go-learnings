package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT")
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")

func logError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	rabbitmqConn, err := amqp.Dial("amqp://" + rabbit_user + ":" + rabbit_password + "@" + rabbit_host + ":" + rabbit_port + "/")
	if err != nil {
		logError(err, "Failed to create RabbitMQ connection")
	} else {
		log.Println("RabbitMQ connection created!!!")
	}
	defer rabbitmqConn.Close()

	app := fiber.New()
	app.Get("/publish", func(c *fiber.Ctx) error {
		msg := c.Query("msg")
		if msg == "" {
			log.Println("msg parameter missing!!")
			return c.SendStatus(500)
		}
		fmt.Println(rabbitmqConn)
		ch, err := rabbitmqConn.Channel()
		if err != nil {
			logError(err, "Failed to open a channel")
		} else {
			log.Println("Channel opened!!")
		}
		defer ch.Close()

		queue, err := ch.QueueDeclare(
			"hello-world",
			false,
			false,
			false,
			false,
			nil,
		)

		logError(err, "Failed to create a queue")
		log.Println("Queue created")

		err = ch.Publish(
			"",
			queue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			},
		)

		logError(err, "Failed to put message in queue")
		log.Println("Message added")

		return c.SendStatus(201)
	})

	log.Fatal(app.Listen(":5050"))
}
