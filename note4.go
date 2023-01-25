// insertion sort

package main

import "fmt"

func insertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	return array
}

func main() {
	array := []int{1, 4, 12, 8, 6, 32, 9}
	fmt.Println(array)
	fmt.Println(insertionSort(array))

}


// ---------------------------------------------- //

// mergesort

package main

import "fmt"

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	fmt.Println(left, right)

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func mergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}

	middle := len(array) / 2
	left := mergeSort(array[:middle])
	right := mergeSort(array[middle:])
	return merge(left, right)
}

func main() {
	
	array := []int{1, 4, 12, 8, 6, 32, 9}
	fmt.Println(array)
	fmt.Println(mergeSort(array))

}

// ---------------------------------------------- //

// quicksort

package main

import "fmt"

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	pivot := array[0]
	left := make([]int, 0, len(array))
	right := make([]int, 0, len(array))

	for _, v := range array[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func main(){

	arr := []int{1, 12, 4, 8, 6, 32, 9}
	fmt.Println(arr)
	t := quickSort(arr)
	fmt.Println(t)
}

// ---------------------------------------------- //

// selection sort

package main

import "fmt"

func selectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
			fmt.Println(array, i, j)
		}
	}
	return array
}

func main() {

	array := []int{1, 4, 12, 8, 6, 32, 9}
	fmt.Println(array)
	fmt.Println(selectionSort(array))
}

// ---------------------------------------------- //

// sortin pazzo

package main

import "fmt"

func main() {

	arr := []int{1, 4, 12, 8, 6, 32, 9}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

// ---------------------------------------------- //