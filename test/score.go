package main

import "fmt"

//实现分数排名
//输入：一个数组，0.00～5.00
//输出对数组的排名：可并列

//思路：
//考虑计数排序

func main() {
	var length int
	var input []float32

	fmt.Scan(&length)
	for i := 0; i < length; i++ {
		var temp float32
		fmt.Scan(&temp)
		input = append(input, temp)

	}

	res := score(input)
	fmt.Println(res)
}

func score(input []float32) []int {
	result := []int{}
	count := [501]int{}
	//计数
	for i := 0; i < len(input); i++ {
		t := int(input[i] * 100)
		count[t]++
	}
	//哈希表从后往前
	mp := make(map[float32]int)
	rank := 1
	for i := 500; i >= 0; i-- {
		//
		if count[i] != 0 {
			mp[float32(i)/100.00] = rank
			rank++
		}
		continue
	}

	for i := 0; i < len(input); i++ {

		result = append(result, mp[input[i]])
	}

	return result
}
