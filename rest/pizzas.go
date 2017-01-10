package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/7joe7/pizzamanagement/resources"
	"github.com/7joe7/pizzamanagement/db"
)

func postRestPizzas(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Request body is invalid.", 403)
		return
	}
	defer r.Body.Close()
	p := resources.NewPizza()
	err = json.Unmarshal(body, p)
	if err != nil {
		http.Error(w, "Request body is not a valid JSON.", 403)
		return
	}
	err = p.IsValid()
	if err != nil {
		http.Error(w, fmt.Sprintf("Pizza is invalid. %v", err), 403)
		return
	}
	for i := 0; i < len(p.Ingredients); i++ {
		exists, err := db.EntityExists(resources.DB_KEY_INGREDIENTS, p.Ingredients[i])
		if err != nil {
			http.Error(w, fmt.Sprintf("Database request to verify existence of ingredient with '%s' failed. %v", p.Ingredients[i], err), 500)
			return
		}
		if !exists {
			http.Error(w, fmt.Sprintf("Ingredient with id '%s' doesn't exist and can't be assigned to a pizza.", p.Ingredients[i]), 403)
			return
		}
	}
	id, err := db.AddEntity(resources.DB_KEY_PIZZAS, p.ToMap())
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request failed. %v", err), 500)
		return
	}
	for i := 0; i < len(p.Ingredients); i++ {
		relId, err := db.AddEntity(fmt.Sprintf("%s:%s:ingredients", resources.DB_KEY_PIZZAS, id), map[string]string{"PizzaId":id,"IngredientId":p.Ingredients[i]})
		if err != nil {
			http.Error(w, fmt.Sprintf("Database request to create a relationship between pizza with id '%s' and ingredient with id '%s' failed. %v", id, p.Ingredients[i], err), 500)
			return
		}
		log.Printf("Created relationship with id '%s' between pizza(%s) and ingredient(%s).", relId, id, p.Ingredients[i])
	}
	w.WriteHeader(201)
	log.Printf("Created pizza with id '%s'.", id)
}

func getRestPizzas(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pizzas, err := db.GetAllEntities(resources.DB_KEY_PIZZAS)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request to retrieve all pizzas failed. %v", err), 500)
		return
	}
	pizzasJson, err := json.Marshal(pizzas)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database value of pizzas altogether is invalid. %v", err), 500)
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
	pid := params.ByName("pid")
	result, err := db.GetEntityById(resources.DB_KEY_PIZZAS, pid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request to get pizza with id '%s' failed. %v", pid, err), 500)
		return
	}
	pizzaJson, err := json.Marshal(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database value of pizza with '%s' is invalid. %v", pid, err), 500)
		return
	}
	_, err = w.Write(pizzaJson)
	if err != nil {
		http.Error(w, fmt.Sprintf("Writing response failed. %v", err), 500)
	}
	w.WriteHeader(200)
}

func deleteRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	result, err := db.DeleteEntity(resources.DB_KEY_PIZZAS, pid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request to delete pizza with id '%s' failed. %v", pid, err), 500)
		return
	}
	if !result {
		http.Error(w, fmt.Sprintf("Pizza with id '%s' was not deleted.", pid), 403)
	}
	w.WriteHeader(200)
	log.Printf("Deleted pizza with id '%s'.", pid)
}

func putRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Request body is invalid.", 403)
		return
	}
	defer r.Body.Close()
	p := resources.NewPizza()
	err = json.Unmarshal(body, p)
	if err != nil {
		http.Error(w, "Request body is not a valid JSON.", 403)
		return
	}
	err = p.IsValid()
	if err != nil {
		http.Error(w, fmt.Sprintf("Pizza is invalid. %v", err), 403)
		return
	}
	err = db.UpdateEntity(resources.DB_KEY_PIZZAS, pid, p.ToMap())
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request to update pizza with id '%s' failed.", pid), 500)
		return
	}
	log.Printf("Updated pizza with id '%s'.", pid)
}

func postRestPizzasPidIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Request body is invalid.", 403)
		return
	}
	defer r.Body.Close()
	iop := resources.NewIngredientOnPizza()
	err = json.Unmarshal(body, iop)
	if err != nil {
		http.Error(w, "Request body is not a valid JSON.", 403)
		return
	}
	iop.PizzaId = pid
	relId, err := db.AddEntity(fmt.Sprintf("%s:%s:ingredients", resources.DB_KEY_PIZZAS, pid), iop.ToMap())
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request to create a relationship between pizza with id '%s' and ingredient with id '%s' failed. %v", iop.PizzaId, iop.IngredientId, err), 500)
		return
	}
	w.WriteHeader(201)
	log.Printf("Created relationship with id '%s' between pizza(%s) and ingredient(%s).", relId, iop.PizzaId, iop.IngredientId)
}

func getRestPizzasPidIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	ingredientsOnPizza, err := db.GetAllEntities(fmt.Sprintf("%s:%s:ingredients", resources.DB_KEY_PIZZAS, pid))
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request to retrieve all ingredients on pizza failed. %v", err), 500)
		return
	}
	ingredientsOnPizzaJson, err := json.Marshal(ingredientsOnPizza)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database value of ingredients on pizza altogether is invalid. %v", err), 500)
		return
	}
	_, err = w.Write(ingredientsOnPizzaJson)
	if err != nil {
		http.Error(w, fmt.Sprintf("Writing response failed. %v", err), 500)
		return
	}
	w.WriteHeader(200)
}

func deleteRestPizzasPidIngredientsIopid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	iopid := params.ByName("iopid")
	result, err := db.DeleteEntity(fmt.Sprintf("%s:%s:ingredients", resources.DB_KEY_PIZZAS, pid), iopid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database request to delete ingredient on pizza with id '%s' failed. %v", iopid, err), 500)
		return
	}
	if !result {
		http.Error(w, fmt.Sprintf("Ingredient on pizza with id '%s' was not deleted. %v", iopid, err), 403)
		return
	}
	w.WriteHeader(200)
	log.Printf("Deleted relationship with id '%s' on pizza with id '%s'.", iopid, pid)
}
