package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"client-server/common"
	"client-server/queue"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	//Set command-liine flags
	client := getFlags()

	//Get queue, defer closing of conn and ch
	conn, ch, q := queue.ConnectAndGetQueue()
	defer conn.Close()
	defer ch.Close()
	fmt.Printf("Queue status: %+v\n", q)

	//Initialize context w 5s timout. Defer cancel
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	//Marshal client
	msg, err := json.Marshal(client)
	if err != nil {
		fmt.Printf("an error %s occured when marshalling the client input", err)
		return
	}

	//Publish message to channel
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         msg,
		})
	queue.HandleIfError(err, "Failed to publish message")

	//log.Printf(" [x] Sent %s", msg)
	fmt.Printf("Succesfully sent message %s to queue\n", msg)
}

// Sets and parses command line flags
// @Returns: *Client with the passed CLI commands
func getFlags() *common.Client {
	//Guards
	if len(os.Args) < 2 {
		fmt.Println("error: expected a method (add, remove, get, getall) subcommand")
		os.Exit(1)
	}
	if os.Args[1] != "add" && os.Args[1] != "remove" && os.Args[1] != "get" && os.Args[1] != "getall" {
		fmt.Println("error: the provided method is not supported.\nSupported methods are: add, remove, get, getall")
		os.Exit(1)
	}

	//Set flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addKey := addCmd.String("key", "", "key of element to add to map")
	addValue := addCmd.String("value", "", "value of element to add to map")

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeKey := removeCmd.String("key", "", "key of element to remove from map")

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getKey := getCmd.String("key", "", "key of element to get from map")

	//Parse values and initialize client
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		validateFlag("add", *addKey, "key")
		validateFlag("add", *addValue, "value")
		return clientConstructor("add", *addKey, *addValue)
	case "remove":
		removeCmd.Parse(os.Args[2:])
		validateFlag("remove", *removeKey, "key")
		return clientConstructor("remove", *removeKey, nil)
	case "get":
		getCmd.Parse(os.Args[2:])
		validateFlag("get", *getKey, "key")
		return clientConstructor("get", *getKey, nil)
	default:
		return clientConstructor("getall", nil, nil)
	}
}

// Validates that a flag exists, exists otherwise
func validateFlag(method string, flag string, def string) {
	if flag == "" {
		fmt.Printf("error: missing flag for %s. Please set the -%s flag\n", method, def)
		os.Exit(1)
	}
}

// Instantiates the Client struct from command line arguments
func clientConstructor(method string, key interface{}, value interface{}) *common.Client {
	return &common.Client{Method: method, Key: key, Value: value}
}
