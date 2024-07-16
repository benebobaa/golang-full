package repository

import (
	"fmt"
	"pokemon_cli/model"
	"time"
)

type PokemonRepository struct {
	Now      time.Time
	Pokemons []model.Pokemon
}

func NewPokemonRepository(now time.Time) *PokemonRepository {
	pokemons := []model.Pokemon{
		{ID: 1, Name: "Pikachu", Type: "Electric", CatchRate: 70, IsRare: false, RegisteredDate: now},
		{ID: 2, Name: "Charmander", Type: "Fire", CatchRate: 40, IsRare: true, RegisteredDate: now.AddDate(0, -1, 0)},
		{ID: 3, Name: "Bulbasaur", Type: "Grass/Poison", CatchRate: 10, IsRare: true, RegisteredDate: now.AddDate(0, -6, 0)},
		{ID: 4, Name: "Dragonite", Type: "Dragon/Flying", CatchRate: 30, IsRare: true, RegisteredDate: now.AddDate(0, -8, 0)},
		{ID: 5, Name: "Mew", Type: "Psychic", CatchRate: 1, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
	}

	return &PokemonRepository{
		Now:      now,
		Pokemons: pokemons,
	}
}

func (p *PokemonRepository) GetAll() []model.Pokemon {

	return p.Pokemons
}

func (p *PokemonRepository) GetById(id int) (*model.Pokemon, error) {

	// Search pokemon by id, iterate sequential
	for _, pokemon := range p.Pokemons {
		if pokemon.ID == id {
			return &pokemon, nil
		}
	}

	return nil, fmt.Errorf("Pokemon with id: %d not found", id)
}
