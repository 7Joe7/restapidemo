package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	redis "gopkg.in/redis.v5"
)

const (
	DEFAULT_PORT int = 49494
)

func main() {
	log.Printf("Starting pizza management microservice on port %d.", DEFAULT_PORT)

	router := httprouter.New()

	router.GET("/rest/pizzas", GetRestPizzas)
	router.POST("/rest/pizzas", PostRestPizzas)

	router.GET("/rest/pizzas/:pid", GetRestPizzasPid)
	router.DELETE("/rest/pizzas/:pid", DeleteRestPizzasPid)

	router.GET("/rest/pizzas/:pid/ingredients", GetRestPizzasPidIngredients)
	router.POST("/rest/pizzas/:pid/ingredients", PostRestPizzasPidIngredients)
	router.PUT("/rest/pizzas/:pid", PutRestPizzasPid)

	router.PUT("/rest/pizzas/:pid/ingredients/:iid", PutRestPizzasPidIngredientsIid)
	router.DELETE("/rest/pizzas/:pid/ingredients/:iid", DeleteRestPizzasPidIngredientsIid)

	// TODO storage

	err := http.ListenAndServe(fmt.Sprintf(":%d", DEFAULT_PORT), router)
	if err != nil {
		log.Fatalf("ListenAndServe failed. %v", err)
	}
	log.Printf("Finishing pizza management microservice.")
}
