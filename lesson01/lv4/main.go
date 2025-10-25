package main

import "fmt"

func main() {
	var userInput string
	fmt.Printf("(Please enter your expr)> ")
	fmt.Scanf("%v", &userInput)
	s := charProcesser(userInput)

	p := initProcessor()
	p.handel(s)

	c := initCaculator(p.s2.data)
	c.runCal()

	fmt.Printf("(The Result:) %v", c.cal.data[0])
}
