package linkedlists

import (
	"errors"
	"testing"
)

func TestAppend(t *testing.T){
	var list *SinglyLinkedList
	err := list.Append(10)
	if err == nil{
		t.Error("error occured, cannot append to empty list")
	}
	
	if errors.Is(err,LINKED_LIST_NIL_ERROR ) == false{ 
		t.Error("error occured, wrong message for nil list")
	}
	list = NewSinglyLinkedList()
	err = list.Append(10)
	if list.start.data != 10{
		t.Error("error occured, append failed for 10")
	}
	list.Append(20)
	if list.start.next.data != 20{
		t.Error("error occured, append failed for 20")
	}
}
func Test_getLastNode(t *testing.T){
	list1 := NewSinglyLinkedList()
	list1.start = &sNode{
		data: 1,
		next: nil,
	}
	list1.start.next = &sNode{
		data: 2,
		next: nil,
	}
	list1.start.next.next = &sNode{
		data: 3,
		next: nil,
	}
	list2 := NewSinglyLinkedList()
	list2.start = &sNode{
		data: 1,
		next: nil,
	}
	list2.start.next = &sNode{
		data: 2,
		next: nil,
	}
	list3 := NewSinglyLinkedList()
	list3.start = &sNode{
		data: 1,
		next: nil,
	}
	list4 := NewSinglyLinkedList() 

	tests := []struct{
		list *SinglyLinkedList
		want any
	}{
		{
			list:list1, 
			want: 3,
		},
		{
			list: list2, 
			want: 2, 
		},
		{
			list: list3, 
			want: 1, 
		},
		{
			list: list4, 
			want: nil,
		},
		{
			list: nil, 
			want: nil,
		},
	}

	for _,test := range tests{
		t.Run("", func(t *testing.T){
			actual := test.list.getLastNode()
			if actual != nil && actual.data != test.want.(int){
				t.Errorf("error occured, wanted %v but got %v", test.want, actual)
			}else if actual == nil && test.want != nil{
				t.Errorf("error occured, expected %v but got %v", test.want, actual)
			}
		})
	}
}
