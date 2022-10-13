// func largest(numbers []int) int {
// 	n := len(numbers)
// 	if 0 {
// 		return numbers[0]
// 	}
// 	return max(numbers[n-1], largest(0))
// }

package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func largest(s []int) int {

	if len(s) == 1 {
		return s[0]
	}

	return max(s[len(s)-1], largest(s[:len(s)-1]))

}

func main() {
	slice := []int{ 1, 2, 5, 7, -2, 10, 9, 21, 3, 8 }

	t := largest(slice)
	fmt.Println(t)

}
