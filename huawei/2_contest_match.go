package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)

	res := contestMatch(num)

	fmt.Println(res)
}

func contestMatch(n int) string {
	strs := []string{}

	for i := 1; i <= n; i++ {
		strs = append(strs, strconv.Itoa(i))
	}

	for n > 1 {
		for i := 0; i < n/2; i++ {
			strs[i] = "(" + strs[i] + "," + strs[len(strs)-1] + ")"
			strs = strs[:len(strs)-1]
		}
		n /= 2
	}

	return strs[0]
}
