package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	DEFAULT_PORT int = 49494
)

func main() {
	log.Printf("Starting pizza management microservice on port %d.", DEFAULT_PORT)

	router := httprouter.New()

	// TODO rest api for this service
	// TODO callers to other services

	// get pizzas
	// get pizzas/id
	// get pizzas/id/ingredients
	// post pizzas
	// post pizzas/id/ingredients
	// delete pizzas/id
	// delete pizzas/id/ingredients/id
	// put pizzas/id
	// put pizzas/id/ingredients/id

	err := http.ListenAndServe(fmt.Sprintf(":%d", DEFAULT_PORT), router)
	if err != nil {
		log.Fatalf("ListenAndServe failed. %v", err)
	}
	log.Printf("Finishing pizza management microservice.")
}
