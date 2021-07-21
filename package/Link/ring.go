package Link

import (
	"fmt"
)

type Linklist struct {
	Next *Linklist
	Data interface{}``
}

//func main() {
//	var Head *Linklist = nil
//	fmt.Println(Head)
//	Head = NewEmpty(3, Head)
//	for current := Head; current != nil; current = current.Next {
//		fmt.Println(current)
//	}
//	AssignmentEmpty(Head, 1, 2, 3)
//	for current := Head; current != nil; current = current.Next {
//		fmt.Println(current)
//	}
//	DeleteData(AddressFind(3, Head), Head)
//	PrintAllList(Head)
//	Change(AddressFind(2, Head), 3)
//	PrintAllList(Head)
//	Add(AddressFind(1, Head))
//	Add(AddressFind(3, Head))
//	AssignmentEmpty(Head,4,5)
//	PrintAllList(Head)
//}

// NewEmpty  创建一个长度为n的空表
func NewEmpty(n int, Head *Linklist) *Linklist {
	if n <= 0 {
		return nil
	}
	var tail *Linklist
	for i := 0; i < n; i++ {
		current := new(Linklist)
		if Head == nil {
			Head = current
		} else {
			tail.Next = current
		}
		tail = current
	}
	return Head
}

// AssignmentEmpty 向空表中赋值
func AssignmentEmpty(Head *Linklist, data ...interface{}) bool {
	if Head == nil {
		return false
	}

	for cnt, current := 0, Head; current != nil; current = current.Next {
		if current.Data == nil {
			current.Data = data[cnt]
			cnt++
		}
	}

	return true
}

// AddressFind 按照所给的值查找所对应的链表块地址
func AddressFind(val interface{}, Head *Linklist) *Linklist {
	for current := Head; current != nil; current = current.Next {
		if current.Data == val {
			return current
		}
	}
	return nil
}

// Add 增加一个结点
func Add(Position *Linklist) bool {
	if Position == nil {
		return false
	}
	current := new(Linklist)
	current.Next = Position.Next
	Position.Next = current
	return true
}

func Change(Position *Linklist, val interface{}) bool {
	if Position == nil {
		return false
	}
	Position.Data = val
	return true
}

// PrintAllList 输出所有内容
func PrintAllList(Head *Linklist) {
	for current := Head; current != nil; current = current.Next {
		fmt.Println(current.Data)
	}
}

// DeleteData 删除指定位置的信息
func DeleteData(Position, Head *Linklist) {
	if Position == nil {
		return
	}
	if Head == Position {
		Head = Head.Next
		return
	}
	for current := Head; current != nil; current = current.Next {
		if current.Next == Position {
			if current.Next.Next == nil {
				current.Next = nil
				return
			}
			current.Next = current.Next.Next
			return
		}
	}
}
