package main

import (
	"fmt"
)

func romanToInt(roman string) int {
	sum := 0

	hashTable := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(roman); i++ {
		if i == len(roman)-1 {
			sum += hashTable[string(roman[i])]
		} else {
			if hashTable[string(roman[i])] < hashTable[string(roman[i+1])] {
				sum -= hashTable[string(roman[i])]
			} else {
				sum += hashTable[string(roman[i])]
			}
		}
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("III"))     // 输出 3
	fmt.Println(romanToInt("MCMXCIV")) // 输出 1994
	fmt.Println(romanToInt("LVII"))    // 输出 57
	fmt.Println(romanToInt("IX"))      // 输出 9
}
