package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Contain(userInput string) string {
	for i := 0; i < len(userInput); i++ {
		c := string(userInput[i])
		//fmt.Printf("%v\n", c)
		switch c {
		case "+":

			return "+"
		case "-":

			return "-"
		case "*":

			return "*"
		case "/":

			return "/"
		default:

			continue
		}
	}
	return "OP Err"
}

func main() {
	var userInput string
	var num1 float64
	var num2 float64
	var op string

	fmt.Printf("Please Input Your expr>")
	fmt.Scanln(&userInput)
	//fmt.Printf(userInput)
	op = Contain(userInput)

	if op == "OP Err" {
		fmt.Printf("OP Error\n")
		return
	}

	//fmt.Printf("%v\n", op)
	res := strings.Split(userInput, op)
	//fmt.Println(res)

	num1, err := strconv.ParseFloat(res[0], 64)
	if err != nil {
		fmt.Printf("Convert Num1 Fail!")
	}

	num2, err = strconv.ParseFloat(res[1], 64)
	if err != nil {
		fmt.Printf("Convert Num2 Fail!")
	}

	switch op {
	case "+":
		fmt.Printf("%f + %f = %.2f", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%f - %f = %.2f", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%f * %f = %.2f", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Printf("Runtime Error: (%v / 0)\n", num1)
			return
		}
		fmt.Printf("%f / %f = %.2f", num1, num2, num1/num2)
	default:
		fmt.Printf("Operation Error!")
	}
}
