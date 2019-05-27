package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	data int
	next *Node
}

func (node Node) print() {
	fmt.Print("[" + strconv.Itoa(node.data) + "] ")
}

func printList(head *Node) {
	tmp := head
	for tmp != nil {
		tmp.print()
		tmp = tmp.next
	}
	fmt.Println()
}

func addEnd(head *Node, data int) {
	tmp := head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &Node{data: data, next: nil}
}

func main() {
	head := Node{data: 1, next: nil}
	node1 := Node{data: 2, next: nil}
	head.next = &node1
	node2 := Node{data: 3, next: nil}
	node1.next = &node2
	addEnd(&head, 100)
	printList(&head)
}
