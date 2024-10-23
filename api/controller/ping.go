package controller

import (
	"fmt"
	"net/http"
	"os"
)

func (RC *RestController) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "- Application Version "+os.Getenv("APP_VERSION"))
}
