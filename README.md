# PokeAPI-Client-Abdelrahman-Mahmoud

## Introduction

A client-server denotes a relationship between cooperating programs in an application, composed of clients initiating requests for services and servers providing that function or service. In this project we use the previously made date time server to get pokemon information.

## Setup

1. Clone the Repository to a directory of your choice.
2. Make sure you have go version 1.22.4 installed on your device
3. Create demo.go file inside the working directory
4. import the package using 
   ```GO
	  import "github.com/codescalersinternships/PokeAPI-Client-Abdelrahman-Mahmoud"
   ```
5. Finish writing your desired code 
6. Open terminal
7. Build the project using
   ```console
   user@user-VirtualBox:~$ go build demo.go
   ```
8. Run the project using
   ```console
   user@user-VirtualBox:~$ ./demo
   ```

## Demo
- Code:
```GO
mt.Println("Client created")

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
```

## Tests

=== RUN   TestOptionFunctions
=== RUN   TestOptionFunctions/happy_path_-_can_add_custom_URLS_using_option_function
=== RUN   TestOptionFunctions/happy_path_-_can_add_custom_client_using_option_function
=== RUN   TestOptionFunctions/happy_path_-_can_add_custom_client_using_option_function#01
--- PASS: TestOptionFunctions (0.01s)
    --- PASS: TestOptionFunctions/happy_path_-_can_add_custom_URLS_using_option_function (0.00s)
    --- PASS: TestOptionFunctions/happy_path_-_can_add_custom_client_using_option_function (0.00s)
    --- PASS: TestOptionFunctions/happy_path_-_can_add_custom_client_using_option_function#01 (0.00s)
=== RUN   TestClientGetPokemonByNameCanHitMockServer
=== RUN   TestClientGetPokemonByNameCanHitMockServer/can_hit_the_mockserver_and_return_pokemon_information
=== RUN   TestClientGetPokemonByNameCanHitMockServer/can_handle_500_status_code
=== RUN   TestClientGetPokemonByNameCanHitMockServer/can_handle_wrong_json_format
--- PASS: TestClientGetPokemonByNameCanHitMockServer (0.01s)
    --- PASS: TestClientGetPokemonByNameCanHitMockServer/can_hit_the_mockserver_and_return_pokemon_information (0.00s)
    --- PASS: TestClientGetPokemonByNameCanHitMockServer/can_handle_500_status_code (0.00s)
    --- PASS: TestClientGetPokemonByNameCanHitMockServer/can_handle_wrong_json_format (0.00s)
=== RUN   TestClientGetAllPokemonsCanHitMockServer
=== RUN   TestClientGetAllPokemonsCanHitMockServer/can_hit_the_mockserver_and_return_all_pokemon_information
=== RUN   TestClientGetAllPokemonsCanHitMockServer/can_handle_500_status_code
=== RUN   TestClientGetAllPokemonsCanHitMockServer/can_handle_wrong_json_format
--- PASS: TestClientGetAllPokemonsCanHitMockServer (0.01s)
    --- PASS: TestClientGetAllPokemonsCanHitMockServer/can_hit_the_mockserver_and_return_all_pokemon_information (0.00s)
    --- PASS: TestClientGetAllPokemonsCanHitMockServer/can_handle_500_status_code (0.00s)
    --- PASS: TestClientGetAllPokemonsCanHitMockServer/can_handle_wrong_json_format (0.00s)
PASS
ok      github.com/codescalersinternships/PokeAPI-Client-Abdelrahman-Mahmoud/client     0.033s
