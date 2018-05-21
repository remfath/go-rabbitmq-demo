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
	body := "hi there"
	err = ch.Publish("", q.Name, false, false,
		amqp.Publishing{ContentType: "text/plain", Body: []byte(body)},
	)
	log.Printf("[v] Sent %s", body)
	util.HandleErr("failed to publish a message", err)
}
