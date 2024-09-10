package controller

import "net/http"

type IRestController interface {
	PingHandler(w http.ResponseWriter, r *http.Request)
}

type RestController struct {
}

func NewRestController() IRestController {
	return &RestController{}
}
