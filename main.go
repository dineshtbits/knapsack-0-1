package main

import (
	"fmt"
)

func main() {
	//inputs
	w := 50
	items := []int{10, 20, 30}
	values := []int{60, 100, 120}

	table := make([][]int, len(items)+1)
	for t := range table {
		table[t] = make([]int, w+1)
	}

	// top down table
	for i := 1; i <= len(items); i++ {
		for j := 1; j <= w; j++ {
			if j >= items[i-1] {
				currentVal := values[i-1]
				remCapacity := j - items[i-1]
				remMaxVal := table[i-1][remCapacity]

				table[i][j] = Max(table[i-1][j-1], currentVal+remMaxVal)
			}
		}
	}
	fmt.Println(table[len(items)][w])

}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
