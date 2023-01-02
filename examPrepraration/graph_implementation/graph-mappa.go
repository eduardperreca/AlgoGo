/* Implementazione di grafo con una mappa delle adiacenze.

Nel caso in cui i vertici sono identificati da stringhe,
il grafo si può implementare con una mappa che associa ad ogni vertice (identificato da una stringa) la slice dei suoi vicini (identificati da una stringa)
l'insieme dei vertici è rappresentato implicitamente dall'insieme delle chiavi della mappa.
NB: se ci sono vertici senza vicini, servirà aggiungerli nella mappa con una slice vuota come slice dei vicini!

La soluzione si può adattare anche al caso in cui i vertici sono identificati da strutture semplici, ad esempio una coppia di interi.
(ma non vanno bene strutture che contengono per esempio delle slice, poiché questo tipo non può essere usato come chiave per una mappa).
*/

package main

import (
	"fmt"
)

func main() {
	g := leggiGrafo()
	stampaGrafo(g)
	aux := make(map[string]bool)
	dfs1(g, "uno", aux)
	aux = make(map[string]bool)
	bfs1(g, "uno", aux)
}

// mappa che associa ad ogni vertice (identificato da una stringa)
// l'insieme dei vertici (stringhe) a lui adiacenti
type grafo map[string][]string

func nuovoGrafo(n int) grafo {
	g := make(map[string][]string)
	return g
}

/* legge da standard input un grafo */
// se ci sono errori restituisce false
//FORMATO
// v1 w1
// v2 w2
// v3 w3
// ...
// vm wm
// NB: in questo formato non sono contemplati vertici senza vicini!
// (se si vuole rappresentare un grafo che ha anche nodi isolati, bisogna cambiare il formato di input in modo da poter specificare l'esistenza di tali vertici)
func leggiGrafo() grafo {
	g := make(map[string][]string)
	fmt.Println("Inserisci gli archi, uno per riga, nel formato v1 v2")
	var v1, v2 string
	_, err := fmt.Scan(&v1, &v2)

	for err == nil {
		//fmt.Println("letti vertici:", v1, v2)
		g[v1] = append(g[v1], v2)
		//g[v2] = append(g[v2], v1)
		_, err = fmt.Scan(&v1, &v2)
	}
	return g
}

func stampaGrafo(g grafo) {
	fmt.Printf("Il grafo ha %d nodi\n", len(g))
	for vertice, lista := range g {
		fmt.Println(vertice, ":", lista)
	}
}

func dfs1(g grafo, v string, aux map[string]bool) {
	fmt.Println(v)
	aux[v] = true
	for _, v2 := range g[v] {
		if !aux[v2] {
			dfs1(g, v2, aux)
		}
	}

}

func bfs1(g grafo, v string, aux map[string]bool) {
	coda := []string{v}
	aux[v] = true

	for len(coda) > 0 {
		v := coda[0]
		coda = coda[1:]
		fmt.Println("\t", v)

		for _, v2 := range g[v] {
			if !aux[v2] {
				coda = append(coda, v2)
				aux[v2] = true
			}
		}
	}
}
