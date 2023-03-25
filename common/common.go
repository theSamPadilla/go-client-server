// Package common declares common types used across the application
package common

// Type Client is the data structure used to communicate between the
// clint and the server.
type Client struct {
	Method string      `json:"method"`
	Key    interface{} `json:"key"`
	Value  interface{} `json:"value"`
}
