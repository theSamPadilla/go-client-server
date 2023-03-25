// Package common declares common types and functions used across packages
package common

import (
	"log"
	"os"
)

// Type Client is the data structure used to communicate between the
// clint and the server.
type Client struct {
	Method string      `json:"method"`
	Key    interface{} `json:"key"`
	Value  interface{} `json:"value"`
}

var (
	WARN   *log.Logger
	INFO   *log.Logger
	RESULT *log.Logger
)

// Initialize 3 global loggers: info, warning and result (error logs written to stdout)
func SetLoggers(path string, nonSequential bool, sleep bool) {
	//Set different log file depending on config
	switch {
	case nonSequential && sleep:
		path = "logs/sleep_nonSequential_server.log"
	case nonSequential:
		path = "logs/nonSequential_server.log"
	case sleep:
		path = "logs/sleep_server.log"
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	INFO = log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	WARN = log.New(file, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	RESULT = log.New(file, "[RESULT] ", log.Ldate|log.Ltime|log.Lshortfile)

	log.Printf("Writing logs to %s\n", path)

}
