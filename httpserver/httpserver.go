package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/thesampadilla/go-client-server/orderedmap"

	"github.com/julienschmidt/httprouter"
)

// Builds the router and registers the routes
func RegisterRoutes(om *orderedmap.OrderedMap) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", GetAllItems(om))
	router.GET("/key/:key", GetItemByKey(om))
	router.GET("/index/:index", GetItemByIndex(om))
	router.POST("/add", AddItem(om))
	router.POST("/remove", RemoveItem(om))

	return router
}

// @ GET /
func GetAllItems(om *orderedmap.OrderedMap) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		result := om.GetAllItemsInOrder()
		fmt.Fprintf(w, "Result:\n%s\n", result)
	}
}

// @ GET /key/:key
func GetItemByKey(om *orderedmap.OrderedMap) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		result, err := om.GetItemByKey(ps.ByName("key"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Fprintf(w, "Result:\n%+v\n", result)
	}
}

// @ GET /index/:index
func GetItemByIndex(om *orderedmap.OrderedMap) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//Check index validity
		ui64, err := strconv.ParseUint(ps.ByName("index"), 10, 64)
		if err != nil {
			fmt.Fprintf(w, "Invalid index. Index must be a positive integer or zero.\n")
			return
		}

		//Get item
		result, err := om.GetItemByIndex(ui64)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintf(w, "Result:\n%+v\n", result)
	}
}

// @ POST
// Requires `key` and `value` in body of request
func AddItem(om *orderedmap.OrderedMap) httprouter.Handle {
	type NewItem struct {
		K interface{} `json:"key"`
		V interface{} `json:"value"`
	}

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		//Check for JSON format of the body
		contentType := r.Header.Get("Content-type")
		if contentType != "application/json" {
			fmt.Fprintf(w, "Invalid Content-type %s\n", contentType)
			return
		}

		//Read the body of the request
		b, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "An error occurred while parsing the content of the request\n")
			return
		}

		//Unmarshal the json and add item to ordered map
		var item NewItem
		err = json.Unmarshal(b, &item)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		om.SetItem(item.K, item.V)
		fmt.Fprintf(w, "Successfully added %s->%s.\n", item.K, item.V)
	}
}

// @ POST
// Requires `key` in body of the request
func RemoveItem(om *orderedmap.OrderedMap) httprouter.Handle {
	type RemoveItem struct {
		K interface{} `json:"key"`
	}

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		//Check for JSON format of the body
		contentType := r.Header.Get("Content-type")
		if contentType != "application/json" {
			fmt.Fprintf(w, "Invalid Content-type %s\n", contentType)
			return
		}

		//Read the body of the request
		b, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "An error occurred while parsing the content of the request\n")
			return
		}

		//Unmarshal the json and remove from ordered map
		var item RemoveItem
		err = json.Unmarshal(b, &item)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		v, err := om.RemoveItemByKey(item.K)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintf(w, "Successfully removed %s->%s from the map.\n", item.K, v)
	}
}
