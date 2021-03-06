package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

const (
	consumerName    = "cli"
	providerName    = "recipes"
	providerVersion = "1.0.0"
	pactDir         = "../../pacts"
	logDir          = "../../logs"
	broker          = "http://localhost:80"
)

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()

	go func() {
		http.HandleFunc("/v1/recipes/", HandleRecipes)
		http.HandleFunc("/setup", handleSetup)
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()

	pact := dsl.Pact{
		Consumer: consumerName,
		Provider: providerName,
		LogDir:   logDir,
		LogLevel: "INFO",
		PactDir:  pactDir,
	}

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port),
		// PactURLs: []string{
		// 	fmt.Sprintf("%s/%s-%s.json", pactDir, consumerName, providerName),
		// },
		ProviderStatesSetupURL: fmt.Sprintf("http://localhost:%d/setup", port),
		BrokerURL:              broker,
		Tags: []string{
			providerName,
			"latest",
		},
		PublishVerificationResults: true,
		ProviderVersion:            providerVersion,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func handleSetup(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var pState types.ProviderState
	json.Unmarshal(body, &pState)

	switch pState.State {
	case "Recipe exists":
		ClearRecipes()
		WriteRecipe("12345678", Recipe{
			ID:   "12345678",
			Name: "Sushi",
			Ingredients: []Ingredient{
				Ingredient{
					Name: "rice",
				},
				Ingredient{
					Name: "vinegar",
				},
			},
		})
	case "Recipe not found":
		ClearRecipes()
		WriteRecipe("00000001", Recipe{
			ID:   "00000001",
			Name: "TKG",
			Ingredients: []Ingredient{
				Ingredient{
					Name: "rice",
				},
				Ingredient{
					Name: "eggs",
				},
				Ingredient{
					Name: "soy sauce",
				},
			},
		})
	}
}
