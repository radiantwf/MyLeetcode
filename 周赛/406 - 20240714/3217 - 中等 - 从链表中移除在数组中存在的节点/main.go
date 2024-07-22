package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// nums =[1]
	nums := []int{1}
	// head =[1,2,1,2,1,2]
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}}}}
	new := modifiedList(nums, head)
	for new != nil {
		println(new.Val)
		new = new.Next
	}
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	mapNums := make(map[int]bool, len(nums))
	for _, num := range nums {
		mapNums[num] = true
	}
	for head.Next != nil {
		if _, ok := mapNums[head.Val]; !ok {
			break
		}
		head = head.Next
	}
	node := head
	for node.Next != nil {
		if _, ok := mapNums[node.Next.Val]; !ok {
			node = node.Next
		} else {
			node.Next = node.Next.Next
		}
	}
	return head
}
