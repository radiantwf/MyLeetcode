package main

import (
	"log"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解题思路：
// 1、 遍历二叉树
// 1.1、 记录每一层的层和
// 1.2、 找出可以击破的节点(子节点<=1),记录这个节点，和这个节点所在的层数，值，以及这个节点下子树每个层的层和
// 2、 计算二叉树的最大层和
// 3、 遍历记录的可以击破的节点，计算击破这个节点后的最大层和，与之前的最大层和比较，取最大值
// 注意：击破节点后，需要重新计算这个节点下的子树的层和，因为这个节点的子树的层和会发生变化（方法为每层层和-记录的节点子树的当前层层和+子树的下层层和）

// BUG:
// 1、当击破节点是末尾节点时，该节点是这层唯一节点，导致本行层和为0，若0大于其他所有层层和，会发生错误
// 2、当击破节点是首节点时，导致最后一层上移，程序逻辑会认为最后一层层和变为0，若0大于其他所有层层和，会发生错误
// 3、当击破节点是层唯一节点时，导致最后一层上移，程序逻辑会认为最后一层层和变为0，若0大于其他所有层层和，会发生错误
// 4、当击破节点后，若击破的节点的子树，深度大于其他旁支子树，择会导致最后一层的层和计算错误结果为0，会发生错误
// 后续遍历 二叉树

func main() {
	// [6,0,3,null,8]	11
	// test_case := &TreeNode{Val: 6, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 8}}, Right: &TreeNode{Val: 3}}
	// [-1,null,-2,null,-3,null,-4,null,-5,null,-6]
	// test_case := &TreeNode{Val: -1, Right: &TreeNode{Val: -2, Right: &TreeNode{Val: -3, Right: &TreeNode{Val: -4, Right: &TreeNode{Val: -5, Right: &TreeNode{Val: -6}}}}}}
	// [-6449,-3776,null,-1138,-4328,null,-6582,null,-7065,null,-1034,-3184,-5485]
	// test_case := &TreeNode{Val: -6449,
	// 	Left: &TreeNode{Val: -3776,
	// 		Left: &TreeNode{Val: -1138,
	// 			Left: nil,
	// 			Right: &TreeNode{Val: -6582,
	// 				Left:  nil,
	// 				Right: &TreeNode{Val: -1034}}},
	// 		Right: &TreeNode{Val: -4328,
	// 			Left: nil,
	// 			Right: &TreeNode{Val: -7065,
	// 				Left:  &TreeNode{Val: -3184},
	// 				Right: &TreeNode{Val: -5485}}}},
	// 	Right: nil}
	// [-7612,-9455,null,-1938,null,null,-4422,-8950,-4982,-4727,null,-6602,null,-2899,null,null,-9071]
	// test_case := &TreeNode{Val: -7612,
	// 	Left: &TreeNode{Val: -9455,
	// 		Left: &TreeNode{Val: -1938,
	// 			Left: nil,
	// 			Right: &TreeNode{Val: -4422,
	// 				Left: &TreeNode{Val: -8950,
	// 					Left: &TreeNode{Val: -4727,
	// 						Left: &TreeNode{Val: -2899}},
	// 					Right: nil},
	// 				Right: &TreeNode{Val: -4982,
	// 					Left: &TreeNode{Val: -6602,
	// 						Left:  nil,
	// 						Right: &TreeNode{Val: -9071}},
	// 					Right: nil}}}}}
	// [-3,0,-5,-3,null,null,-4,-1,null,null,-6,-4]
	test_case := &TreeNode{Val: -3,
		Left: &TreeNode{Val: 0,
			Left: &TreeNode{Val: -3,
				Left: &TreeNode{Val: -1,
					Left:  &TreeNode{Val: -4},
					Right: nil},
				Right: nil},
			Right: nil},
		Right: &TreeNode{Val: -5,
			Left: nil,
			Right: &TreeNode{Val: -4,
				Left:  nil,
				Right: &TreeNode{Val: -6}}}}

	println(getMaxLayerSum(test_case))
}

// import "math"
var laySumList [][]*int
var canBreakNodeTreeMap map[int]map[int]bool

func getMaxLayerSum(root *TreeNode) int {
	laySumList = make([][]*int, 0)
	canBreakNodeTreeMap = make(map[int]map[int]bool)

	postOrderTraversal(root, 0, nil)
	maxLayerSum := math.MinInt32
	maxLayer := len(laySumList[0]) - 1

	for index := 0; index < len(laySumList); index++ {
		var top, bottom int
		if index == 0 {
			top = 0
			bottom = maxLayer
		} else {
			top = *laySumList[index][0]
			bottom = top + len(laySumList[index]) - 2

			max := 0
			subTree := canBreakNodeTreeMap[index]

			for i := 1; i < len(laySumList); i++ {
				if i == index {
					continue
				}
				if _, ok := subTree[i]; ok {
					continue
				}
				tmp := *laySumList[i][0] + len(laySumList[i]) - 2
				if max < tmp {
					max = tmp
				}
			}
			if max < bottom {
				bottom = max
			}
		}
		for l := bottom; l >= top; l-- {
			var layerSum int
			if index == 0 {
				layerSum = *laySumList[0][l]
			} else {
				if l > 0 && top == 0 {
					continue
				}

				tmp := l - top + 1
				layerSum = *laySumList[0][l] - *laySumList[index][tmp]
				if tmp+1 <= len(laySumList[index])-1 {
					layerSum = layerSum + *laySumList[index][tmp+1]
				}
			}
			if layerSum > maxLayerSum {
				maxLayerSum = layerSum
			}
			if layerSum == 0 {
				log.Println("LayerSum", layerSum)
			}
		}
	}

	return maxLayerSum
}

func postOrderTraversal(root *TreeNode, layer int, SubTreesLayerSumIndex []int) {
	if root == nil {
		return
	}

	if layer == 0 {
		laySumList = append(laySumList, make([]*int, 0))
	}

	canBreak := false
	currentBreakNodeIndex := -1
	if (root.Left == nil || root.Right == nil) && layer > 0 {
		canBreak = true
	}

	if canBreak {
		subTreeLayerSum := []*int{&layer}
		laySumList = append(laySumList, subTreeLayerSum)
		index := len(laySumList) - 1
		currentBreakNodeIndex = index
		if SubTreesLayerSumIndex == nil {
			SubTreesLayerSumIndex = make([]int, 0)
		}
		SubTreesLayerSumIndex = append(SubTreesLayerSumIndex, index)
	}

	postOrderTraversal(root.Left, layer+1, SubTreesLayerSumIndex)
	postOrderTraversal(root.Right, layer+1, SubTreesLayerSumIndex)
	value := root.Val

	for _, index := range SubTreesLayerSumIndex {
		subTreeRootLayer := *laySumList[index][0]
		l := layer - subTreeRootLayer + 1
		if len(laySumList[index]) <= l {
			laySumList[index] = append(laySumList[index], make([]*int, l-len(laySumList[index])+1)...)
		}
		if laySumList[index][l] == nil {
			laySumList[index][l] = &value
			// log.Println("add laySumList", index, l, root.Val, *laySumList[index][l])
		} else {
			tmp := *laySumList[index][l] + root.Val
			laySumList[index][l] = &tmp
			// log.Println("add laySumList", index, l, root.Val, *laySumList[index][l])
		}
	}
	if currentBreakNodeIndex != -1 {
		for _, index := range SubTreesLayerSumIndex {
			if index == currentBreakNodeIndex {
				continue
			}
			if _, ok := canBreakNodeTreeMap[index]; !ok {
				canBreakNodeTreeMap[index] = make(map[int]bool)
			}
			if _, ok := canBreakNodeTreeMap[currentBreakNodeIndex]; !ok {
				canBreakNodeTreeMap[currentBreakNodeIndex] = make(map[int]bool)
			}
			canBreakNodeTreeMap[index][currentBreakNodeIndex] = true
			canBreakNodeTreeMap[currentBreakNodeIndex][index] = true
		}

	}

	if len(laySumList[0]) <= layer {
		laySumList[0] = append(laySumList[0], make([]*int, layer-len(laySumList[0])+1)...)
	}
	if laySumList[0][layer] == nil {
		laySumList[0][layer] = &value
	} else {
		tmp := *laySumList[0][layer] + root.Val
		laySumList[0][layer] = &tmp
	}

}
