package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	client "github.com/codescalersinternships/PokeAPI-Client-Abdelrahman-Mahmoud/client"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	fmt.Println("Client created")

	myClient := client.NewClient()

	fmt.Println("GET /pokemon/ditto  -> ditto infotmation")
	returnedPokemon, err := myClient.GetPokemonByName(context.Background(), "ditto")

	if err != nil {
		log.Fatalf("error getting pokemon by name: %s", err)
	} else {
		fmt.Println(returnedPokemon)
	}

	fmt.Println("GET /pokemon/charmeleon  -> return charmeleon infotmation")
	returnedPokemon, err = myClient.GetPokemonByName(context.Background(), "charmeleon")

	if err != nil {
		log.Fatalf("error getting pokemon by name: %s", err)
	} else {
		fmt.Println(returnedPokemon)
	}

	fmt.Println("GET /pokemon/squirtle  -> return squirtle infotmation")
	returnedPokemon, err = myClient.GetPokemonByName(context.Background(), "squirtle")

	if err != nil {
		log.Fatalf("error getting pokemon by name: %s", err)
	} else {
		fmt.Println(returnedPokemon)
	}

	fmt.Println("GET /pokemon  -> return all pokemons infotmation")
	returnedPokemons, err := myClient.GetAllPokemons(context.Background())

	if err != nil {
		log.Fatalf("error getting all pokemons: %s", err)
	} else {
		fmt.Println(returnedPokemons)
	}

}
