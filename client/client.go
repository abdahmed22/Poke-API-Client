// Package httpclient provides a testing enviroment to test poki API server.
package httpclient

import (
	"net/http"
	"os"
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
