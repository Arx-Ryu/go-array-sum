package main

import (
	"fmt"
)

func main() {

	//Input
	var num int
	arr1 := []int{1,2,3,4,5}
	fmt.Print("Masukan Bilangan : ")
	fmt.Scanln(&num)

	//Process
	result := arraySum(arr1, num)

	//Output
	fmt.Print("Kemungkinan Penjumlahan : ", result)
}

func arraySum(arr1 []int, num int) int {
	result := 0
	for x := 0; x < len(arr1); x++ {
		for y := 0; y < len(arr1)-x; y++ {
			if arr1[x] + arr1[y+x] == num && arr1[x] != arr1[y+x] {
				result++
			}
		}
	}

	return result
}