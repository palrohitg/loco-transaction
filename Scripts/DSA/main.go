package main

import "fmt"

func ways(n int, m int, i int, j int) int {
	// base
	if i < 0 || i > n || j < 0 || j > m {
		return 0
	}

	if i == m-1 && j == n-1 {
		return 1
	}

	return ways(m, n, i+1, j) + ways(m, n, i, j+1)
}

func main() {
	m := 2
	n := 2
	ans := ways(n, m, 0, 0)
	fmt.Println(ans)
	return
}
