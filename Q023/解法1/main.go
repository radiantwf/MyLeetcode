package main

import "fmt"

// 定义链表节点结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	lists := []*ListNode{
		{
			Val: 0,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	}
	ret := mergeKLists(lists)
	fmt.Println(ret)
}

func mergeKLists(lists []*ListNode) *ListNode {
	newLists := make([]*ListNode, 0)
	for _, list := range lists {
		if list != nil {
			newLists = append(newLists, list)
		}

	}
	if len(newLists) == 0 {
		return nil
	}
	// 对链表头节点进行排序
	listsLen := len(newLists)
	quickSort(newLists, 0, len(newLists)-1)

	topListNode := newLists[0]
	minIndex := 0
	currentNode := &ListNode{
		Val:  0,
		Next: topListNode}
	for {
		if minIndex == listsLen {
			break
		}
		currentNode.Next = newLists[minIndex]
		currentNode = currentNode.Next

		if currentNode.Next == nil {
			minIndex++
			continue
		}

		x, y := minIndex+1, listsLen-1
		if y >= x {
			for x < y {
				mid := x + (y-x)/2
				if newLists[mid].Val < currentNode.Next.Val {
					x = mid + 1
				} else if newLists[mid].Val == currentNode.Next.Val {
					x = mid
					break
				} else {
					y = mid
				}
			}
			if newLists[x].Val > currentNode.Next.Val {
				x--
			}
			for i := minIndex; i < x; i++ {
				newLists[i] = newLists[i+1]
			}
			newLists[x] = currentNode.Next
		} else {
			newLists[minIndex] = currentNode.Next
		}
	}
	return topListNode
}
func quickSort(lists []*ListNode, low, high int) {
	if low < high {
		pi := partition(lists, low, high)
		quickSort(lists, low, pi-1)
		quickSort(lists, pi+1, high)
	}
}

func partition(lists []*ListNode, low, high int) int {
	pivot := lists[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if lists[j].Val < pivot.Val {
			i++
			lists[i], lists[j] = lists[j], lists[i]
		}
	}
	lists[i+1], lists[high] = lists[high], lists[i+1]
	return i + 1
}
