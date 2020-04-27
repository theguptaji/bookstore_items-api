package app

import (
	"net/http"

	"github.com/theguptaji/bookstore_items-api/src/controllers"
)

func MapUrls() {
	r.HandleFunc("/ping", controllers.PingController.Get).Methods(http.MethodGet)
	r.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	r.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	r.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)

}
