/* implementazione di grafi come vettore (slice) di insiemi (slice) di adiacenza

// L'insieme dei vertici (individuati dagli interi 0,1,... n-1) è definito implicitamente da una variabile intera n.
*/

package main

import (
	"fmt"
)

type adjSet []int

type graph struct {
	n   int
	adj []adjSet
}

func nuovoGrafo(n int) *graph {
	var g graph
	g.n = n
	g.adj = make([]adjSet, n)

	stampaGrafo(&g)
	return &g
}

/* legge da standard input un grafo */
// se ci sono errori restituisce false
// assumo che le mappe per vertici e archi siano già state create
//FORMATO
// n
// v1 w1
// v2 w2
// v3 w3
// ...
// vm wm
func leggiGrafo() (g *graph) {
	//fmt.Println("Inserisci il numero di vertici")
	var n int
	fmt.Scan(&n)
	g = nuovoGrafo(n)

	fmt.Println("Inserisci gli archi, uno per riga, nel formato v1 v2")
	var v1, v2 int
	_, err := fmt.Scan(&v1, &v2)

	for err == nil {
		//fmt.Println("letti vertici:", v1, v2)
		g.adj[v1] = append(g.adj[v1], v2)
		//g.adj[v2] = append(g.adj[v2], v1)
		//stampaGrafo(g)
		_, err = fmt.Scan(&v1, &v2)
	}
	return g
}

func stampaGrafo(g *graph) {
	fmt.Printf("Il grafo ha %d nodi\n", g.n)
	for i, adiacenti := range g.adj {
		fmt.Println(i, ": ", adiacenti)
	}
}

func dfs1(g *graph, v int, aux []bool) {
	fmt.Println(v)
	aux[v] = true
	for _, v2 := range g.adj[v] {
		if !aux[v2] {
			dfs1(g, v2, aux)
		}
	}

}

func bfs1(g *graph, v int, aux []bool) {
	coda := []int{v}
	aux[v] = true

	for len(coda) > 0 {
		v := coda[0] // nota: se, invece di togliere dall'inizio, togliessimo dalla fine, avremmo una pila, e quindi si avrebbe una dfs
		coda = coda[1:]
		fmt.Println("\t", v)

		for _, v2 := range g.adj[v] {
			if !aux[v2] {
				coda = append(coda, v2)
				aux[v2] = true
			}
		}
	}
}

func main() {
	g := leggiGrafo()
	stampaGrafo(g)

	aux := make([]bool, g.n)
	dfs1(g, 1, aux)
	aux = make([]bool, g.n)
	bfs1(g, 1, aux)
}
