package main

import "fmt"

func main() {
	var price, coin int
	const fiveHundred int = 500

	fmt.Scan(&price, &coin)

	remainder := price % fiveHundred

	if remainder <= coin {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
