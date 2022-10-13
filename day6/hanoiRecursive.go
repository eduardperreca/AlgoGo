/*
Scrivete una funzione hanoi( n , from, temp, to int ) che stampi le
mosse per spostare n dischi dal paletto from al paletto to aiutandosi,
se necessario, con il paletto ausiliario temp.
*/

package main

import "fmt"

func hanoi(n int, from, temp, to string){
	if n == 1 {
		fmt.Println("sposta il disco da", from, "a", to)
	} else {
		hanoi(n-1, from, to, temp)
		fmt.Println("sposta il disco da", from, "a", to)
		hanoi(n-1, temp, from, to)
	}
}

func main(){
	hanoi(4, "A", "B", "C")

}

//ep 6