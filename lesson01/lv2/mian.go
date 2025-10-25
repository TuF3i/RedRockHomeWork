package main

import (
	"fmt"
)

var num1 float64
var operation string
var num2 float64
var mode int

func calculatorSwitch() (float64, error) {
	switch operation {
	case "+":

		return num2 + num1, nil

	case "-":

		return num1 - num2, nil

	case "*":

		return num1 * num2, nil

	case "/":

		return num1 / num2, nil

	default:

		return 0, fmt.Errorf("你输了什么，害怕")
	}
}

func calculatorIf() (float64, error) {
	if operation == "+" {
		return num1 + num2, nil
	}

	if operation == "-" {
		return num1 - num2, nil
	}

	if operation == "*" {
		return num1 * num2, nil
	}

	if operation == "/" {
		return num1 / num2, nil
	}

	return 0, fmt.Errorf("你输了什么，害怕")
}

func main() {
	fmt.Printf("选择一个运行模式(0:if 1:switch): ")
	_, err := fmt.Scanf("%v", &mode)
	if err != nil {
		fmt.Printf("你输了什么，害怕")
		return
	}

	fmt.Printf("请输入第一个数: ")
	_, err = fmt.Scanf("%v", &num1)
	if err != nil {
		fmt.Printf("你输了什么，害怕")
		return
	}

	fmt.Printf("请输入操作符: ")
	_, err = fmt.Scanf("%v", &operation)
	if err != nil {
		fmt.Printf("你输了什么，害怕")
		return
	}

	fmt.Printf("请输入第二个数: ")
	_, err = fmt.Scanf("%v", &num2)
	if err != nil {
		fmt.Printf("你输了什么，害怕")
		return
	}

	if mode == 0 {
		res, err := calculatorIf()
		if err != nil {
			fmt.Printf("你输了什么，害怕")
			return
		}
		fmt.Printf("结果是: %f", res)
		return
	}

	if mode == 1 {
		res, err := calculatorSwitch()
		if err != nil {
			fmt.Printf("你输了什么，害怕")
			return
		}
		fmt.Printf("结果是: %f", res)
		return
	}

	fmt.Printf("你输了什么，害怕")
}
