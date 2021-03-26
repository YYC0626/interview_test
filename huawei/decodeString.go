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
			//重复次数记录完成，flag重新置为false
			flag = false
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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
递归做法：
	遇到小括号则调用递归函数，其内部完成：获取repeat，获取repeatTime，返回拼接好的字符串和完成后的末尾下标

*/

func main() {
	var str string
	fmt.Scan(&str)

	res := decodeStr(str)

	fmt.Println(res)
	return
}

//题解函数，开启主循环
func decodeStr(code string) string {
	result := ""

	var i int
	for i = 0; i < len(code); i++ {
		if code[i] == '(' {
			var temp string
			temp, i = recur(code, i+1)
			result += temp
			//		continue
		}
		if code[i] != '(' && code[i] != ')' && code[i] != '<' && code[i] != '>' {
			result += string(code[i])
		}
		continue
	}
	return result
}

//递归函数：接收原字符串和传递而来的下标
//该函数是碰到'('时调用的，所以碰到')'就要结束自身的循环，然后完成处理
func recur(code string, index int) (string, int) {
	result := ""
	temp := ""

	//从i开始遍历
	i := index

	for ; code[i] != ')'; i++ {
		//再次遇到则继续递归调用
		if code[i] == '(' {
			s, ind := recur(code, i+1)

			//重复串处理后加在temp之后，i值更新
			temp += s
			i = ind
		}

		temp += string(code[i])
	}

	//循环体结束了，在这里更新repeat的值
	//获取重复次数
	//更新返回值字符串
	//i是’）‘，i+1就是‘《’
	n, newindex := getTime(code, i+1)
	for j := 0; j < n; j++ {
		result += temp
	}

	return result, newindex

}

//该函数获取尖括号内的次数
func getTime(code string, index int) (int, int) {
	time := ""
	for i := index; i < len(code); i++ {

		if code[i] == '<' {
			continue
		}
		if code[i] == '>' {
			num, _ := strconv.Atoi(time)
			return num, i + 1
		}
		time += string(code[i])
	}
	//调用时一定保证从if里返回
	return -1, index
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
