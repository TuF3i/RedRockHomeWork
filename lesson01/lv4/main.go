package main

import "fmt"

func main() {
	var userInput string
	fmt.Printf("(Please enter your expr)> ")
	fmt.Scanln(&userInput)

	if !safetyCheck(userInput) {
		fmt.Println("Division by zero")
		return
	}

	s := charProcesser(userInput)

	p := initProcessor()
	p.handle(s)

	c := initCaculator(p.s2.data)
	c.runCal()

	fmt.Printf("(The Result:) %v", c.cal.data[0])
}
