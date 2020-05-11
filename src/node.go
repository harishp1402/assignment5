package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type List struct {
	tail  *Node
	start *Node
}

func (l *List) First() *Node {
	return l.start
}

func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}

func main() {
	values := &List{}
	size := 10
	for i := 0; i < size; i++ {
		node := Node{data: rand.Intn(100)}
		if node.data == 0 {
			i = i - 1
			continue
		}
		values.insertNode(&node)
		values.Display()
		values.Delete(node.data)
		fmt.Printf("%v and number is %v\n", i, node.data)
	}
	values.Display()
	// for j := 0; j < size; j++ {
	// 	values.Delete(81)
	// }
	// values.Display()
}

func (l *List) insertNode(newNode *Node) {
	if l.start == nil {
		l.start = newNode
		l.tail = newNode
	} else {
		currentNode := l.start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		l.tail = newNode
	}
}

//Display list of items
func (l *List) Display() {
	list := l.start
	for list != nil {
		fmt.Printf("value = %v and prev = %v and next= %v\n", list.data, list.prev, list.next)
		list = list.next
	}
	fmt.Println()
}

func (l *List) Delete(data int) bool {
	success := false
	node2del := l.Find(data)
	if node2del != nil {
		fmt.Println("Delete - FOUND: ", data)
		prevnode := node2del.prev
		nextnode := node2del.next
		// Remove this node
		prevnode.next = node2del.next
		nextnode.prev = node2del.prev
		success = true
	}
	return success
}

func (l *List) Find(data int) *Node {
	found := false
	var ret *Node = nil
	for n := l.First(); n != nil && !found; n = n.Next() {
		if n.data == data {
			found = true
			ret = n
		}
	}
	return ret
}
