package main

import "fmt"

func main() {
	var price, coin int
	const fiveHundred int = 500

	fmt.Scan(&price, &coin)

	if price%fiveHundred == 0 {
		fmt.Println("Yes")
		return
	}

	for {
		if price < fiveHundred {
			break
		}
		price -= fiveHundred
	}

	for i := 0; i < coin+1; i++ {
		if price == 0 {
			fmt.Println("Yes")
			return
		}
		price -= 1
	}
	fmt.Println("No")
}
