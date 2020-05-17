package main

import (
	"fmt"

	"github.com/apex/log"
	"github.com/streadway/amqp"
)

func main() {
	server()
}

func server() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello rabbitMQ"),
	}

	for {
		ch.Publish("", q.Name, false, false, msg)
	}
	// Args are... exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing
	// Empty string for exhange means default exchange
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	// Args are...message string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table
	failOnError(err, "Failed to declare a queue")

	return conn, ch, &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
