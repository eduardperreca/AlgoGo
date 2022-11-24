package main

import (
	"fmt"
)

func quantiSenzaSubordinati(mappaDipendenti map[string][]string) int {
	var count int
	ignore := false
	for _, v := range mappaDipendenti {
		for _, x := range v {
			ignore = false
			if _, ok := mappaDipendenti[x]; ok {
				// fmt.Println(ok, " ", x)
				ignore = true
			}
			if !ignore {
				count++
			}

		}
	}
	return count
}

func stampaSubordinati(mappaDipendenti map[string][]string) {
	var lista []string
	for _, v := range mappaDipendenti {
		for _, x := range v {
			lista = append(lista, x)
		}
	}
	for _, k := range lista {
		fmt.Println(k)
	}
}

func supervisore(mappaDipendenti map[string][]string, name string) string {
	for k, v := range mappaDipendenti {
		for _, x := range v {
			if x == name {
				return k
			}
		}
	}
	return "not found"
}

// complessitá temporale O(n^2)

func recursiveTrap(mappaDipendenti map[string][]string, name string, lista []string) []string {
	lista = append(lista, name)
	for k, v := range mappaDipendenti {
		for _, x := range v {
			if x == name {
				lista = recursiveTrap(mappaDipendenti, k, lista)
			}
		}
	}
	return lista
}

func stampaImpiegatiSopra(mappaDipendenti map[string][]string, name string) {
	var lista []string
	lista = recursiveTrap(mappaDipendenti, name, lista)
	if len(lista) > 1 {
		for i := 1; i < len(lista); i++ {
			fmt.Println(lista[i])
		}
	} else {
		for _, x := range lista {
			fmt.Println(x)
		}
	}
}

func stampaImpiegatiConSupervisore(mappaDipendenti map[string][]string) {
	for _, v := range mappaDipendenti {
		for _, x := range v {
			fmt.Println(x)
		}
	}
}

func main() {

	var mappaDipendenti = make(map[string][]string)

	mappaDipendenti["A"] = []string{"B", "C", "D"}
	mappaDipendenti["B"] = []string{"E", "F"}
	mappaDipendenti["G"] = []string{"H"}

	for k, v := range mappaDipendenti {
		fmt.Println(k, "->", v)
		fmt.Println("Dipendenti di", k, ":", len(v))
		fmt.Println(v, "->", k)
	}

	test := quantiSenzaSubordinati(mappaDipendenti)
	fmt.Println(test)

	fmt.Println("stampo i subordinati kekw")
	stampaSubordinati(mappaDipendenti)

	fmt.Println("stampo il supervisore di H")
	test2 := supervisore(mappaDipendenti, "H")
	fmt.Println(test2)

	fmt.Println("stampo gli impiegati sopra di H")
	stampaImpiegatiSopra(mappaDipendenti, "F")

	fmt.Println("stampo gli impiegati con supervisore")
	stampaImpiegatiConSupervisore(mappaDipendenti)

}
