package rest

import (
	"net/http"
	"encoding/json"
	"log"

	"github.com/julienschmidt/httprouter"
)

func getRestPizzas(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// TODO get pizzas from redis
}

func postRestPizzas(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p Pizza
	err := decoder.Decode(&p)
	if err != nil {
		http.Error(w, "Request body is not a valid JSON.", 400)
	}
	defer r.Body.Close()
	// TODO save p
}

func getRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func getRestPizzasPidIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func postRestPizzasPidIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func deleteRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func deleteRestPizzasPidIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func putRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func putRestPizzasPidIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func logRequest(h httprouter.Handle) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		log.Printf("Accepted request from '%s' to address '%s', method '%s'.", r.RemoteAddr, r.RequestURI, r.Method)
		h(w, r, params)
	}
}
