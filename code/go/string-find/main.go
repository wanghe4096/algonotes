package main

import "fmt"

/*
* 字符串匹配算法， BF算法， 时间复杂度O(n*m)
 */
func findAtString(A, B string) int {
	n, m := len(A), len(B)
	i, j := 0, 0
	for i = 0; i < n-m+1; i++ {
		for j = i; j < m; j++ {
			if A[i+j] != B[j] {
				break
			}
		}
		if j == m {
			return i
		}
	}
	return -1
}

func main() {
	s := "CABCDEFGH"
	m := "AB"
	res := findAtString(s, m)
	fmt.Printf("主串: %s\n", s)
	fmt.Printf("模式串：%s\n", m)
	fmt.Printf("结果: %d\n", res)
}
