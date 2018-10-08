package client

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

const (
	consumerName    = "cli"
	consumerVersion = "1.0.0"
	providerName    = "recipes"
	pactDir         = "../../pacts"
	logDir          = "../../logs"
	broker          = "http://localhost:80"
)

var pact dsl.Pact

func TestMain(m *testing.M) {
	pact = dsl.Pact{
		Consumer:                 consumerName,
		Provider:                 providerName,
		LogDir:                   logDir,
		PactDir:                  pactDir,
		LogLevel:                 "INFO",
		DisableToolValidityCheck: true,
		PactFileWriteMode:        "overwrite",
	}

	exitCode := m.Run()

	pact.WritePact()
	pact.Teardown()

	publisher := dsl.Publisher{}
	err := publisher.Publish(types.PublishRequest{
		PactBroker: broker,
		PactURLs: []string{
			fmt.Sprintf("%s/%s-%s.json", pactDir, consumerName, providerName),
		},
		ConsumerVersion: consumerVersion,
		Tags: []string{
			providerName,
		},
	})
	if err != nil {
		log.Println("ERROR: ", err)
	}

	os.Exit(exitCode)
}

func TestPactConsumerRecipesHandler_RecipeExists(t *testing.T) {
	var testSushiExists = func() error {
		c := New("localhost", pact.Server.Port)
		_, err := c.GetRecipe("12345678")
		if err != nil {
			return err
		}

		return nil
	}

	pact.
		AddInteraction().
		Given("Recipe exists").
		UponReceiving("A request to get a recipe").
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.Term("/v1/recipes/12345678", "/v1/recipes/[0-9a-z]+"),
		}).
		WillRespondWith(dsl.Response{
			Status: http.StatusOK,
			Body: dsl.Like(map[string]interface{}{
				"id":   dsl.Like("12345678"),
				"name": dsl.Like("Sushi"),
				"ingredients": dsl.EachLike(map[string]interface{}{
					"name": dsl.Like("rice"),
				}, 1),
			}),
			Headers: dsl.MapMatcher{
				"Content-Type": dsl.Term("application/json", `application\/json`),
			},
		})

	err := pact.Verify(testSushiExists)
	if err != nil {
		t.Fatalf("Error on Verify: %v", err)
	}
}

func TestPactConsumerRecipesHandler_RecipeNotFound(t *testing.T) {
	var testSushiExists = func() error {
		c := New("localhost", pact.Server.Port)
		_, err := c.GetRecipe("99999999")
		if err != ErrorRecipeNotFound {
			return err
		}

		return nil
	}

	pact.
		AddInteraction().
		Given("Recipe not found").
		UponReceiving("A request to get a recipe").
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.Term("/v1/recipes/99999999", "/v1/recipes/[0-9a-z]+"),
		}).
		WillRespondWith(dsl.Response{
			Status: http.StatusNotFound,
			Body:   nil,
			Headers: dsl.MapMatcher{
				"Content-Type": dsl.Term("application/json", `application\/json`),
			},
		})

	err := pact.Verify(testSushiExists)
	if err != nil {
		t.Fatalf("Error on Verify: %v", err)
	}
}
