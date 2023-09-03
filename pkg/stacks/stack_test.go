package stacks

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stk := NewStack()

	if stk == nil {
		t.Fatalf("Stack is nil")
	}
	if stk.size != 0 {
		t.Errorf("empty stack size not zero")
	}
	if stk.top != nil {
		t.Errorf("top of new stack is not nil")
	}
	if stk.size != 0 {
		t.Errorf("empty stack cannot have size > 0")
	}

}
func TestIsEmpty(t *testing.T) {
	stk := NewStack()
	if stk.IsEmpty() == false {
		t.Errorf("new stack is always empty")
	}
	stk.Push(22)
	stk.Push(33)
	if stk.IsEmpty() == true {
		t.Errorf("error: stack is actually not empty")
	}
	stk.Pop()
	stk.Pop()
	if stk.IsEmpty() == false {
		t.Errorf("error: stack is actually empty now")
	}
}
func TestPush(t *testing.T) {
	stk := NewStack()

	stk.Push(22)
	if stk.size != 1 {
		t.Errorf("pushed first element, size should be 1")
	}
	if stk.top.data != 22 {
		t.Errorf("wrong element pushed")
	}
	stk.Push(33)
	if stk.size != 2 && stk.top.data != 33 {
		t.Errorf("push() not working, expected size 2 and expected data 33")
	}
}
func TestPop(t *testing.T) {
	stk := NewStack()

	var val any
	var err error

	stk.Push(22)
	stk.Push(33)
	stk.Push(44)
	val, _ = stk.Pop()

	if val != 44 {
		t.Errorf("wrong element popped from stack")
	}
	val, _ = stk.Pop()
	if val != 33 {
		t.Errorf("wrong element popped from stock")
	}
	stk.Pop()
	_, err = stk.Pop()
	if err == nil {
		t.Errorf("popping out from empty stack")
	}
}
func TestPeek(t *testing.T) {
	stk := NewStack()

	stk.Push(22)
	stk.Push(44)
	stk.Push(33)
	stk.Pop()
	stk.Pop()
	val, _ := stk.Peek()
	if val != 22 {
		t.Errorf("peek returned wrong value from stack")
	}
}
func TestSize(t *testing.T) {
	stk := NewStack()

	if stk.Size() != 0 {
		t.Errorf("empty stack size != 0")
	}
	stk.Push(33)
	stk.Push(44)
	stk.Push(55)
	stk.Pop()
	if stk.Size() != 2 {
		t.Errorf("wrong stack size returned")
	}
}
