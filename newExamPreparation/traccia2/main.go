package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Oggetto struct {
	Nome string
	Val  int
	Dx   string
	Sx   string
	Op   string
	Type string
}

func readInput() map[string]*Oggetto {

	var mappa = make(map[string]*Oggetto)
	f, _ := os.Open("input.txt")
	b := bufio.NewScanner(f)

	for b.Scan() {
		split := strings.Split(b.Text(), " ")
		nome := strings.ReplaceAll(split[0], ":", "")
		if len(split) <= 2 {
			val, _ := strconv.Atoi(split[1])
			mappa[nome] = &Oggetto{
				Nome: nome,
				Val:  val,
				Dx:   "",
				Sx:   "",
				Op:   " ",
				Type: "num",
			}
		} else {
			mappa[nome] = &Oggetto{
				Nome: nome,
				Val:  0,
				Dx:   split[1],
				Sx:   split[3],
				Op:   split[2],
				Type: "Ogg",
			}
		}
	}
	return mappa
}

func stampaFigli(mappa map[string]*Oggetto, k string){
	if 
	
}

func main() {
	mappa := readInput()
	stampaFigli()
}
