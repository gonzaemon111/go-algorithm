package main

import (
	"fmt"
)

// ソート済みの配列をマージする
func Merge(a, b []int) []int {
	result := make([]int, len(a) + len(b))

	var i, j, cnt int
	for i + j < len(a) + len(b){
		if a[i] < b[j] {
			result[cnt] = a[i]
			i++
		} else {
			result[cnt] = b[j]			
			j++
		}
		cnt++
		
		if i == len(a) {
			for j < len(b) {
				result[cnt] = b[j]
				cnt++
				j++
			}
			break
		}

		if j == len(b) {
			for i < len(a) {
				result[cnt] = a[i]
				cnt++
				i++
			}
			break
		}
	}
	
	return result
}

func DivideArray(a []int) ([]int, []int) {
	return a[:len(a) / 2], a[len(a) / 2:]
}

func MergeSort(a []int) []int {
	a1, a2 := DivideArray(a)

	if len(a1) > 1 {
		a1 = MergeSort(a1)
	}
	if len(a2) > 1{
		a2 = MergeSort(a2)
	}
	a = Merge(a1, a2)
	return a
}

func main()  {
	a := []int{2, 4, 5, 1, 3}
	fmt.Println(MergeSort(a))
}
