package models

type Response struct {
	Results []SearchPokemon
}

type SearchPokemon struct {
	Name      string
	ID        int
	Url       string
	Weight    int
	Height    int
	Types     []Types
	Abilities []Abilities
}

type Types struct {
	Type Type
}

type Type struct {
	Name string
}

type Abilities struct {
	Ability Ability
}

type Ability struct {
	Name string
}
