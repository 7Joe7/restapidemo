package rest

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

/*
 Returns a list of pizzas in JSON format

 JSON format:
 ["id":{"name":"","ingredients":[]}]
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
func GetRestIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getRestPizzasPidIngredients(w, r, params)
}

/*
 Adds an ingredient to a pizza specified by its id

 Accepts JSON:
 {"name":""}
 */
func PostRestIngredients(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
func DeleteRestIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
func PutRestIngredientsIid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	putRestPizzasPidIngredientsIid(w, r, params)
}

/*
 Logs request parameters
 */
func LogRequest(h httprouter.Handle) httprouter.Handle {
	return logRequest(h)
}
