package controller

import (
	"fmt"
	"net/http"
)

func (RC *RestController) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "- Test Commit 7")
}
