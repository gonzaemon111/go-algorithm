# 基本的なソートアルゴリズムをGolangで実装してみた

## 背景

ソートアルゴリズムの話をするときに、自分の中で整理できていなかったので以下のアルゴリズムを実装してみました。

- バブルソート
- 選択ソート
- 挿入ソート
- シェルソート
- マージソート
- クイックソート
- ヒープソート

Goで書いたのは筋トレも兼ねているだけです。

そもそもGoにはsortパッケージ があります。
ちなみにリンク先を見て頂ければわかりますが、sortパッケージはクイックソートで実装されています。
(実体としては、データ長に応じて、ヒープソートと挿入ソートを併用する実装になっています)

ある程度ソート済みのアルゴリズムでは挿入ソートのほうが高速なのでソートしたいデータの特性によってはsortパッケージ以外を使用する選択肢もありそうです。

### バブルソート

隣同士比較して交換するだけです.

```go
package main

import (
	"fmt"
)

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
```

### 選択ソート

最小値(あるいは最大値)を「選択」し、小さい順に並べていくソートです。


```go
package main

import (
	"fmt"
)

func Min(a []int) (idx, n int) {
	n = a[0]
	idx = 0
	for i, v := range a {
		if n > v {
			n = v
			idx = i
		}
	}

	return
}

func SelectionSort(a []int) []int {
	for i, _ := range a {
		idx, _ := Min(a[i:len(a)])
		a[i], a[i + idx] = a[i + idx], a[i]
	}

	return a
}

func main()  {
	a := []int{2, 4, 5, 1, 3}
	fmt.Println(SelectionSort(a))
}
```

### 挿入ソート

整列済みのデータに対して、要素を正しい位置に挿入します。 整列済みであることを意識して、比較は後ろから行うようにしてみました。

連結リストで実装するのがよかったですね......。

```go
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
```

### シェルソート

適当な間隔hiずつ空けて、グループを作り、グループ内でソートします。さらに間隔を狭めていき、間隔が1になったところで全体をソートします。

上記の過程である程度ソートされていき、最後は挿入ソートで締めます。

Knuth先生のThe Art of Computer Programmingによると以下を満たすよう間隔を選んでいくとO(n1.25)になるそうです。
(未読ですが......)


hi+1=3hi+1
h0=1,hi∈


```go
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
```

### マージソート

データを適当に分けていき、それをマージする際にソートします。

```go
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
```

### クイックソート

一般なデータに対しては最も高速といわれているそうです。 ピボットより小さい値を持つグループと大きい値を持つグループにわけ、再帰的にソートしていきます。

ピボットの選び方によっては無限ループに入ってしまったり、計算量が増えてしまうので注意が必要です。

2017/1/14 訂正（詳細は記事冒頭に記載）


```go

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
```

### ヒープソート

ヒープ木では根が最大(または最小)になるので、根を取り除き、残った部分をさらにヒープ木にすることを繰り返します。

UpHeapとDownHeapをまとめてヒープ木を作る関数にするのもありかと思ったのですが、それだと必要以上に計算が多くなるので微妙なんでしょうか。


```go

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
```

---

全体的にGoの配列(slice)の操作に苦戦しました。 複数の値をまとめて代入できなかったり、配列を添え字で抜き出すところとか自分の想像する挙動と微妙に異なるので境界は特に気を遣いながら書きました。

上記以外にも様々なソートアルゴリズムがあるので実装していきたいです。 特にデータの特性に応じて、一定の条件下ではより高速になるようなソートはおもしろそうですね。



