package main

import (
	"fmt"
	"strconv"
)

// Interface is just a collection of method signatures
// To implement interface, just implement the methods defined in the interface
type LinkedList interface {
	add(val int)
	removeLast() int
	size() int
}

type SLLNode struct {
	next *SLLNode
	val int
}

// SinglyLinkedList is implementing the LinkedList interface as it is implementing all the methods defined in the interface
type SinglyLinkedList struct {
	count int
	head *SLLNode
	tail *SLLNode
}

// Implementing add(val int) method in LinkedList interface
func (sll *SinglyLinkedList) add(val int) {
	sllNode := new(SLLNode)
	sllNode.val = val
	sllNode.next = nil
	if sll.head == nil {
		sll.head = sllNode
		sll.tail = sllNode
	} else {
		tmpHead := sll.head
		for tmpHead.next != nil {
			tmpHead = tmpHead.next
		}
		tmpHead.next = sllNode
		sll.tail = sllNode
	}
	sll.count++
}

// Implementing removeLast() method in LinkedList interface
func (sll *SinglyLinkedList) removeLast() int {
	prevNode := sll.head
	tmpHead := sll.head
	for tmpHead.next != nil {
		prevNode = tmpHead;
		tmpHead = tmpHead.next
	}

	if (sll.size() > 1) {
		prevNode.next = nil
		sll.tail = prevNode
	} else {
		sll.head = nil
		sll.tail = nil

	}

	sll.count--
	return tmpHead.val
}

// Implementing size() method in LinkedList interface
func (sll *SinglyLinkedList) size() int {
	return sll.count
}

// Implementing String() method in Stringer interface. Similar to toString() in java
func (sll *SinglyLinkedList) String() string {
	resStr := ""
	for tmpHead := sll.head; tmpHead != nil; tmpHead = tmpHead.next {
		resStr += "->" + strconv.Itoa(tmpHead.val)
	}

	return resStr
}

func main() {
	sll := new(SinglyLinkedList)
	sll.add(10)
	sll.add(20)
	sll.add(30)
	fmt.Println("Size of the LinkedList : ", sll.size())
	fmt.Println("Contents of the LinkedList : ", sll)
	fmt.Println("Removing : ", sll.removeLast())
	fmt.Println("Removing : ", sll.removeLast())
	//fmt.Println("Removing : ", sll.removeLast())
	fmt.Println("Contents of the LinkedList after removal : ", sll)
	fmt.Println("Size of the LinkedList after removal : ", sll.size())
}