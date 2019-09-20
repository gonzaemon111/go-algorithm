package main

import (
	"fmt"
)

func CalcInterval(n int) int {
	h := 1
	
	for h < n {
		h = 3 * h + 1
	}
	
	h = (h - 1) / 3
	
	return h
}

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

func ShellSort(a []int) []int {
	h := CalcInterval(len(a))
	
	for h > 1 {
		for i := 0; i < h; i++ {
			// hずつ飛ばしたグループを作る
			b := make([]int, len(a) / h + 1)
			cnt := 0
			for j := 0; j < len(a) / h + 1; j++ {
				if i + h * j < len(a){
					b[j] = a[i + h * j]
					cnt++
				}
			}
			
			// 抜き出したグループに対して挿入ソート
			c := InsertionSort(b[:cnt])
			fmt.Println(c)
			
			// ソート済みのものを代入
			for j := 0; j < len(c); j++ {
				if i + h * j < len(a){
					a[i + h * j] = c[j]
				}
			}
			
		}
		h = (h - 1) / 3
	}
	
	a = InsertionSort(a)
	return a
}

func main()  {
	// a := []int{2, 4, 5, 1, 3}
	a := []int{2, 4, 5, 1, 3, 4, 5, 10, 11, 15, 1, 4, 2}
	fmt.Println(ShellSort(a))
}
