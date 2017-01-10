package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/7joe7/pizzamanagement/resources"
	"github.com/7joe7/pizzamanagement/db"
)

func postRestIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "Request body to create ingredient is invalid."
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	defer r.Body.Close()
	in := resources.NewIngredient()
	err = json.Unmarshal(body, in)
	if err != nil {
		msg := "Request body to create ingredient is not a valid JSON."
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	err = in.IsValid()
	if err != nil {
		msg := fmt.Sprintf("Ingredient is invalid. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	id, err := db.AddEntity(resources.DB_KEY_INGREDIENTS, in.ToMap())
	if err != nil {
		msg := fmt.Sprintf("Database request failed. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	w.WriteHeader(201)
	log.Printf("Created ingredient with id '%s'.", id)
}

func getRestIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ingredients, err := db.GetAllEntities(resources.DB_KEY_INGREDIENTS)
	if err != nil {
		msg := fmt.Sprintf("Database request to retrieve all ingredients failed. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	ingredientsJson, err := json.Marshal(ingredients)
	if err != nil {
		msg := fmt.Sprintf("Database value of ingredients altogether is invalid. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	_, err = w.Write(ingredientsJson)
	if err != nil {
		msg := fmt.Sprintf("Writing response failed. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	w.WriteHeader(200)
}

func getRestIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	iid := params.ByName("iid")
	exists, err := db.EntityExists(resources.DB_KEY_INGREDIENTS, iid)
	if err != nil {
		msg := fmt.Sprintf("Database request to verify existence of ingredient with id '%s' failed. %v", iid, err)
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
	result, err := db.GetEntityById(resources.DB_KEY_INGREDIENTS, iid)
	if err != nil {
		msg := fmt.Sprintf("Database request to get ingredient with id '%s' failed. %v", iid, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	ingredientJson, err := json.Marshal(result)
	if err != nil {
		msg := fmt.Sprintf("Database value of ingredient with '%s' is invalid. %v", iid, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	_, err = w.Write(ingredientJson)
	if err != nil {
		msg := fmt.Sprintf("Writing response failed. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	w.WriteHeader(200)
}

func deleteRestIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	iid := params.ByName("iid")
	result, err := db.DeleteEntity(resources.DB_KEY_INGREDIENTS, iid)
	if err != nil {
		msg := fmt.Sprintf("Database request to delete ingredient with id '%s' failed. %v", iid, err)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	if !result {
		msg := fmt.Sprintf("Ingredient with id '%s' was not deleted. %v", iid, err)
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	w.WriteHeader(200)
	log.Printf("Deleted ingredient with id '%s'.", iid)
}

func putRestIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	iid := params.ByName("iid")
	exists, err := db.EntityExists(resources.DB_KEY_INGREDIENTS, iid)
	if err != nil {
		msg := fmt.Sprintf("Database request to verify existence of ingredient with id '%s' failed. %v", iid, err)
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "Request body is invalid."
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	defer r.Body.Close()
	in := resources.NewIngredient()
	err = json.Unmarshal(body, in)
	if err != nil {
		msg := "Request body is not a valid JSON."
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	err = in.IsValid()
	if err != nil {
		msg := fmt.Sprintf("Ingredient is invalid. %v", err)
		log.Printf(msg)
		http.Error(w, msg, 400)
		return
	}
	in.Id = iid
	err = db.UpdateEntity(resources.DB_KEY_INGREDIENTS, iid, in.ToMap())
	if err != nil {
		msg := fmt.Sprintf("Database request to update ingredient with id '%s' failed.", iid)
		log.Printf(msg)
		http.Error(w, msg, 500)
		return
	}
	log.Printf("Updated ingredient with id '%s'.", iid)
}
