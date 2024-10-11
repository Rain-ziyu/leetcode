package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	tmp := pre
	for pre != nil {
		if pre.Next.Val == 0 {
			pre.Next = pre.Next.Next
			pre = pre.Next
		} else {
			pre.Val += pre.Next.Val
			pre.Next = pre.Next.Next
		}
	}
	return tmp.Next
}

func removeStars(s string) string {
	stack := []uint8{}
	for i := range s {
		if s[i] == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
