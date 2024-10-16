package main

import "fmt"

type Node struct {
	value interface{}
	front *Node
	back  *Node
}

type Deck struct {
	size  int
	last  *Node
	first *Node
}

//Addfront, AddBack, PopBack, PopFront, isExist

func (d *Deck) AddFront(value interface{}) {

	newNode := &Node{value: value, back: d.first}
	if d.size != 0 {
		d.first.front = newNode
		d.first = newNode
	} else {
		d.first = newNode
		d.last = newNode
	}

	d.size++

}

func (d *Deck) AddBack(value interface{}) {

	newNode := &Node{value: value, front: d.last}

	if d.size != 0 {
		d.last.back = newNode
		d.last = newNode
	} else {
		d.first = newNode
		d.last = newNode
	}

	d.size++

}

func (d *Deck) PopFront() {
	if d.size != 0 {

		if d.size != 1 {
			d.first = d.first.back
			d.first.front = nil
		} else {
			d.first = nil
			d.last = nil
		}
		d.size--
	} else {
		fmt.Println("Deck is empty!")
	}
}

func (d *Deck) PopBack() {
	if d.size != 0 {
		if d.size != 1 {

			d.last = d.last.front
			d.last.back = nil

		} else {
			d.last = nil
			d.first = nil
		}
		d.size--
	} else {
		fmt.Println("Deck is empty!")
	}
}

func (d *Deck) isExist(element interface{}) bool {
	currentNode := d.last

	isExist := false

	for i := 0; i < d.size; i++ {
		if currentNode.value == element {
			isExist = true
			break
		}
		currentNode = currentNode.front
	}

	return isExist
}

func main() {
	deck := new(Deck)

	deck.AddFront(123)
	deck.AddBack("hello")
	deck.AddFront("Sosi")
	deck.AddBack('r')

	fmt.Println(deck.isExist(""))

}
