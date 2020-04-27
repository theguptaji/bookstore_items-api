package controllers

import "net/http"

const (
	pong = "pong"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Get(http.ResponseWriter, *http.Request)
}

type pingController struct{}

func (i *pingController) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
