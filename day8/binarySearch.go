package main

func binarySearchRecursive(arr []int, value int) int {
	if len(arr) == 0 {
		return -1
	}
	mid := len(arr) / 2
	if arr[mid] == value {
		return mid
	}
	if arr[mid] > value {
		return binarySearchRecursive(arr[:mid], value)
	}
	return binarySearchRecursive(arr[mid+1:], value)
}

func main(){

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
}