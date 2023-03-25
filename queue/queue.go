// Package queue decleres a work queue using RabbitMQ.
// The queue declaration is called by both client and server
// to manage tasks concurrently
package queue

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Error handler for queue functionalities
func HandleIfError(err error, msg string) {
	if err != nil {
		fmt.Printf("queue error:%s\napplication message:%s", err, msg)
		log.Panicf("%s: %s", msg, err)
	}
}

// Common initialization for server and client queue. The function
// returns a pointer to the connection and channel to be closed by
// either the client or the server.
// @Returns: (*amqp.Channel, *amqp.Connection, amqp.Queue)
func ConnectAndGetQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	//Connect locally to the RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	HandleIfError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	HandleIfError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"task_queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	HandleIfError(err, "Failed to declare a queue")

	return conn, ch, &q
}
