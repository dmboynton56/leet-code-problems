package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// solution returns true if the linked list has a cycle.
func solution(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	// Cycle: tail connects to index 1
	n3 := &ListNode{Val: 0}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 2, Next: n2}
	head := &ListNode{Val: 3, Next: n1}
	n3.Next = n1
	fmt.Println(solution(head)) // true

	// No cycle
	single := &ListNode{Val: 1}
	fmt.Println(solution(single)) // false
}
