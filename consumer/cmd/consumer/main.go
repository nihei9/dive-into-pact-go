package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nihei9/dive-into-pact-go/consumer/client"
)

func main() {
	var (
		host = flag.String("h", "localhost", "Host")
		port = flag.Int("p", 10000, "Port")
	)
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Please specify a recipe id")
		os.Exit(1)
	}

	c := client.New(*host, *port)
	recipe, err := c.GetRecipe(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(recipe.ID, recipe.Name)
	for _, ingredient := range recipe.Ingredients {
		fmt.Println("*", ingredient.Name)
	}
}
