package main

import (
	"fmt"
)

// 中間値を返す
func Med3(x, y, z int) int {
	if x < y {
		if y < z {
			return y
		} else if x < z {
			return z
		} else {
			return x
		}
	} else {
		if x < z {
			return x
		} else if y < z {
			return z
		} else {
			return y
		}
	}
}

func QuickSort(a []int) {
	pivot := Med3(a[0], a[len(a) / 2], a[len(a) - 1])
	left := 0
	right := len(a) - 1
	for {
		// 交換する対象を探す
		for a[left] < pivot {
			left++
		}
	
		for a[right] > pivot {
			right--
		}
		
		// 左右からの探索が交差したら終了
		if left >= right {
			break
		}
		
		// 対象を交換
		a[left], a[right] = a[right], a[left]

		flg := true
		if a[right] == pivot {
			left++
			flg = false
		}
		if a[left] == pivot && flg {
			right--
		}

	}
	
	a1 := a[:left]
	if len(a1) > 1 {
		QuickSort(a1)
	}

	a2 := a[right+1:]
	if len(a2) > 1 {
		QuickSort(a2)
	}
	
	cnt := 0
	for _, v := range a1 {
		a[cnt] = v
		cnt++
	}
	a[cnt] = pivot
	cnt++
	for _, v := range a2 {
		a[cnt] = v
		cnt++
	}
	
}

func main()  {
	a := []int{2, 4, 5, 1, 3}
	// a := []int{1, 0, 2}
	// a := []int{2, 4, 1, 3, 1}
	QuickSort(a)
	fmt.Println(a)
}
