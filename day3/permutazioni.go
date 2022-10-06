package main

import "fmt"

type Permutazioni struct {
	chiave int
	nome   string
}

func main() {

	// Creo un array di struct
	permutazioni := []Permutazioni{
		{chiave: 6, nome: "Francesco"},
		{chiave: 1, nome: "Andrea"},
		{chiave: 5, nome: "Elisa"},
		{chiave: 2, nome: "Beatrice"},
		{chiave: 3, nome: "Carlo"},
		{chiave: 4, nome: "Dino"},
		{chiave: 7, nome: "Giorgia"},
		{chiave: 9, nome: "Irene"},
		{chiave: 8, nome: "Henry"},
	}

	fmt.Println("Array di struct non ordinato: ", permutazioni)
	for i := 0; i < len(permutazioni); i++ {
		chiave := permutazioni[i].chiave
		permutazioni[i], permutazioni[len(permutazioni)-chiave] = permutazioni[len(permutazioni)-chiave], permutazioni[i]
	}

	fmt.Println("Array di struct ordinato: ", permutazioni)

}
