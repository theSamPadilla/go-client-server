package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/* Register Routes */
func RegisterRoutes() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", GetAllItems)
	router.GET("/:id", GetItemById)
	router.POST("/add", AddItem)
	router.POST("/remove", RemoveItem)

	return router
}

// @ GET
func GetAllItems(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "You hit the GetAllItems endpoit.\nURL: %s\n", r.URL.Path)
}

// @ GET
func GetItemById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "You hit the GetItemById endpoit. You passed in ID %s.\nURL: %s\n", ps.ByName("id"), r.URL.Path)
}

// @ POST
func AddItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "You hit the AddItem endpoit as a POST.\nURL: %s\n", r.URL.Path)
}

// @ POST
func RemoveItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "You hit the RemoveItem endpoit as a POST.\nURL: %s\n", r.URL.Path)
}
