package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	var count int
	fmt.Scanf("%d", &n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		input := sc.Text()
		nums[i], _ = strconv.Atoi(input)
	}
	result := shift(nums, count)
	fmt.Println(result)

}

func shift(nums []int, count int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i]%2 != 0 {
			return count
		}
	}
	count += 1
	for i := 0; i < n; i++ {
		nums[i] = nums[i] / 2
	}
	count = shift(nums, count)
	return count
}
