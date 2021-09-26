package main

import "fmt"

func main() {

	fmt.Println(insertionSort([]int{3, 5, 6, 3, 7, 8, 2, 4}))
	fmt.Println(bubbleSort([]int{3, 5, 6, 3, 7, 8, 2, 4}))
}

//based on: https://www.youtube.com/watch?v=vg4GS4LSJXI
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

//Based on: https://tutorialedge.net/courses/go-algorithms-course/21-bubble-sort-in-go/
func bubbleSort(input []int) []int {
	n := len(input)
	swapped := true
	for swapped {

		swapped = false

		for i := 0; i < n-1; i++ {

			if input[i] > input[i+1] {

				input[i], input[i+1] = input[i+1], input[i]

				swapped = true
			}
		}
	}
	return input
}
