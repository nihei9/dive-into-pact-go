package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ErrorRecipeNotFound = errors.New("Recipe not found")
)

type Client struct {
	host string
	port int
	c    http.Client
}

type Recipe struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name string `json:"name"`
}

func New(host string, port int) *Client {
	return &Client{
		host: host,
		port: port,
		c:    http.Client{},
	}
}

func (c *Client) GetRecipe(id string) (*Recipe, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s:%v/v1/recipes/%s", c.host, c.port, id), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if res.StatusCode == http.StatusNotFound {
			return nil, ErrorRecipeNotFound
		}
		return nil, fmt.Errorf("Server returned error status: status = %d(%s)", res.StatusCode, http.StatusText(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	recipe := Recipe{}
	err = json.Unmarshal(body, &recipe)
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}
