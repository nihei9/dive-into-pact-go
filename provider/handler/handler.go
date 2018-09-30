package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type Recipe struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name string `json:"name"`
}

var recipeBook = map[string]Recipe{}

func WriteRecipe(id string, recipe Recipe) {
	recipeBook[id] = recipe
}

func ClearRecipes() {
	recipeBook = map[string]Recipe{}
}

func HandleRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	re := regexp.MustCompile("^/v1/recipes/(\\d+)$")
	g := re.FindSubmatch([]byte(r.URL.Path))
	if len(g) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := string(g[1])
	recipe, ok := recipeBook[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, _ := json.Marshal(recipe)
	w.Write(b)
}
