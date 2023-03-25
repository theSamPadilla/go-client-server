// Common functions used by the server package
package server

import (
	"bytes"
	"client-server/common"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

// Gets a goroutine ID
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// Sleeps the goroutine if sleep is True
func sleepIfOn(sleep bool) {
	if sleep {
		common.INFO.Printf("Worker GID %d sleeping...\n\n", getGID())
		time.Sleep(time.Duration(rand.Intn(5-1)+1) * time.Second)
	}
}

// Acknowledges message to the queue for next dispatch, if sequential execution
func ackIfSequential(ack bool, r *amqp091.Delivery) {
	common.INFO.Printf("Goroutines running at the end of GID %d WORKER: %d\n\n", getGID(), runtime.NumGoroutine())
	if ack {
		r.Ack(false)
	}
}

// Logs Server initialization, including goroutines at runtime, queue status, and configs.
func logServerStart(nonSequential bool, sleep bool, q *amqp091.Queue) {
	common.INFO.Printf("STARTING SERVER\n")
	common.INFO.Printf("Number of goroutines on server start: %d (1 * Main, 2 * amqp server connection, 1 * consumer)\n", runtime.NumGoroutine())
	common.INFO.Printf("Sequential Execution: %v\n", !nonSequential)
	common.INFO.Printf("Sleep Simulation: %v\n", sleep)
	common.INFO.Printf("Queue status: %+v\n\n\n", q)

	//Alert stdout that server is running
	log.Printf("Server waiting for messages. Press CTRL+C to exit\n\n")
}
