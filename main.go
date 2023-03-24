package main

import (
	"fmt"
	"net/http"

	"client-server/orderedmap"
	"client-server/server"
)

func main() {
	//Initialize ordered map and server
	om := orderedmap.Constructor()
	r := server.RegisterRoutes(om)

	fmt.Println("\nServing responses on port 6969 and waiting for clients...")
	http.ListenAndServe(":6969", r)
}
