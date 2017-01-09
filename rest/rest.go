package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/7joe7/pizzamanagement/db"
	"github.com/7joe7/pizzamanagement/resources"
	"runtime/debug"
)

func getRestPizzas(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pizzas, err := db.GetAll(resources.DB_KEY_PIZZAS)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request failed. %v", err), 500)
	}
	w.Write()
	// TODO get pizzas from redis
}

func postRestPizzas(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Request body is invalid.", 400)
		return
	}
	p := resources.NewPizza()
	err = json.Unmarshal(body, p)
	if err != nil {
		http.Error(w, "Request body is not a valid JSON.", 400)
		return
	}
	defer r.Body.Close()
	err = p.IsValid()
	if err != nil {
		http.Error(w, fmt.Sprintf("Pizza is invalid. %v", err), 400)
		return
	}
	err = db.Add(resources.DB_KEY_PIZZAS, string(body))
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request failed. %v", err), 500)
		return
	}
	w.WriteHeader(201)
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
