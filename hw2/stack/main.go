package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	size    int
	topNode *Node
}

func (f *Stack) Add(element interface{}) {
	newNode := &Node{value: element, next: f.topNode}

	f.topNode = newNode

	f.size++
}

func (f *Stack) Pop() {
	if f.size != 0 {
		f.topNode = f.topNode.next

		f.size--
	} else {
		fmt.Print("Stack is empty!")
	}
}

func (f *Stack) isExist(element interface{}) bool {
	currentNode := f.topNode

	isExist := false

	for i := 0; i < f.size; i++ {
		if currentNode.value == element {
			isExist = true
			break
		}
		currentNode = currentNode.next
	}

	return isExist
}

func main() {
	
}
