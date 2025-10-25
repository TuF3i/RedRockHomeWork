package main

import "fmt"

type stack struct {
	data []string
}

func initStack() stack {
	return stack{data: make([]string, 0)}
}

func (s *stack) push(value string) {
	s.data = append(s.data, value)
}

func (s *stack) pop() (string, error) {
	if s.ifEmpty() {
		return "0", fmt.Errorf("Stack is Empty!")
	}

	peek := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return peek, nil
}

func (s *stack) getPeek() (string, error) {
	if s.ifEmpty() {
		return "0", fmt.Errorf("Stack is Empty!")
	}
	return s.data[len(s.data)-1], nil
}

func (s *stack) ifEmpty() bool {
	return len(s.data) == 0
}
