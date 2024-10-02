package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	stdinText, _ := reader.ReadString('\n')
	stdinText = strings.TrimSpace(stdinText)

	var result int

	if stdinText[0] == '1' {
		result++
	}

	if stdinText[1] == '1' {
		result++
	}

	if stdinText[2] == '1' {
		result++
	}
	fmt.Println(result)
}
