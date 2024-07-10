package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// func TestClientCanHitPokeAPIServer(t *testing.T) {

// 	t.Run("can hit the poke api server and return ditto information", func(*testing.T) {

// 		myClient := NewClient()
// 		returnedPokemon, err := myClient.GetPokemonByName(context.Background(), "ditto")

// 		expectedPokemon := Pokemon{
// 			ID:             132,
// 			Name:           "ditto",
// 			BaseExperience: 101,
// 			Height:         3,
// 		}

// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedPokemon, returnedPokemon)
// 	})

// 	t.Run("can hit the poke api server and return charmeleon information", func(*testing.T) {

// 		myClient := NewClient()
// 		returnedPokemon, err := myClient.GetPokemonByName(context.Background(), "charmeleon")

// 		expectedPokemon := Pokemon{
// 			ID:             5,
// 			Name:           "charmeleon",
// 			BaseExperience: 142,
// 			Height:         11,
// 		}

// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedPokemon, returnedPokemon)
// 	})

// 	t.Run("can hit the poke api server and return squirtle information", func(*testing.T) {

// 		myClient := NewClient()
// 		returnedPokemon, err := myClient.GetPokemonByName(context.Background(), "squirtle")

// 		expectedPokemon := Pokemon{
// 			ID:             7,
// 			Name:           "squirtle",
// 			BaseExperience: 63,
// 			Height:         5,
// 		}

// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedPokemon, returnedPokemon)
// 	})

// 	t.Run("can hit the poke api server and return all pokemon information", func(*testing.T) {

// 		myClient := NewClient()
// 		returnedPokemonSlice, err := myClient.GetAllPokemons(context.Background())

// 		assert.NoError(t, err)
// 		assert.NotEmpty(t, returnedPokemonSlice)
// 	})
// }

func TestOptionFunctions(t *testing.T) {
	t.Run("happy path - can add custom URLS using option function", func(*testing.T) {

		myClient := NewClient(
			CustomURL("custom_url"),
		)

		assert.Equal(t, "custom_url", myClient.URL)
	})

	t.Run("happy path - can add custom client using option function", func(*testing.T) {

		myClient := NewClient(
			CustomEndPoint("/endpoint"),
		)

		assert.Equal(t, "/endpoint", myClient.endPoint)
	})

	t.Run("happy path - can add custom client using option function", func(*testing.T) {

		myClient := NewClient(
			CustomClient(&http.Client{
				Timeout: 5 * time.Second,
			}),
		)

		assert.Equal(t, 5*time.Second, myClient.httpClient.Timeout)
	})

}

func TestClientGetPokemonByNameCanHitMockServer(t *testing.T) {
	t.Run("can hit the mockserver and return pokemon information", func(*testing.T) {

		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, `{"id": 7, "name": "squirtle", "base_experience": 63, "height": 5}`)
				}),
		)

		defer mockServer.Close()

		myClient := NewClient(
			CustomURL(mockServer.URL),
		)

		returnedPokemon, err := myClient.GetPokemonByName(context.Background(), "squirtle")

		expectedPokemon := Pokemon{
			ID:             7,
			Name:           "squirtle",
			BaseExperience: 63,
			Height:         5,
		}

		assert.NoError(t, err)

		assert.Equal(t, expectedPokemon, returnedPokemon)

	})

	t.Run("can handle 500 status code", func(*testing.T) {

		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				}),
		)

		defer mockServer.Close()

		myClient := NewClient(
			CustomURL(mockServer.URL),
		)

		_, err := myClient.GetPokemonByName(context.Background(), "squirtle")

		assert.Error(t, err)

	})

	t.Run("can handle wrong json format", func(*testing.T) {

		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, `"id": 7, "name": "squirtle", "base_experience": 63, "height": 5`)
				}),
		)

		defer mockServer.Close()

		myClient := NewClient(
			CustomURL(mockServer.URL),
		)

		_, err := myClient.GetPokemonByName(context.Background(), "squirtle")

		assert.Error(t, err)

	})
}

func TestClientGetAllPokemonsCanHitMockServer(t *testing.T) {
	t.Run("can hit the mockserver and return all pokemon information", func(*testing.T) {

		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, `{"results": [{"id": 132, "name": "ditto", "base_experience": 101, "height": 3},
					{"id": 5, "name": "charmeleon", "base_experience": 142, "height": 11},
					{"id": 7, "name": "squirtle", "base_experience": 63, "height": 5}]}`)
				}),
		)

		defer mockServer.Close()

		myClient := NewClient(
			CustomURL(mockServer.URL),
		)

		returnedPokemonSlice, err := myClient.GetAllPokemons(context.Background())

		expectedPokemonSlice := []Pokemon{
			{
				ID:             132,
				Name:           "ditto",
				BaseExperience: 101,
				Height:         3,
			},
			{
				ID:             5,
				Name:           "charmeleon",
				BaseExperience: 142,
				Height:         11,
			},
			{
				ID:             7,
				Name:           "squirtle",
				BaseExperience: 63,
				Height:         5,
			},
		}

		assert.NoError(t, err)

		assert.Equal(t, expectedPokemonSlice, returnedPokemonSlice)

	})

	t.Run("can handle 500 status code", func(*testing.T) {

		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				}),
		)

		defer mockServer.Close()

		myClient := NewClient(
			CustomURL(mockServer.URL),
		)

		_, err := myClient.GetAllPokemons(context.Background())

		assert.Error(t, err)

	})

	t.Run("can handle wrong json format", func(*testing.T) {

		mockServer := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, `["id": 7, "name": "squirtle", "base_experience": 63, "height": 5]`)
				}),
		)

		defer mockServer.Close()

		myClient := NewClient(
			CustomURL(mockServer.URL),
		)

		_, err := myClient.GetAllPokemons(context.Background())

		assert.Error(t, err)

	})
}
