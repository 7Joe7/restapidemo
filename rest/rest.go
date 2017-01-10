package rest

import (
	"fmt"
	"net/http"
	"log"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"
)

func logRequest(h httprouter.Handle) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		log.Printf("Accepted request from '%s' to address '%s', method '%s'.", r.RemoteAddr, r.RequestURI, r.Method)
		defer func () {
			if err := recover(); err != nil {
				log.Printf("Panic occurred while responding to %s, method %s. %v\n", r.RequestURI, r.Method, err)
				log.Printf("Failure stacktrace: %s\n", string(debug.Stack()))
				http.Error(w, fmt.Sprintf("Server error. %v", err), 500)
			}
		}()
		h(w, r, params)
	}
}
