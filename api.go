package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

/*
 Returns a list of pizzas in JSON format
 */
func GetRestPizzas(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getRestPizzas(w, r, params)
}

/*
 Adds a new pizza

 Accepts JSON:
 {"name":"","ingredients":[]}
 */
func PostRestPizzas(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postRestPizzas(w, r, params)
}

/*
 Returns a pizza with details based on its id in JSON format
 */
func GetRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getRestPizzasPid(w, r, params)
}

/*
 Returns a list of ingredients of a pizza specified by its id
 */
func GetRestPizzasPidIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getRestPizzasPidIngredients(w, r, params)
}

/*
 Adds an ingredient to a pizza specified by its id

 Accepts JSON:
 {"name":""}
 */
func PostRestPizzasPidIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postRestPizzasPidIngredients(w, r, params)
}

/*
 Deletes a pizza specified by its id
 */
func DeleteRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	deleteRestPizzasPid(w, r, params)
}

/*
 Deletes an ingredient from a pizza specified by pizza id and ingredient id
 */
func DeleteRestPizzasPidIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	deleteRestPizzasPidIngredientsIid(w, r, params)
}

/*
 Modifies a pizza specified by its id

 Accepts JSON:
 {"name":"","ingredients":[]}
 */
func PutRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	putRestPizzasPid(w, r, params)
}

/*
 Modifies an ingredient of a pizza specified byt the pizza id and ingredient id

 Accepts JSON:
 {"name":""}
 */
func PutRestPizzasPidIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	putRestPizzasPidIngredientsIid(w, r, params)
}
