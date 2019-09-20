package main

import "fmt"

func BubbleSort(a []int) []int {
    for i := 0; i < len(a) - 1; i++ {
        for j := 0; j < len(a) - i - 1; j++ {
            if a[j] > a[j + 1] {
                a[j], a[j + 1] = a[j + 1], a[j]
            }
        }
    }
    return a
}


func main()  {
	a := []int{2, 4, 5, 1, 3}

	fmt.Println(BubbleSort(a))
}
