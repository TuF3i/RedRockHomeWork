package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {

		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v ", j, i, j*i)
		}

		fmt.Println()
	}
}
