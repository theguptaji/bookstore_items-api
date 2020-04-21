package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	r = mux.NewRouter()
)

func StartApplication() {
	MapUrls()

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}

	if err := srv.ListenAndServe(); err != nil {

	}
}
