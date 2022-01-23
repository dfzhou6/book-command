# topK
```go
package main

import "fmt"

const maxN = 10

var arr []int

func main() {
	// arr := []int{5, 2, 3, 4, 6, 9, 0, 1, 4, 7}
	arr = []int{5, 2, 3, 4, 6, 9, 0}
	heapify_build()
	fmt.Println(arr)
	heapifyAdd(1)
	fmt.Println(arr)
	heapifyAdd(4)
	fmt.Println(arr)
	heapifyAdd(7)
	fmt.Println(arr)
	heapifyAdd(10)
	fmt.Println(arr)
	heapifyAdd(0)
	fmt.Println(arr)

	for len(arr) > 0 {
		fmt.Println(heapifyDel())
	}
}

func heapify_build() {
	n := len(arr)
	for i := (n-2)/2; i >= 0; i-- {
		heapify(n, i)
	}
}

func heapify(n, i int) {
	max := i
	l, r := 2*i+1, 2*i+2
	if l <= n-1 && arr[max] < arr[l] {
		max = l
	}
	if r <= n-1 && arr[max] < arr[r] {
		max = r
	}
	if max != i {
		arr[max], arr[i] = arr[i], arr[max]
		heapify(n, max)
	}
}

func heapifyAdd(v int) {
	n := len(arr)
	if n < maxN {
		arr = append(arr, v)
		n++
		for i := n-1; (i-1)/2 >= 0; {
			p := (i-1)/2
			if arr[i] > arr[p] {
				arr[i], arr[p] = arr[p], arr[i]
				i = p
			} else {
				break
			}
		}
	} else {
		heapifyRpl(v)
	}
}

func heapifyDel() int {
	n := len(arr)
	v := arr[0]
	arr[0], arr[n-1] = arr[n-1], arr[0]
	arr = arr[:n-1]
	heapify(n-1, 0)
	return v
}

func heapifyRpl(v int) {
	if v < arr[0] {
		arr[0] = v
		heapify(len(arr), 0)
	}
}

```