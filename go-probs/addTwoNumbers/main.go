package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// solution adds two numbers represented as reversed linked lists.
// Go: carry as int, modulo/division; nil-safe with helper or inline checks.
func solution(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		carry = sum / 10 // Go int division truncates toward zero — fine for positive sums
	}
	return dummy.Next
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
	l1 := buildList(2, 4, 3) // 342
	l2 := buildList(5, 6, 4) // 465
	fmt.Println(listToSlice(solution(l1, l2))) // [7 0 8] => 807
}
