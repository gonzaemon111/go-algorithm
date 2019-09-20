package main

import (
	"fmt"
)

func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i - j - 1] > a[i - j] {
				a[i - j - 1], a[i - j] = a[i - j], a[i - j - 1]
			} else {
				break
			}
		}
	}

	return a
}

func main()  {
	a := []int{2, 4, 5, 1, 3}
	fmt.Println(InsertionSort(a))
}
