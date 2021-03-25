package main 

import (
	"fmt"
)

func main(){
	var str string 
	fmt.Scan(&str)

	length := subString(str)
	fmt.Println(length)
	
	return 
}

func subString(str string) int {
	res := 0
	l, r := 0, 0
	//byte 1字节
	//rune 4字节
	mp := make(map[byte]int)

	for  ; r < len(str) ; r++ {
		mp[str[r]]++

		char := str[l]
		if mp[char] > 1 {
			mp[char]--
			l++
		}

		res = max(res, r - l + 1)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


/*
func subString(str string) int {
	res := 0
	l, r := 0, 0
	//byte 1字节
	//rune 4字节
	mp := make(map[byte]int)

	for  ; r < len(str) ; r++ {
		mp[str[r]]++

		for len(mp) > 2 {
			char := str[l]
			l++

			if mp[char] > 1 {
				mp[char]--
			} else {
				delete(mp, char)
			}

		}
		res = max(res, r - l + 1)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

*/