package main

import (
	"fmt"
	"strconv"
)

func main() {
	var rCard, gCard, bCard string

	fmt.Scan(&rCard, &gCard, &bCard)

	linedCard := rCard + gCard + bCard
	num, _ := strconv.Atoi(linedCard)
	if num%4 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
