package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"pokemon_cli/repository"
	"strconv"
	"strings"
	"time"
)

func main() {

	var name string

	// Init Repository
	now := time.Now()
	pokemonRepository := repository.NewPokemonRepository(now)
	userRepository := repository.NewUserRepository()

	fmt.Println("--Pokemon Game--")
	fmt.Print("Enter your name: ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	err := scanner.Err()

	if err != nil {
		log.Fatal("Error get your name :: ", err)
	}

	name = scanner.Text()
	userRepository.Name = name

	fmt.Println("Hello! ", scanner.Text())
	fmt.Println("--Menu--")
	fmt.Println("[1] Input 'all_pokemon' To get all available pokemons")
	fmt.Println("[2] Input 'catch_{pokemon_id}' Enter Pokémon ID to attempt to catch")
	fmt.Println("[3] Input 'get_collections' To get all your pokemon catch collections")
	fmt.Println("[4] Input 'exit' To exit the application")

	for {
		fmt.Print(name, "! >>> ")
		scanner.Scan()

		err := scanner.Err()
		if err != nil {
			fmt.Println("Error input :: ", err)
		}

		input := scanner.Text()

		input_arr := strings.Split(input, "_")

		// Show all available pokemons in console
		// if user input 'all_pokemon'
		if input == "all_pokemon" {
			pokemons := pokemonRepository.GetAll()

			fmt.Println("--Available Pokémon--")

			for _, pokemon := range pokemons {
				fmt.Printf("ID: %d, Name: %s, Type: %s, Catch Rate: %d%%, Is Rare: %t, Registered Date: %s\n",
					pokemon.ID, pokemon.Name, pokemon.Type, pokemon.CatchRate, pokemon.IsRare, pokemon.RegisteredDate.Format(time.RFC1123))
			}

		} else if input == "get_collections" {
			favPokemons := userRepository.GetAllPokemons()

			if favPokemons != nil {
				for _, pokemon := range favPokemons {
					fmt.Printf("ID: %d, Name: %s, Type: %s, Catch Rate: %d%%, Is Rare: %t, Registered Date: %s\n",
						pokemon.ID, pokemon.Name, pokemon.Type, pokemon.CatchRate, pokemon.IsRare, pokemon.RegisteredDate.Format(time.RFC1123))
				}
			} else {
				fmt.Println("You dont have any pokemons")
			}

		} else if input_arr[0] == "catch" && len(input_arr) > 1 {
			// Try catch the pokemon when prefix input 'catch'
			// Split the string and get the index 0
			// Convert type from string to id

			pokemonId := strings.TrimSpace(input_arr[1])

			if pokemonId == "" {
				fmt.Println("Pokemon id cannot empty or null!")
				continue
			}

			id, err := strconv.Atoi(pokemonId)

			if err != nil {
				fmt.Println("Pokemon id not valid")
				continue
			}

			fmt.Printf("Trying catch pokemon id: %s", pokemonId)
			for i := 0; i < 5; i++ {

				time.Sleep(500 * time.Millisecond)
				fmt.Print(" . ")
			}

			fmt.Println("")

			// Find pokemon and save to user collection pokemons
			pokemon, err := pokemonRepository.GetById(id)

			if err != nil {
				fmt.Println("Failed catch pokemon:", err.Error())
				continue
			}

			// Catch catchProbability
			result, ok := catchProbability(pokemon.CatchRate)

			fmt.Printf("You attempted to catch %s (%s type) with a catch rate of %d%%: %s\n",
				pokemon.Name, pokemon.Type, pokemon.CatchRate, result)

			if ok {

				userRepository.SavePokemon(pokemon)

				fmt.Printf("Success saved pokemon %s with id %d to your collections", pokemon.Name, pokemon.ID)
				fmt.Println("")

			} else {

				fmt.Println("Invalid Pokémon ID. Please try again.")

			}

		} else if input == "exit" {
			// Break the loop if user input "exit"
			break

		} else {
			fmt.Println("Your input is wrong, please try again")
			continue
		}
	}

	fmt.Println("Thankyou! ", name)
}

// catchProbability checks if catching is successful given a percentage rate
func catchProbability(rate int) (string, bool) {
	if rate < 0 || rate > 100 {
		return "Invalid rate. Please provide a rate between 0 and 100.", false
	}

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	chance := rand.Intn(100) + 1     // Generate a random number between 1 and 100

	if chance <= rate {
		return "SUCCESS, you caught it", true
	}
	return "FAIL, it got away", false
}
