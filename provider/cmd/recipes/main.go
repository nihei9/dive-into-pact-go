package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/nihei9/dive-into-pact-go/provider/handler"
)

func main() {
	var (
		port = flag.Int("p", 10000, "port")
	)
	flag.Parse()

	writeRecipes()

	http.HandleFunc("/v1/recipes/", handler.HandleRecipes)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func writeRecipes() {
	recipes := []handler.Recipe{
		handler.Recipe{
			ID:   "12345678",
			Name: "Sushi",
			Ingredients: []handler.Ingredient{
				handler.Ingredient{
					Name: "rice",
				},
				handler.Ingredient{
					Name: "vinegar",
				},
			},
		},
	}

	for _, recipe := range recipes {
		handler.WriteRecipe(recipe.ID, recipe)
	}
}
