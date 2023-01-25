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


// --------------------------------------------- //


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


// --------------------------------------------- //


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




// --------------------------------------------- //


package main

import "fmt"

type Node struct {
	Row    int
	Column int
}

func main() {

	matrix := [][]int{
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 0, 0, 2},
	}

	// 1. Create a stack
	stack := []Node{}

	// 2. Add the first node to the stack
	stack = append(stack, Node{0, 0})

	// 3. Create visited map
	visited := make(map[Node]bool)

	// 4. While the stack is not empty
	for len(stack) > 0 {

		// 5. Pop the first node
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(node)
		fmt.Println("Stack:", stack)
		fmt.Println("Visited:", visited)

		// 6. Check if it is the goal
		if matrix[node.Row][node.Column] == 2 {
			fmt.Println("END!, we spent", len(visited), "steps")
			return
		}

		// 7. Add the neighbors to the stack
		if node.Row > 0 {
			if _, ok := visited[Node{node.Row - 1, node.Column}]; !ok && matrix[node.Row-1][node.Column] == 0 || matrix[node.Row-1][node.Column] == 2 {
				visited[Node{node.Row - 1, node.Column}] = true
				stack = append(stack, Node{node.Row - 1, node.Column})
			}
		}

		if node.Column > 0 {
			if _, ok := visited[Node{node.Row, node.Column - 1}]; !ok && matrix[node.Row][node.Column-1] == 0 || matrix[node.Row][node.Column-1] == 2 {
				visited[Node{node.Row, node.Column - 1}] = true
				stack = append(stack, Node{node.Row, node.Column - 1})
			}
		}
		if node.Column < len(matrix[0])-1 {
			if _, ok := visited[Node{node.Row, node.Column + 1}]; !ok && matrix[node.Row][node.Column+1] == 0 || matrix[node.Row][node.Column+1] == 2 {
				visited[Node{node.Row, node.Column + 1}] = true
				stack = append(stack, Node{node.Row, node.Column + 1})
			}
		}
		if node.Row < len(matrix)-1 {
			if _, ok := visited[Node{node.Row + 1, node.Column}]; !ok && matrix[node.Row+1][node.Column] == 0 || matrix[node.Row+1][node.Column] == 2 {
				visited[Node{node.Row + 1, node.Column}] = true
				stack = append(stack, Node{node.Row + 1, node.Column})
			}
		}

	}

}



// --------------------------------------------- //


// bfs 

package main

import "fmt"

type Node struct {
	Row    int
	Column int
}

func main() {

	matrix := [][]int{
		{7, 1, 8, 4, 4},
		{9, 6, 22, 4, 4},
		{1, 1, 8, 4, 4},
		{1, 9, 8, 4, 4},
		{1, 1, 1, 1, 12},
	}

	// 1. Create a queue
	queue := []Node{}

	// 2. Add the first node to the queue
	queue = append(queue, Node{0, 0})
	visited := make(map[Node]bool)
	// 3. While the queue is not empty
	for len(queue) > 0 {

		// 4. Dequeue the first node
		node := queue[0]
		queue = queue[1:]

		fmt.Println("visiting", node)
		fmt.Println("with value: ", matrix[node.Row][node.Column])
		fmt.Println("visited", visited)

		// 5. Check if it is the goal
		if matrix[node.Row][node.Column] == 6 {
			println("Found!")
			return
		}

		// 6. Add the neighbors to the queue
		if node.Row > 0 {
			if _, ok := visited[Node{node.Row - 1, node.Column}]; !ok {
				visited[Node{node.Row - 1, node.Column}] = true
				queue = append(queue, Node{node.Row - 1, node.Column})
			}

		}
		if node.Row < len(matrix)-1 {
			if _, ok := visited[Node{node.Row + 1, node.Column}]; !ok {
				visited[Node{node.Row + 1, node.Column}] = true
				queue = append(queue, Node{node.Row + 1, node.Column})
			}
		}
		if node.Column > 0 {
			if _, ok := visited[Node{node.Row, node.Column - 1}]; !ok {
				visited[Node{node.Row, node.Column - 1}] = true
				queue = append(queue, Node{node.Row, node.Column - 1})
			}
		}
		if node.Column < len(matrix[0])-1 {
			if _, ok := visited[Node{node.Row, node.Column + 1}]; !ok {
				visited[Node{node.Row, node.Column + 1}] = true
				queue = append(queue, Node{node.Row, node.Column + 1})
			}
		}
		fmt.Println("queue", queue)
	}

}


// --------------------------------------------- //

// dfs


package main

import "fmt"

type Node struct {
	Row    int
	Column int
}

