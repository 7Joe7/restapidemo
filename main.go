package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/7joe7/pizzamanagement/rest"
	"github.com/7joe7/pizzamanagement/db"
)

const (
	DEFAULT_PORT = 49494
	REDIS_PORT   = 6379
	REDIS_HOST   = "localhost"
)

func init() {
	log.Printf("Starting pizza management microservice on port %d.", DEFAULT_PORT)
	redisAddress := fmt.Sprintf("%s:%d", REDIS_HOST, REDIS_PORT)
	err := db.SetupConnection(redisAddress)
	if err != nil {
		log.Printf("Connection to database failed. %v", err)
		os.Exit(-1)
	}
}

func main() {
	log.Printf("Setting up REST API.")

	router := httprouter.New()

	router.GET("/rest/pizzas", rest.LogRequest(rest.GetRestPizzas))
	router.POST("/rest/pizzas", rest.LogRequest(rest.PostRestPizzas))

	router.GET("/rest/pizzas/:pid", rest.LogRequest(rest.GetRestPizzasPid))
	router.DELETE("/rest/pizzas/:pid", rest.LogRequest(rest.DeleteRestPizzasPid))

	router.GET("/rest/pizzas/:pid/ingredients", rest.LogRequest(rest.GetRestPizzasPidIngredients))
	router.POST("/rest/pizzas/:pid/ingredients", rest.LogRequest(rest.PostRestPizzasPidIngredients))
	router.PUT("/rest/pizzas/:pid", rest.LogRequest(rest.PutRestPizzasPid))

	router.PUT("/rest/pizzas/:pid/ingredients/:iid", rest.LogRequest(rest.PutRestPizzasPidIngredientsIid))
	router.DELETE("/rest/pizzas/:pid/ingredients/:iid", rest.LogRequest(rest.DeleteRestPizzasPidIngredientsIid))

	err := http.ListenAndServe(fmt.Sprintf(":%d", DEFAULT_PORT), router)
	if err != nil {
		log.Fatalf("ListenAndServe failed. %v", err)
	}
	log.Printf("Finishing pizza management microservice.")
}
