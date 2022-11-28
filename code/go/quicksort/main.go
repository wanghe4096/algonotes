package main

import "fmt"

func sort(arr []int) {
	p, r := 0, len(arr)-1
	qsort(arr, p, r)
}
func qsort(arr []int, p, r int) {
	if r <= p {
		return
	}
	q := partition(arr, p, r)
	qsort(arr, p, q-1)
	qsort(arr, q+1, r)
	return
}

func partition(arr []int, p, r int) int {
	pivot := arr[p]
	i, j := p+1, r
	for {
		for arr[i] < pivot {
			i++
			if i == r {
				break
			}
		}

		for pivot < arr[j] {
			j--
			if j == p {
				break
			}
		}
		if i >= j {
			break
		}
		exch(arr, i, j)
	}
	//	fmt.Printf("[p=%d] j=%d arr= %+v\n", pivot, j, arr)
	exch(arr, p, j)
	return j
}

func exch(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 第一种方法
func part(a []int, p, r int) int {
	pivot := a[p]
	j := p
	for i := p; i <= r; i++ {
		if a[i] < pivot {
			j++
			fmt.Printf("[swap] i=%d\tj=%d\n", i, j)
			a[i], a[j] = a[j], a[i]
		}
	}

	fmt.Printf("j=%d\ta==%+v\n", j, a)

	a[p], a[j] = a[j], a[p]
	return j
}

func main() {
	arr := []int{5, 3, 1, 8, 9, 4, 6, 2, 7}

	low, high := 0, len(arr)-1
	fmt.Printf("low=%d\thigh=%d\n", low, high)
	qsort(arr, low, high)
	fmt.Printf("result=%+v\n", arr)

}
