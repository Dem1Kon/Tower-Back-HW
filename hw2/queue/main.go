package main

import "fmt"

type Node struct {
	value interface{}
	front *Node
	back  *Node
}

type Queue struct {
	size         int
	firstToLeave *Node
	Last         *Node
}

func (q *Queue) Add(value interface{}) {
	newNode := &Node{value: value, front: q.Last}

	if q.size != 0 {
		q.Last.back = newNode
	} else {
		q.firstToLeave = newNode
	}

	q.Last = newNode
	q.size++
}

func (q *Queue) Pop() {

	if q.size != 0 {

		if q.size == 1 {

			q.firstToLeave = nil
			q.Last = nil

		} else {

			q.firstToLeave = q.firstToLeave.back
			q.firstToLeave.front = nil

		}

		q.size--

	} else {

		fmt.Println("Queue is empty")

	}
}

func (q *Queue) isExist(element interface{}) bool {
	currentNode := q.Last

	isExist := false

	for i := 0; i < q.size; i++ {
		if currentNode.value == element {
			isExist = true
			break
		}

		currentNode = currentNode.front
	}
	return isExist
}

func main() {

}
