// boardgame
package main

type Category struct {
	name        string
	description string
}

type Mechanism struct {
	name        string
	description string
}

type Range struct {
	min uint16
	max uint16
}

type game struct {
	primary_name string
	yearReleased uint16
	noPlayers    Range
	playTime     Range
	age          string
	categories   []Category
	mechanisms   []Mechanism
}
