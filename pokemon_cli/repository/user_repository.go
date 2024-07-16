package repository

import (
	"pokemon_cli/model"
)

type UserRepository struct {
	Name     string
	Pokemons []model.Pokemon
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) GetAllPokemons() []model.Pokemon {
	return u.Pokemons
}

func (u *UserRepository) SavePokemon(pokemon *model.Pokemon) {
	u.Pokemons = append(u.Pokemons, *pokemon)
}
