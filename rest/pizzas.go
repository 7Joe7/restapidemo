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
		msg := "Request body for creation of pizza is invalid."
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	defer r.Body.Close()
	p := resources.NewPizza()
	err = json.Unmarshal(body, p)
	if err != nil {
		msg := "Request body for creation of pizza is not a valid JSON."
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	err = p.IsValid()
	if err != nil {
		msg := fmt.Sprintf("Pizza is invalid. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	for i := 0; i < len(p.Ingredients); i++ {
		exists, err := db.EntityExists(resources.DB_KEY_INGREDIENTS, p.Ingredients[i])
		if err != nil {
			msg := fmt.Sprintf("Database request to verify existence of ingredient with '%s' failed. %v", p.Ingredients[i], err)
			log.Printf(msg)
			http.Error(w, msg, 500)
			return
		}
		if !exists {
			msg := fmt.Sprintf("Ingredient with id '%s' doesn't exist and can't be assigned to a pizza.", p.Ingredients[i])
			log.Printf(msg)
			http.Error(w, msg, 400)
			return
		}
	}
	id, err := db.AddEntity(resources.DB_KEY_PIZZAS, p.ToMap())
	if err != nil {
		msg := fmt.Sprintf("Database request failed. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	log.Printf("Created pizza with id '%s'.", id)
	for i := 0; i < len(p.Ingredients); i++ {
		relId, err := db.AddEntity(fmt.Sprintf("%s:%s:ingredients", resources.DB_KEY_PIZZAS, id), map[string]string{"pizzaId":id,"ingredientId":p.Ingredients[i]})
		if err != nil {
			msg := fmt.Sprintf("Database request to create a relationship between pizza with id '%s' and ingredient with id '%s' failed. %v", id, p.Ingredients[i], err)
			log.Printf(msg)
			http.Error(w, msg, 500)
			return
		}
		log.Printf("Created relationship with id '%s' between pizza(%s) and ingredient(%s).", relId, id, p.Ingredients[i])
	}
	w.WriteHeader(201)
}

func getRestPizzas(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pizzas, err := db.GetAllEntities(resources.DB_KEY_PIZZAS)
	if err != nil {
		msg := fmt.Sprintf("Database request to retrieve all pizzas failed. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	pizzasJson, err := json.Marshal(pizzas)
	if err != nil {
		msg := fmt.Sprintf("Database value of pizzas altogether is invalid. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	_, err = w.Write(pizzasJson)
	if err != nil {
		msg := fmt.Sprintf("Writing response failed. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	w.WriteHeader(200)
}

func getRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	result, err := db.GetEntityById(resources.DB_KEY_PIZZAS, pid)
	if err != nil {
		msg := fmt.Sprintf("Database request to get pizza with id '%s' failed. %v", pid, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	pizzaJson, err := json.Marshal(result)
	if err != nil {
		msg := fmt.Sprintf("Database value of pizza with '%s' is invalid. %v", pid, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	_, err = w.Write(pizzaJson)
	if err != nil {
		msg := fmt.Sprintf("Writing response failed. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
	}
	w.WriteHeader(200)
}

func deleteRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	result, err := db.DeleteEntity(resources.DB_KEY_PIZZAS, pid)
	if err != nil {
		msg := fmt.Sprintf("Database request to delete pizza with id '%s' failed. %v", pid, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	if !result {
		msg := fmt.Sprintf("Pizza with id '%s' was not deleted.", pid)
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	w.WriteHeader(200)
	log.Printf("Deleted pizza with id '%s'.", pid)
}

func putRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "Request body is invalid."
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	defer r.Body.Close()
	p := resources.NewPizza()
	err = json.Unmarshal(body, p)
	if err != nil {
		msg := "Request body is not a valid JSON."
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	err = p.IsValid()
	if err != nil {
		msg := fmt.Sprintf("Pizza is invalid. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	for i := 0; i < len(p.Ingredients); i++ {
		exists, err := db.EntityExists(resources.DB_KEY_INGREDIENTS, p.Ingredients[i])
		if err != nil {
			msg := fmt.Sprintf("Database request to verify existence of ingredient with '%s' failed. %v", p.Ingredients[i], err)
			log.Printf(msg)
			http.Error(w, msg, 500)
			return
		}
		if !exists {
			msg := fmt.Sprintf("Ingredient with id '%s' doesn't exist and can't be assigned to a pizza.", p.Ingredients[i])
			log.Printf(msg)
			http.Error(w, msg, 400)
			return
		}
	}
	p.Id = pid
	err = db.UpdateEntity(resources.DB_KEY_PIZZAS, pid, p.ToMap())
	if err != nil {
		msg := fmt.Sprintf("Database request to update pizza with id '%s' failed.", pid)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	for i := 0; i < len(p.Ingredients); i++ {
		relId, err := db.AddEntity(fmt.Sprintf("%s:%s:ingredients", resources.DB_KEY_PIZZAS, pid), map[string]string{"pizzaId":pid,"ingredientId":p.Ingredients[i]})
		if err != nil {
			msg := fmt.Sprintf("Database request to create a relationship between pizza with id '%s' and ingredient with id '%s' failed. %v", pid, p.Ingredients[i], err)
			log.Printf(msg)
			http.Error(w, msg, 500)
			return
		}
		log.Printf("Created relationship with id '%s' between pizza(%s) and ingredient(%s).", relId, pid, p.Ingredients[i])
	}
	log.Printf("Updated pizza with id '%s'.", pid)
}

func postRestPizzasPidIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	exists, err := db.EntityExists(resources.DB_KEY_PIZZAS, pid)
	if err != nil {
		msg := fmt.Sprintf("Database request to verify existence of pizza with id '%s' failed. %v", pid, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	if !exists {
		msg := fmt.Sprintf("Specified pizza with id '%s' doesn't exist.", pid)
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "Request body for creation of relationship between pizza and ingredient is invalid."
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	defer r.Body.Close()
	iop := resources.NewIngredientOnPizza()
	err = json.Unmarshal(body, iop)
	if err != nil {
		msg := "Request body is not a valid JSON."
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	exists, err = db.EntityExists(resources.DB_KEY_INGREDIENTS, iop.IngredientId)
	if err != nil {
		msg := fmt.Sprintf("Database request to verify existence of ingredient with id '%s' failed. %v", iop.IngredientId, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	if !exists {
		msg := "Specified ingredient doesn't exist."
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	iop.PizzaId = pid
	relId, err := db.AddEntity(fmt.Sprintf("%s:%s:ingredients", resources.DB_KEY_PIZZAS, pid), iop.ToMap())
	if err != nil {
		msg := fmt.Sprintf("Database request to create a relationship between pizza with id '%s' and ingredient with id '%s' failed. %v", iop.PizzaId, iop.IngredientId, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	w.WriteHeader(201)
	log.Printf("Created relationship with id '%s' between pizza(%s) and ingredient(%s).", relId, iop.PizzaId, iop.IngredientId)
}

func getRestPizzasPidIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	exists, err := db.EntityExists(resources.DB_KEY_PIZZAS, pid)
	if err != nil {
		msg := fmt.Sprintf("Database request to verify existence of pizza with id '%s' failed. %v", pid, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	if !exists {
		msg := fmt.Sprintf("Specified pizza with id '%s' doesn't exist.", pid)
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	ingredientsOnPizza, err := db.GetAllEntities(fmt.Sprintf("%s:%s:ingredients", resources.DB_KEY_PIZZAS, pid))
	if err != nil {
		msg := fmt.Sprintf("Database request to retrieve all ingredients on pizza failed. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	ingredientsOnPizzaJson, err := json.Marshal(ingredientsOnPizza)
	if err != nil {
		msg := fmt.Sprintf("Database value of ingredients on pizza altogether is invalid. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	_, err = w.Write(ingredientsOnPizzaJson)
	if err != nil {
		msg := fmt.Sprintf("Writing response failed. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	w.WriteHeader(200)
}

func deleteRestPizzasPidIngredientsIopid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pid := params.ByName("pid")
	exists, err := db.EntityExists(resources.DB_KEY_PIZZAS, pid)
	if err != nil {
		msg := fmt.Sprintf("Database request to verify existence of pizza with id '%s' failed. %v", pid, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	if !exists {
		msg := fmt.Sprintf("Specified pizza with id '%s' doesn't exist.", pid)
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	iopid := params.ByName("iopid")
	result, err := db.DeleteEntity(fmt.Sprintf("%s:%s:ingredients", resources.DB_KEY_PIZZAS, pid), iopid)
	if err != nil {
		msg := fmt.Sprintf("Database request to delete ingredient on pizza with id '%s' failed. %v", iopid, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	if !result {
		msg := fmt.Sprintf("Ingredient on pizza with id '%s' was not deleted. %v", iopid, err)
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	w.WriteHeader(200)
	log.Printf("Deleted relationship with id '%s' on pizza with id '%s'.", iopid, pid)
}
