package linkedlists

import (
	"errors"
)
var LINKED_LIST_NIL_ERROR = errors.New("linked list is nil")
type sNode struct{
	data int
	next *sNode
}
type SinglyLinkedList struct{
	start *sNode
}

func NewSinglyLinkedList() *SinglyLinkedList{
	return &SinglyLinkedList{}
}

func (list *SinglyLinkedList) Append(data int) error{
	if list == nil {
		return LINKED_LIST_NIL_ERROR
	}
	if list.start == nil{
		list.start = &sNode{
			data: data,
			next: nil,
		}
	}else{
		last := list.getLastNode()
		last.next = &sNode{
			data: data,
			next: nil,
		}	
	}
	return nil
}

/*
This function checks a given list and returns the last node. If the list is empty or nil, it returns nil.
*/
func (list *SinglyLinkedList) getLastNode() *sNode{
	if list == nil || list.start == nil{
		return nil
	}
	start := list.start
	for start.next != nil{
		start = start.next	
	}
	return start
}
