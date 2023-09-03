package stacks

import "errors"

type Node struct {
	data interface{}
	next *Node
}
type Stack struct {
	top  *Node
	size int
}

func NewStack() *Stack {
	return &Stack{}
}
func (s *Stack) Size() int {
	return s.size
}
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
func (s *Stack) Peek() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("empty stack")
	}
	return s.top.data, nil
}
func (s *Stack) Push(data interface{}) {
	s.top = &Node{data, s.top}
	s.size++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("stack empty")
	}
	data := s.top.data
	s.top = s.top.next
	s.size--
	return data, nil
}
