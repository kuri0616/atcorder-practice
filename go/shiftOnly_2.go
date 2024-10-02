package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	ary := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&ary[i])
	}

	res := 0

	for {
		existOdd := false

		for i := 0; i < n; i++ {
			if ary[i]%2 != 0 {
				existOdd = true
				break
			}
		}

		if existOdd {
			break
		}

		for i := 0; i < n; i++ {
			ary[i] /= 2
		}
		res++
	}

	fmt.Println(res)
}
