// Package main is the entrypoint for module client-server. Main orchestrates usage
// settings to allow for either an http server to be spun up, or a client-worker
// orchestration using RabbitMQ.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/thesampadilla/go-client-server/common"
	"github.com/thesampadilla/go-client-server/httpserver"
	"github.com/thesampadilla/go-client-server/orderedmap"
	"github.com/thesampadilla/go-client-server/server"
)

func main() {
	//Get flags, set loggers, and set orderedmap
	serverType, nonSequential, sleep := getFlags()
	common.SetLoggers("logs/server.log", nonSequential, sleep)
	om := orderedmap.Constructor()

	//HTTP Server, no other flags are relevant.
	if serverType == "http" {
		r := httpserver.RegisterRoutes(om)
		fmt.Printf("--STARTING HTTP SERVER--")
		fmt.Printf("\n[*] Serving http responses on port 6969 and waiting for clients...\n\n")
		http.ListenAndServe(":6969", r)
	} else {
		server.StartServer(om, nonSequential, sleep)
	}
}

func getFlags() (string, bool, bool) {
	//Default Execution
	if len(os.Args) == 1 {
		return "tcp", false, false
	}

	//Parse and return arguments
	serverType := flag.String("server-type", "tcp", "Server type (either 'http' or 'tcp'). Default is tcp.")
	isSequential := flag.Bool("non-sequential", false, "If present, queue dispatches messages without receiving acknowledgement from workers. This may impact the sequentiality of operations performed by workers, especially if '-sleep' is present. Default is false.")
	sleep := flag.Bool("sleep", false, "If present, server simulates IDLE time for the workers before carrying on the work. Default is false.")
	flag.Parse()

	if *serverType != "http" {
		*serverType = "tcp"
	}

	return *serverType, *isSequential, *sleep

}
