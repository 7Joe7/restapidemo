package rest

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"github.com/7joe7/pizzamanagement/resources"
	"github.com/7joe7/pizzamanagement/db"
)

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
	err = db.AddEntity(resources.DB_KEY_PIZZAS, p.ToMap())
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request failed. %v", err), 500)
		return
	}
	w.WriteHeader(201)
}

func getRestPizzas(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pizzas, err := db.GetAllEntities(resources.DB_KEY_PIZZAS)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request failed. %v", err), 500)
		return
	}
	pizzasJson, err := json.Marshal(pizzas)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database value id invalid. %v", err), 500)
		return
	}
	_, err = w.Write(pizzasJson)
	if err != nil {
		http.Error(w, fmt.Sprintf("Writing response failed. %v", err), 500)
		return
	}
	w.WriteHeader(200)
}

func getRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//pid := params.ByName("pid")
	//pizza, err := db.GetEntityById(resources.DB_KEY_PIZZAS, pid)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Database request for pizza with id '%s' failed. %v", pid, err), 500)
	//	return
	//}
	//p := resources.NewPizza()
	//err = json.Unmarshal([]byte(pizza), p)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Database value is invalid. %v", err), 500)
	//	return
	//}
	//_, err = w.Write([]byte(pizza))
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Unable to write response. %v", err), 500)
	//	return
	//}
	//w.WriteHeader(200)
}

func deleteRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	result, err := db.DeleteEntity(resources.DB_KEY_PIZZAS, pid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request to delete pizza with id '%s' failed. %v", pid, err), 500)
		return
	}
	if !result {
		http.Error(w, fmt.Sprintf("Pizza with id '%s' was not deleted.", pid), 400)
	}
	w.WriteHeader(200)
}

func putRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//pid := params.ByName("pid")
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, "Request body is invalid.", 400)
	//	return
	//}
	//p := resources.NewPizza()
	//err = json.Unmarshal(body, p)
	//if err != nil {
	//	http.Error(w, "Request body is not a valid JSON.", 400)
	//	return
	//}
	//defer r.Body.Close()
	//err = p.IsValid()
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Pizza is invalid. %v", err), 400)
	//	return
	//}
	//result, err := db.UpdateEntity(resources.DB_KEY_PIZZAS, pid, body)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Database request to modify pizza with id "))
	//	return
	//}
	//
}
