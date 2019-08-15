package main

import (
	"fmt"
)

func sumSubarrayMins1(A []int) int {
	var sum int
	double := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		double[i] = make([]int, len(A)+1)
	}
	fmt.Println(double)
	for i := 0; i < len(A); i++ {
		double[i][i] = A[i]
		for j := i; j < len(A); j++ {
			if A[j] < double[i][j] {
				double[i][j+1] = A[j]
				sum += A[j]
			} else {
				double[i][j+1] = double[i][j]
				sum += double[i][j]
			}
			//sum += findMin(A[i:j+1])
		}
	}
	fmt.Println(double)
	return sum % (1e9 + 7)
}
func sumSubarrayMins(A []int) int { //动态规划
	var sum int
	for i := 0; i < len(A); i++ {
		min := A[i]
		for j := i; j < len(A); j++ {
			if A[j] < min {
				min = A[j] //子数组的最小值
			}
			sum += min
		}
	}
	return sum % (1e9 + 7)
}

func main() {
	A := []int{3, 1, 2, 4}
	fmt.Println(sumSubarrayMins(A))
	/*fmt.Println(2508796223%int(1e9+7))
	var arr = []int{33,2}
	fmt.Println(arr)*/

	d := make([]int, 1)
	fmt.Println(d)
}
