package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/7joe7/pizzamanagement/rest"
	"github.com/7joe7/pizzamanagement/db"
	"github.com/7joe7/pizzamanagement/resources"
)

func init() {
	log.Printf("Starting pizza management microservice on port %d.", resources.DEFAULT_PORT)
	redisAddress := fmt.Sprintf("%s:%d", resources.REDIS_HOST, resources.REDIS_PORT)
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
	router.PUT("/rest/pizzas/:pid", rest.LogRequest(rest.PutRestPizzasPid))

	router.GET("/rest/pizzas/:pid/ingredients", rest.LogRequest(rest.GetRestPizzasPidIngredients))
	router.POST("/rest/pizzas/:pid/ingredients", rest.LogRequest(rest.PostRestPizzasPidIngredients))

	router.DELETE("/rest/pizzas/:pid/ingredients/:iopid", rest.LogRequest(rest.DeleteRestPizzasPidIngredientsIopid))

	router.GET("/rest/ingredients", rest.LogRequest(rest.GetRestIngredients))
	router.POST("/rest/ingredients", rest.LogRequest(rest.PostRestIngredients))

	router.GET("/rest/ingredients/:iid", rest.LogRequest(rest.GetRestIngredientsIid))
	router.PUT("/rest/ingredients/:iid", rest.LogRequest(rest.PutRestIngredientsIid))
	router.DELETE("/rest/ingredients/:iid", rest.LogRequest(rest.DeleteRestIngredientsIid))

	err := http.ListenAndServe(fmt.Sprintf(":%d", resources.DEFAULT_PORT), router)
	if err != nil {
		log.Fatalf("ListenAndServe failed. %v", err)
	}
	log.Printf("Finishing pizza management microservice.")
}
