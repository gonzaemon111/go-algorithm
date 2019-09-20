package main

import (
	"fmt"
	"math"
)

// 二分木を確認するための補助関数
// 二桁になると表示のバランスが崩れるのでdebug用
// 本関数を使用しない場合はmathパッケージをimportしなくてよい
func PrintBinaryTree(a []int) {
	fmt.Println("--------Binary tree begin--------")
	depth := int(math.Log2(float64(len(a))) + 1)
	cnt := 1
	for i, v := range a {
		for j := 0; j < int(math.Pow(2, float64(depth - cnt)) - 1); j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", v)
		for j := 0; j < int(math.Pow(2, float64(depth - cnt)) - 1); j++ {
			fmt.Printf(" ")
		}
		fmt.Printf(" ")
		
		if i == int(math.Pow(2, float64(cnt))) - 2 {
			fmt.Println("")
			fmt.Println("")
			cnt++
		}
	}
	
	fmt.Println("\n---------Binary tree end---------")
}

func UpHeap(a []int, n int) []int {
	a = append(a, n)
 	child := len(a) - 1
	var parent int = (child + 1) / 2 - 1

	for {
		if a[child] > a[parent] {
			a[child], a[parent] = a[parent], a[child]
			child = parent
			parent = (child + 1) / 2 - 1
			if parent < 0 {
				break
			}
		} else {
			break
		}
	}
	return a	
}


func DownHeap(a []int) []int {
	a[0], a[len(a) - 1] = a[len(a) - 1], a[0]
	a = a[:len(a)-1]
	
	parent := 0
	var child int
	
	for {
		child = 2 * parent + 1
		
		// 子どもがいなければ親が葉になるので終了
		if child > len(a) - 1 {
			break
		}
		
		// 親との比較は子どものうち大きい方とのみしたいので前処理
		if child != len(a) - 1 {
			if a[child] < a[child + 1] {
				child++
			}
		}
		
		if a[parent] >= a[child] {
			break
		} 
		
		a[parent], a[child] = a[child], a[parent]
		parent = child
	}
	
	return a
}

func HeapSort(a []int) []int {
	var b []int
	b = append(b, a[0])
	
	for i := 1; i < len(a); i++ {
		b = UpHeap(b, a[i])
	}

	for i := 0; i < len(a); i++ {
		a[len(a) - 1 - i] = b[0]
		b = DownHeap(b)
	}
	
	return a
}

func main()  {
	a := []int{2, 4, 5, 1, 3}
	// a := []int{2, 4, 5, 1, 3, 2, 1, 5, 1, 6, 8}

	fmt.Println(HeapSort(a))
}
