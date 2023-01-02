/* Implementazione di grafo con struttura vertice

// Se ai vertici del grafo sono associate informazioni più articolate
// (che non possono essere usate come chiavi di una mappa),
// conviene usare una struttura vertice contenente anche una slice di puntatori ai vertici vicini.

(Questa rappresentazione ricorda quella degli alberi binari con una struttura nodo con puntatori ai figli sinistro e destro.)
*/

package main

import "fmt"

type item int
type chiave string

type vertice struct {
	valore item
	k      chiave
	adj    []*vertice // insieme dei vertici adiacenti
}

type graph map[chiave]*vertice

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
func leggiGrafo() (g graph) {
	//fmt.Println("Inserisci il numero di vertici")
	g = make(map[chiave]*vertice)

	i := 0
	var v1, v2 chiave
	_, err := fmt.Scan(&v1, &v2)

	for err == nil {
		//fmt.Println("letti vertici:", v1, v2)
		_, ok := g[v1]
		if !ok {
			g[v1] = &vertice{item(i), v1, nil}
			i++
		}
		_, ok = g[v2]
		if !ok {
			g[v2] = &vertice{item(i), v2, nil}
			i++
		}
		g[v1].adj = append(g[v1].adj, g[v2])
		//g[v2].adj = append(g[v2].adj,g[v1])
		_, err = fmt.Scan(&v1, &v2)
	}
	return g
}

func stampaVertice(v *vertice) {
	fmt.Print(v.k, ", ", v.valore, ", adiacenti: ")
	for _, el := range v.adj {
		fmt.Print(el.k, " ")
	}
	fmt.Println()

}

func stampaGrafo(g graph) {
	fmt.Printf("Il grafo ha %d nodi\n", len(g))
	for _, vertice := range g {
		stampaVertice(vertice)
	}
}

func dfs1(g graph, v *vertice, aux map[chiave]bool) {
	fmt.Println(v.k)
	aux[v.k] = true
	for _, v2 := range v.adj {
		_, ok := aux[v2.k]
		if !ok {
			dfs1(g, v2, aux)
		}
	}

}

func bfs1(g graph, v *vertice, aux map[chiave]bool) {
	coda := []*vertice{v}
	aux[v.k] = true

	for len(coda) > 0 {
		v := coda[0]
		coda = coda[1:]
		fmt.Println(v.k)

		for _, v2 := range v.adj {
			_, ok := aux[v2.k]
			if !ok {
				coda = append(coda, v2)
				aux[v2.k] = true
			}
		}
	}
}

func main() {
	g := leggiGrafo()
	stampaGrafo(g)

	aux := make(map[chiave]bool)
	v := g["uno"]
	dfs1(g, v, aux)
	aux = make(map[chiave]bool)
	v = g["due"]
	fmt.Println()
	bfs1(g, v, aux)
}
