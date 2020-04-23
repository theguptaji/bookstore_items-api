package app

import (
	"net/http"
	"time"

	"github.com/theguptaji/bookstore_items-api/clients/elasticsearch"

	"github.com/gorilla/mux"
)

var (
	r = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	MapUrls()

	srv := &http.Server{
		Handler:      r,
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Addr:         "127.0.0.1:8000",
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
