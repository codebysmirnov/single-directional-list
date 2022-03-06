package main

import "fmt"

// Node element has information, data and information of next element
type Node struct {
	Data interface{}
	Next *Node
}

// List represents the starting point of the relationship of elements
type List struct {
	begin  *Node
	length int
}

// Remove remove element from List by passed index
func (l *List) Remove(index int) {
	if l.begin == nil {
		panic("Out of range!")
	}
	if index < 0 || index >= l.length {
		panic("Out of range!")
	}
	if index == 0 {
		l.begin = l.begin.Next
		return
	}

	current := l.begin
	count := 0
	for current.Next != nil {
		if index == count+1 {
			next := current.Next
			current.Next = next.Next
			return
		}
		count++
		current = current.Next
	}
	panic("Out of range!")
}

// Len return's length of List
func (l List) Len() int {
	return l.length
}

// Get return element of List by passed index
func (l List) Get(index int) interface{} {
	if l.begin == nil {
		panic("Out of range!")
	}
	if index < 0 || index >= l.length {
		panic("Out of range!")
	}
	if index == 0 {
		return l.begin.Data
	}
	count := 0
	current := l.begin
	for current.Next != nil {
		if count == index {
			return current.Data
		}
		current = current.Next
		count++
	}
	panic("Out of range!")
}

// PushFront add passed element to the beginning List
func (l *List) PushFront(element interface{}) {
	l.length++
	if l.begin == nil {
		l.begin = &Node{Data: element}
		return
	}
	first := l.begin
	l.begin = &Node{Data: element, Next: first}

	return
}

// Append passed element to the end of List
func (l *List) Append(element interface{}) {
	l.length++
	if l.begin == nil {
		l.begin = &Node{Data: element}
		return
	}
	last := l.begin
	for last.Next != nil {
		last = last.Next
	}
	last.Next = &Node{Data: element}
	return
}

// String returns string of all List elements
func (l List) String() string {
	if l.begin == nil {
		return "[]"
	}

	result := fmt.Sprintf("[%v", l.begin.Data)
	current := l.begin
	for current.Next != nil {
		current = current.Next
		result += fmt.Sprintf(", %v", current.Data)
	}
	result += "]"
	return result
}

func main() {
	var l List

	l.Append("cat")
	l.Append("dog")
	l.Append("tiger")

	l.PushFront("elephant")

	fmt.Printf("List: %s\n", l)
	fmt.Printf("Get: %s\n", l.Get(1))
	fmt.Printf("Length: %d\n", l.Len())

	l.Remove(1)
	fmt.Printf("Removed: %s\n", l)

}
