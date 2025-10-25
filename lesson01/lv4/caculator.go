package main

import (
	"strconv"
)

func initCaculator(p []string) caculator {
	return caculator{
		source: p,
		cal:    initStack(),
	}
}

type caculator struct {
	source []string
	cal    stack
}

func (c *caculator) FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func (c *caculator) runCal() {
	for _, item := range c.source {
		if _, err := strconv.ParseFloat(item, 64); err == nil {
			c.cal.push(item)
		} else {
			r, _ := c.cal.pop()
			l, _ := c.cal.pop()
			op := item

			right, _ := strconv.ParseFloat(r, 64)
			left, _ := strconv.ParseFloat(l, 64)

			switch op {
			case "+":
				res := left + right
				c.cal.push(c.FloatToString(res))
			case "-":
				res := left - right
				c.cal.push(c.FloatToString(res))
			case "*":
				res := left * right
				c.cal.push(c.FloatToString(res))
			case "/":
				res := left / right
				c.cal.push(c.FloatToString(res))
			}
		}
	}
}
