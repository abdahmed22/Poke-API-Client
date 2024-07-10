// Package httpclient provides a testing enviroment to test poki API server.
package httpclient

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Client struct {
	httpClient *http.Client
	URL        string
	endPoint   string
}

type Option func(c *Client)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
}

type BodyFormat struct {
	Pokemons []Pokemon `json:"results"`
}

// NewClient creates a new client
func NewClient(options ...Option) *Client {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	client := &Client{
		httpClient: http.DefaultClient,
		URL:        os.Getenv("URL"),
		endPoint:   os.Getenv("ENDPOINT"),
	}

	for _, option := range options {
		option(client)
	}
	return client
}

// CustomURL provides the option to change default URL
func CustomURL(URL string) Option {
	return func(c *Client) {
		c.URL = URL
	}
}

// CustomEndPoint provides the option to change default endpoint
func CustomEndPoint(endPoint string) Option {
	return func(c *Client) {
		c.endPoint = endPoint
	}
}

// CustomClient provides the option to change default Client
func CustomClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
