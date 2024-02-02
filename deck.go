package main

import "fmt"

type deck []string

func (cards deck) print() {
	for _, card := range cards {
		fmt.Println(card)
	}
}