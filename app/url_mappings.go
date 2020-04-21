package app

import (
	"net/http"

	"github.com/theguptaji/bookstore_items-api/controllers"
)

func MapUrls() {
	r.HandleFunc("/ping", controllers.PingController.Get).Methods(http.MethodGet)
	r.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
