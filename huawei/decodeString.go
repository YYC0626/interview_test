package main

import (
	"fmt"
	"strconv"
)

/*

func main() {
	var str string
	fmt.Scan(&str)

	res := decoder(str)

	fmt.Println(res)
	return
}

//括号匹配考虑使用栈来实现
//遇到左括号压栈，遇到右括号弹出

//维护字符串的切片temp，存储当前需要重复的序列。（当作栈来用）
//维护一个flag值，为true的时候将遍历到的字符算作重复次数的部分，为false的时候将遍历到的字符算作需要重复的字串部分。（置于最末，即最后才进行判断）

func decoder(code string) string {
	stk := []rune{}
	temp := []string{}
	repeat := ""
	repeatTime := ""
	flag := false

	result := ""

	for _, char := range code {

		//	fmt.Println(stk)
		//	fmt.Println(temp)
		//	fmt.Println(repeatTime)

		//
		//遇到左括号，开始计temp
		if char == '(' {
			//左小括号压入，开始统计重复的序列
			stk = append(stk, char)
			//压入新的temp值
			temp = append(temp, "")
			//
			flag = false
			continue
		}

		//	temp已记录完毕
		if char == ')' {
			//重复段结束,‘）’对应的‘（’弹出，此时的temp顶部存储的即为需要重复的串
			stk = stk[:len(stk)-1]
			//复制栈顶字符串，弹栈
			repeat = temp[len(temp)-1]
			temp = temp[:len(temp)-1]

			continue
		}

		//
		//遇到‘<'，压栈，开始记录重复次数字符串
		if char == '<' {
			//左尖括号压栈，开始统计重复次数
			stk = append(stk, char)
			//
			flag = true
			repeatTime = ""
			continue
		}

		//遇到'>'，弹出对应的'<'，然后把记录的重复次数变为int
		if char == '>' {
			//重复次数统计结束
			stk = stk[:len(stk)-1]
			//记录的次数转换为数字
			num, _ := strconv.Atoi(repeatTime)
			//重复添加num次（添加到外一层的temp里）
			for i := 0; i < num; i++ {
				if len(temp) == 0 {
					result += repeat
					continue
				}
				temp[len(temp)-1] += repeat
			}

			continue
		}

		//栈为空时（尚未碰到括号）
		//直接附加在result后
		if len(stk) == 0 {
			result += string(char)
			continue
		}

		if flag == false {
			temp[len(temp)-1] += string(char)
		} else {
			repeatTime += string(char)
		}

	}

	return result
}

*/

/*
递归做法：
	遇到小括号则调用递归函数，其内部完成：获取repeat，获取repeatTime，返回拼接好的字符串和完成后的末尾下标
*/

func main() {
	var str string
	fmt.Scan(&str)

	res, _ := recur(str, 0)

	fmt.Println(res)
	return
}

func recur(code string, index int) (string, int) {
	result := ""
	temp := ""

	//从i开始遍历
	i := index
	for ; i < len(code); i++ {
		//如果当前字符不是‘（’，说明是第一次遍历（非内层递归调用），直接附加在result后
		if code[i] != '(' {
			result += string(code[i])
			continue
		}

		//是‘（’
		if code[i] == '(' {
			//内层处理完，返回结果和下标，此处是‘》’的下标，因为循环体结束还有i++操作
			temp, i = recur(code, i+1)

			result += temp
			return result, i
		}

		//
		if code[i] == ')' {
			//循环体结束了，在这里更新repeat的值
			//获取重复次数
			//更新返回值字符串
			n := 0
			//i是’）‘，i+1就是‘《’
			n, i = getTime(code, i+1)
			for j := 0; j < n; j++ {
				result += temp
			}

			return result, i
		}
		temp += string(code[i])

	}
	return result, i
}

func getTime(code string, index int) (int, int) {
	time := ""
	for i := index; i < len(code); i++ {

		if code[i] == '<' {
			continue
		}
		if code[i] == '>' {
			num, _ := strconv.Atoi(time)
			return num, i
		}
		time += string(code[i])
	}
	//调用时一定保证从if里返回
	return -1, index
}

/*
func main() {
	var str string
	fmt.Scan(&str)

	res := recur(str, 0)

	fmt.Println(res)
	return
}

// *	递归函数实现的思路：
// *		遇到‘（’即调用递归函数，开启记录temp的flag，遇到‘）’则停止记录；
// *		遇到‘《’则开始记录重复次数，遇到‘》’则停止记录，将所记录的次数由字符串转换为数字，然后在result后将temp连续附加time次；
// *		然后返回result
// *




//考虑递归实现，遇到左括号则递归进入下一层，遇到右括号且栈为空
func recur(code string, start int) string {
	//flag := false
	result := ""

	time := ""
	temp := ""

	dotemp := false
	dotime := false

	for i := start; i < len(code); i++ {
		//如果遇到了左括号并且标识位是真（不是第一次出现），则递归
		//if code[i] == '(' && flag == false {
		//	flag = true
		//	continue
		//}

		//如果不是第一次出现，就递归进入下一层
		if code[i] == '(' {
			//recur
			temp += recur(code, i)
		}
		//右括号记录temp
		if code[i] == ')' {
			dotemp = false
			continue
		}
		if i < len(code) && dotemp == true {
			temp += string(code[i])
			continue
		}

		//统计重复的次数
		//左括号开始，标志符置为true
		//右括号结束，开始往result后添加
		if code[i] == '<' {
			time = ""
			dotime = true
			continue
		}
		if code[i] == '>' {
			n, _ := strconv.Atoi(time)

			for i := 0; i < n; i++ {
				result += temp
			}
			return result
		}
		if i < len(code) && dotime == true {
			time += string(code[i])
			continue
		}

		result += string(code[i])
		//	temp += code[i]
	}

	return result
}
*/
