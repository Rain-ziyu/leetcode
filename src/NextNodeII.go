package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	start := root
	//start会进行变换，每次循环start相当于是每一层的第一个元素
	for start != nil {
		//每次循环开始时记录下一层的最左元素
		var nextStart *Node
		var last *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		//从每一行的第一个元素开始进行串联其子元素
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		//	每一次循环完就可以去循环下一层的最左元素了
		start = nextStart

	}
	return root
}
