package main

import (
	"github.com/streadway/amqp"
	"log"
	"github.com/remfath/go-rabbitmq-demo/util"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.HandleErr("failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	util.HandleErr("failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare("test-queue", false, false, false, false, nil)
	util.HandleErr("failed to declare a queue", err)

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	util.HandleErr("failed to register a consumer", err)
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
}
