package main

import "fmt"

func main() {
	var s string = "Hola Mundo"
	var b string = "Chao Mundo"

	fmt.Println(s, b)
	fmt.Println(insertionSort([]int{3, 5, 6, 3, 7, 8, 2, 4}))
}

func insertionSort(A []int) []int {
	for j := 0; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i > -1 && A[i] > key {
			A[i+1] = A[i]
			i -= 1
		}
		A[i+1] = key
	}
	return A
}

func selectionSort(A []int) []int {

}
