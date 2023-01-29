package main

import "fmt"

type Arch struct {
	destination int
	brightness  int
}

type Graph struct {
	node         int
	destinations []Arch
}

func main() {

	n, m, h, s := 0, 0, 0, 0

	fmt.Println("Inserisci la prima riga: ")
	fmt.Scan(&n, &m, &h, &s)
	fmt.Println(n, m, h, s)
	// h harmony start
	// s sarah end

	for i := 0; i < m; i++ {
		
	}

}
