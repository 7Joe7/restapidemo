package rest

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

/*
 Returns a list of pizzas in JSON format

 JSON format:
 [{"id":"","name":""}]
 */
func GetRestPizzas(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getRestPizzas(w, r, params)
}

/*
 Adds a new pizza

 Accepts JSON:
 {"name":"","ingredients":["id1", "id2"]}
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
 Returns a list of ingredients
 */
func GetRestIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getRestIngredients(w, r, params)
}

/*
 Returns information about ingredient specified by id
 */
func GetRestIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getRestIngredientsIid(w, r, params)
}

/*
 Adds an ingredient

 Accepts JSON:
 {"name":""}
 */
func PostRestIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postRestIngredients(w, r, params)
}

/*
 Deletes a pizza specified by its id
 */
func DeleteRestPizzasPid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	deleteRestPizzasPid(w, r, params)
}

/*
 Deletes an ingredient specified by its id
 */
func DeleteRestIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	deleteRestIngredientsIid(w, r, params)
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
 Modifies an ingredient specified by its id

 Accepts JSON:
 {"name":""}
 */
func PutRestIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	putRestIngredientsIid(w, r, params)
}

/*
 Logs request parameters
 */
func LogRequest(h httprouter.Handle) httprouter.Handle {
	return logRequest(h)
}
