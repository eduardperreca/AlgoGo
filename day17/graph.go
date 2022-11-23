package main

type Vertex struct {
	Value int
	key  int
}

type Graph struct {
	Vertices map[int]*Vertex
	Edges    map[*Vertex][]*Vertex
}

func main(){


	
}