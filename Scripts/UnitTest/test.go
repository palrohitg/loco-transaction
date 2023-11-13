package main

import (
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestAdd(t *testing.T) {
	//var tests = []struct {
	customSet := CreateCustomerString()

	//	a, b int
	//	want int
	//}{
	//	{0, 1, 0},
	//	{1, 0, 0},
	//	{2, -2, -2},
	//	{0, -1, -1},
	//	{-1, 0, -1},
	//}
	//
	//for _, tt := range tests {
	//
	//	testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
	//	t.Run(testname, func(t *testing.T) {
	//		ans := IntMin(tt.a, tt.b)
	//		if ans != tt.want {
	//			t.Errorf("got %d, want %d", ans, tt.want)
	//		}
	//	})
	//}
}

func TestContains(t *testing.T) {
}

func main() {
	TestAdd()
}
