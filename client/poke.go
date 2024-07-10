// Package httpclient provides a testing enviroment to test poke API server.
package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
)

// GetPokemonByName mimics a user performing a request to get a certain pokemon from the server
func (c *Client) GetPokemonByName(ctx context.Context, pokemonName string) (Pokemon, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.URL+c.endPoint+"/"+pokemonName, nil)

	if err != nil {
		return Pokemon{}, err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 3 * time.Second

	var (
		res    *http.Response
		resErr error
	)

	retryable := func() error {
		res, resErr = c.httpClient.Do(req)
		return resErr
	}

	notify := func(err error, t time.Duration) {
		log.Printf("error: %v happened at time: %v", err, t)
	}

	err = backoff.RetryNotify(retryable, b, notify)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}

	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code")
	}
	var pokemon Pokemon

	err = json.NewDecoder(res.Body).Decode(&pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

// GetAllPokemons mimics a user performing a request to get all pokemons from the server
func (c *Client) GetAllPokemons(ctx context.Context) ([]Pokemon, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.URL+c.endPoint, nil)

	if err != nil {
		return []Pokemon{}, err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 3 * time.Second

	var (
		res    *http.Response
		resErr error
	)

	retryable := func() error {
		res, resErr = c.httpClient.Do(req)
		return resErr
	}

	notify := func(err error, t time.Duration) {
		log.Printf("error: %v happened at time: %v", err, t)
	}

	err = backoff.RetryNotify(retryable, b, notify)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}

	if err != nil {
		return []Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []Pokemon{}, fmt.Errorf("unexpected status code")
	}

	var body BodyFormat

	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return []Pokemon{}, err
	}

	return body.Pokemons, nil
}
