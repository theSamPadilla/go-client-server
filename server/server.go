// Package server implements a consumer of RabbitMQ messages and routes them to
// the appropriate handler according to the method sent by the client. Messages
// can add, remove, getItem, getIndex or getall itesm from an ordered map.
//
// This file contains the consumer definition and router definition. Handlers
// and utilities in handlers.go and utils.go respectively.
package server

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"

	"github.com/thesampadilla/go-client-server/common"
	"github.com/thesampadilla/go-client-server/orderedmap"
	"github.com/thesampadilla/go-client-server/queue"

	"github.com/rabbitmq/amqp091-go"
)

// Starts the server, blocks channel to listen to requests, and invokes router
func StartServer(om *orderedmap.OrderedMap, nonSequential bool, sleep bool) {
	//Defer wg
	//defer wg.Done()

	//Get q, conn and ch and defer their close
	conn, ch, q := queue.ConnectAndGetQueue()
	defer conn.Close()
	defer ch.Close()
	var err error

	//Dispatch messages without ack from workers
	if nonSequential {
		err = ch.Qos(0, 0, false)

	} else { //Dispatches messages to workers on ack, ensuring sequentiality
		err = ch.Qos(1, 0, false)
	}
	queue.HandleIfError(err, "Failed to set QoS")

	//Consume messages
	msgs, err := ch.Consume(
		q.Name,        // queue
		"",            // consumer
		nonSequential, // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	queue.HandleIfError(err, "Failed to register a consumer")

	//Log Server start, block channel and dispatch messages to router
	logServerStart(nonSequential, sleep, q)
	var repeat chan struct{}
	MessageRouter(&msgs, om, nonSequential, sleep)
	common.INFO.Printf("Goroutines running at the end of SERVER: %d\n\n", runtime.NumGoroutine())
	<-repeat
}

// Dispatches messages to workers according to the method type
// Each worker runs in its own goroutine.
// Runs indefinitely iterating over the message channel.
func MessageRouter(msgs *<-chan amqp091.Delivery, om *orderedmap.OrderedMap, nonSequential bool, sleep bool) {
	for request := range *msgs {
		//Reinitialize new empty Client and unmarshal the body into it
		var client common.Client
		err := json.Unmarshal(request.Body, &client)
		if err != nil {
			fmt.Printf("error unmarshaling the message: %s\n", err)
			log.Panic(err)
		}

		//Route request to new goroutine, pass data structure and configs.
		switch client.Method {
		case "add":
			go AddItem(client.Key, client.Value, om, &request, nonSequential, sleep)
		case "remove":
			go RemoveItem(client.Key, om, &request, nonSequential, sleep)
		case "get":
			go GetItem(client.Key, om, &request, nonSequential, sleep)
		case "geti":
			go GetItemByIndex(client.Key, om, &request, nonSequential, sleep)
		default:
			go GetAllItems(om, &request, nonSequential) //Get All never waits
		}
	}
	common.INFO.Printf("Goroutines running at the end of ROUTER: %d\n", runtime.NumGoroutine())
}