func main() {

	matrix := [][]int{
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 1, 0, 0, 2},
	}

	// 1. Create a stack
	stack := []Node{}

	// 2. Add the first node to the stack
	stack = append(stack, Node{0, 0})

	// 3. Create visited map
	visited := make(map[Node]bool)

	// 4. While the stack is not empty
	for len(stack) > 0 {

		// 5. Pop the first node
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(node)
		fmt.Println("Stack:", stack)
		fmt.Println("Visited:", visited)

		// 6. Check if it is the goal
		if matrix[node.Row][node.Column] == 2 {
			fmt.Println("END!, we spent", len(visited), "steps")
			return
		}

		// 7. Add the neighbors to the stack
		if node.Row > 0 {
			if _, ok := visited[Node{node.Row - 1, node.Column}]; !ok && matrix[node.Row-1][node.Column] == 0 || matrix[node.Row-1][node.Column] == 2 {
				visited[Node{node.Row - 1, node.Column}] = true
				stack = append(stack, Node{node.Row - 1, node.Column})
			}
		}

		if node.Column > 0 {
			if _, ok := visited[Node{node.Row, node.Column - 1}]; !ok && matrix[node.Row][node.Column-1] == 0 || matrix[node.Row][node.Column-1] == 2 {
				visited[Node{node.Row, node.Column - 1}] = true
				stack = append(stack, Node{node.Row, node.Column - 1})
			}
		}
		if node.Column < len(matrix[0])-1 {
			if _, ok := visited[Node{node.Row, node.Column + 1}]; !ok && matrix[node.Row][node.Column+1] == 0 || matrix[node.Row][node.Column+1] == 2 {
				visited[Node{node.Row, node.Column + 1}] = true
				stack = append(stack, Node{node.Row, node.Column + 1})
			}
		}
		if node.Row < len(matrix)-1 {
			if _, ok := visited[Node{node.Row + 1, node.Column}]; !ok && matrix[node.Row+1][node.Column] == 0 || matrix[node.Row+1][node.Column] == 2 {
				visited[Node{node.Row + 1, node.Column}] = true
				stack = append(stack, Node{node.Row + 1, node.Column})
			}
		}

	}

}


// --------------------------------------------- //


// shortest path with strange prio-queue lol

package main

import "fmt"

type Node struct {
	Row    int
	Column int
	Val    int
}

func min(a []Node, target int) Node {
	min := a[0].Val
	for _, v := range a {
		if v.Val == target {
			return v
		}
	}

	for _, v := range a {
		if v.Val < min {
			min = v.Val
		}
	}
	for _, v := range a {
		if v.Val == min {
			return v
		}
	}
	return Node{}
}

func main() {
	matrix := [][]int{
		{7, 1, 8, 4, 4},
		{9, 1, 8, 4, 4},
		{1, 1, 8, 4, 4},
		{1, 9, 8, 4, 4},
		{1, 1, 1, 1, 1},
	}

	// 1. Create a stack
	stack := []Node{}

	// 2. Add the first node to the stack
	stack = append(stack, Node{0, 0, 0})
	dummy := []Node{}
	// 3. Create visited map
	visited := make(map[Node]bool)
	count := 0
	valSum := 0
	// 4. While the stack is not empty
	for len(stack) > 0 {

		// 5. Pop the first node
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(node)
		fmt.Println("value:", matrix[node.Row][node.Column])

		// 6. Check if it is the goal
		if node.Column == len(matrix)-1 && node.Row == len(matrix[0])-1 {
			fmt.Println("END!, we spent", count, "steps")
			fmt.Println("END!, we spent", valSum, "value")
			return
		}
		dummy = []Node{}
		// 7. Add the neighbors to the stack
		if node.Row > 0 {
			if _, ok := visited[Node{node.Row - 1, node.Column, matrix[node.Row-1][node.Column]}]; !ok {
				visited[Node{node.Row - 1, node.Column, matrix[node.Row-1][node.Column]}] = true
				dummy = append(dummy, Node{node.Row - 1, node.Column, matrix[node.Row-1][node.Column]})
			}
		}

		if node.Column > 0 {
			if _, ok := visited[Node{node.Row, node.Column - 1, matrix[node.Row][node.Column-1]}]; !ok {
				visited[Node{node.Row, node.Column - 1, matrix[node.Row][node.Column-1]}] = true
				dummy = append(dummy, Node{node.Row, node.Column - 1, matrix[node.Row][node.Column-1]})
			}
		}
		if node.Column < len(matrix[0])-1 {
			if _, ok := visited[Node{node.Row, node.Column + 1, matrix[node.Row][node.Column+1]}]; !ok {
				visited[Node{node.Row, node.Column + 1, matrix[node.Row][node.Column+1]}] = true
				dummy = append(dummy, Node{node.Row, node.Column + 1, matrix[node.Row][node.Column+1]})
			}
		}
		if node.Row < len(matrix)-1 {
			if _, ok := visited[Node{node.Row + 1, node.Column, matrix[node.Row+1][node.Column]}]; !ok {
				visited[Node{node.Row + 1, node.Column, matrix[node.Row+1][node.Column]}] = true
				dummy = append(dummy, Node{node.Row + 1, node.Column, matrix[node.Row+1][node.Column]})
			}
		}

		stack = append(stack, min(dummy, matrix[len(matrix)-1][len(matrix[0])-1]))
		fmt.Println("dummy:", dummy)
		count++
		valSum += stack[0].Val

	}
}
