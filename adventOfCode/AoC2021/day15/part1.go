/*

AOC 2021 Day 15 Part 1
Time complexity: O(shit)
Space complexity: ahahah lol

--
grafico implementato attraverso una matrice, creo una struttura "Node" per definirmi un nodo, con valore, riga e colonna, cosí da poter navigare nella matrice in maniera piú semplice.
per salvarmi le varie posizioni utilizzo una mappa, per poter avere un accesso diretto in memoria, salvando l'intero oggetto (avrei potuto salvare anche direttammente le coordinate).
implemento una coda di prioritá ma dato che non avevo assolutamente voglia di implementare un heap, ho fatto una coda normale e ogni volta che devo estrarre il minimo,
lo cerco tra tutti gli elementi della coda, roba assurda perché si sente tanto la differenza di complessitá di tempo essendo la ricerca O(n) invece di O(logn).
per ogni nodo estratto dalla coda, controllo se sono arrivato alla fine, se sì, stampo il valore e termino il programma, altrimenti controllo se posso andare in basso, in alto, a destra e a sinistra.
accodo tutti i nodi e i loro valori sommando il cammino con quello del nodo precedente.
quando ho finito di accodare tutti i nodi, li ordino in base al valore e estraggo il minimo, lo rimuovo dalla coda e ripeto il processo.
la funzoine returnPositionMin ritorna la posizione del minimo nella coda e il minimo stesso, cosí da poterlo rimuovere dalla coda, una brutta implementazione ma funzionale, sull'input della parte 1
funziona ma va moooolto lento.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Val int
	Row int
	Col int
}

func returnPositionMin(queue []Node) (int, Node) {
	min := queue[0]
	minRet := 0
	for i, k := range queue {
		if k.Val < min.Val {
			min = k
			minRet = i

		}
	}
	return minRet, min
}

func main() {

	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	matrix := [][]int{}

	for b.Scan() {
		row := []int{}
		for _, k := range b.Text() {
			numb, _ := strconv.Atoi(string(k))
			row = append(row, numb)
		}
		matrix = append(matrix, row)
	}

	queue := []Node{}
	visited := make(map[Node]bool)

	queue = append(queue, Node{0, 0, 0})

	visited[Node{0, 0, 0}] = true
	fmt.Println(queue)
	for len(queue) > 0 {
		fmt.Println(queue, "that's the queue before removing the min")
		x, min := returnPositionMin(queue)
		queue = append(queue[:x], queue[x+1:]...)
		fmt.Println(min, "min value")
		fmt.Println(queue, "that's the queue after removing the min")

		if min.Row == len(matrix)-1 && min.Col == len(matrix[0])-1 {
			fmt.Println(min.Val)
			break
		}
		fmt.Println("sono qui lol")

		if min.Row+1 < len(matrix) && !visited[Node{matrix[min.Row+1][min.Col], min.Row + 1, min.Col}] {
			fmt.Println("sono qui")
			queue = append(queue, Node{matrix[min.Row+1][min.Col] + min.Val, min.Row + 1, min.Col})
			visited[Node{matrix[min.Row+1][min.Col] + min.Val, min.Row + 1, min.Col}] = true
		}
		if min.Col+1 < len(matrix[0]) && !visited[Node{matrix[min.Row][min.Col+1], min.Row, min.Col + 1}] {
			queue = append(queue, Node{matrix[min.Row][min.Col+1] + min.Val, min.Row, min.Col + 1})
			visited[Node{matrix[min.Row][min.Col+1], min.Row, min.Col + 1}] = true
		}
		if min.Row-1 >= 0 && !visited[Node{matrix[min.Row-1][min.Col], min.Row - 1, min.Col}] {
			queue = append(queue, Node{matrix[min.Row-1][min.Col] + min.Val, min.Row - 1, min.Col})
			visited[Node{matrix[min.Row-1][min.Col], min.Row - 1, min.Col}] = true
		}
		if min.Col-1 >= 0 && !visited[Node{matrix[min.Row][min.Col-1], min.Row, min.Col - 1}] {
			queue = append(queue, Node{matrix[min.Row][min.Col-1] + min.Val, min.Row, min.Col - 1})
			visited[Node{matrix[min.Row][min.Col-1], min.Row, min.Col - 1}] = true
		}
		fmt.Println(queue)
	}
}
