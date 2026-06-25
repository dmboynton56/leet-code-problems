package main

import "fmt"

// ListNode matches LeetCode's linked list definition.
// Go: struct with pointer next; Python often uses a class or nested ListNode in LeetCode stub.
type ListNode struct {
	Val  int
	Next *ListNode
}

// solution merges two sorted linked lists into one sorted list.
// Go: pointer manipulation; dummy head avoids special-casing the first node.
func solution(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // heap-allocated via &; Python objects are always references.
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next // advance pointer; Python: list1 = list1.next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	// Attach remainder — at most one is non-nil.
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return dummy.Next // skip sentinel node
}

func buildList(vals ...int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	out := []int{}
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func main() {
	l1 := buildList(1, 2, 4)
	l2 := buildList(1, 3, 4)
	fmt.Println(listToSlice(solution(l1, l2))) // [1 2 3 4 4]
}
