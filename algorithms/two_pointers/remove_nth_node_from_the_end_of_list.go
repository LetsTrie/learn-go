package main

import (
	"fmt"
)

type ListNode struct {
	value int
	next  *ListNode
}

// removeNthFromEnd removes the nth node from the end of the linked list and returns the modified list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	currentPtr := &ListNode{next: head}
	rightPtr := currentPtr
	leftPtr := currentPtr

	for i := 0; i < n; i++ {
		if rightPtr.next == nil {
			return head
		}

		rightPtr = rightPtr.next
	}

	for rightPtr.next != nil {
		rightPtr = rightPtr.next
		leftPtr = leftPtr.next
	}

	leftPtr.next = leftPtr.next.next
	return currentPtr.next
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	// Create a sample linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{value: 1}
	head.next = &ListNode{value: 2}
	head.next.next = &ListNode{value: 3}
	head.next.next.next = &ListNode{value: 4}
	head.next.next.next.next = &ListNode{value: 5}

	n := 5 // Remove the 2nd node from the end

	head = removeNthFromEnd(head, n)
	printList(head)
}
