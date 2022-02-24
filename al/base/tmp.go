package main

import "fmt"

func main() {
	i := run_right([]int{1,3,5,6,6,6,8,11,19,30,50,100}, 4)
	fmt.Println(i)
}

func run(arr []int, v int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l+(r-l)/2
		if v > arr[m] {
			l = m+1
		} else if v < arr[m] {
			r = m-1
		} else {
			return m
		}
	}
	return -1
}

func run_left(arr []int, v int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l+(r-l)/2
		if arr[m] >= v {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	if l == len(arr) || arr[l] != v {
		return -1
	}
	return l
}

func run_right(arr []int, v int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l+(r-l)/2
		if arr[m] <= v {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if r < 0 {
		return -1
	}
	return r
}