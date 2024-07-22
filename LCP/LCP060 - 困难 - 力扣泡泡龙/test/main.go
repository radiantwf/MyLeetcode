package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
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

// LevelInfo 用于存储每层的信息，包括前缀和、是否被占用以及左、右子节点的指针。
type LevelInfo struct {
	PrefixSum, IsOccupied, LeftPtr, RightPtr int
}

// RemoveNode 用于存储需要被移除的节点信息。
type RemoveNode struct {
	NodeIndex int
	Level     int
}

// levelInfos 用于存储每一层的信息。
var levelInfos [][]LevelInfo

// removeNodes 用于存储需要被移除的节点信息。
var removeNodes []RemoveNode

func collectNodes(level int, node *TreeNode) int {
	if node == nil {
		return 0
	}

	// 扩展 levelInfos 列表以适应新层。
	for len(levelInfos) <= level+1 {
		levelInfos = append(levelInfos, []LevelInfo{{PrefixSum: 0, IsOccupied: -1, LeftPtr: -1, RightPtr: -1}})
	}

	// 更新当前层的信息。
	lastInfo := levelInfos[level][len(levelInfos[level])-1]
	newInfo := LevelInfo{
		PrefixSum:  lastInfo.PrefixSum + node.Val, // 计算前缀和
		IsOccupied: -1,                            // 初始化为未占用
		LeftPtr:    len(levelInfos[level+1]),      // 设置左子节点的指针
		RightPtr:   -1,                            // 初始化右子节点的指针为-1
	}
	levelInfos[level] = append(levelInfos[level], newInfo)

	// 用于识别节点的索引。
	node.Val = len(levelInfos[level]) - 1

	// 如果节点有两个子节点，则添加到 removeNodes 列表中。
	if collectNodes(level+1, node.Left)+collectNodes(level+1, node.Right) != 2 {
		removeNodes = append(removeNodes, RemoveNode{NodeIndex: len(levelInfos[level]) - 1, Level: level})
	}

	// 更新右子节点的指针。
	levelInfos[level][len(levelInfos[level])-1].RightPtr = len(levelInfos[level+1]) - 1

	return 1
}

// getMaxLayerSum 计算树的最大层和。
func getMaxLayerSum(root *TreeNode) int {
	collectNodes(0, root)

	// 获取树的高度。
	treeHeight := len(levelInfos) - 1

	// 初始化最大层和。
	maxSum := math.MinInt32

	// 计算没有移除任何节点时的最大层和。
	for level := 0; level < treeHeight; level++ {
		maxSum = max(maxSum, levelInfos[level][len(levelInfos[level])-1].PrefixSum)
	}

	// 遍历 removeNodes 列表，并尝试移除每一个节点以计算新的最大层和。
	for _, removeInfo := range removeNodes {
		nodeIndex, startLevel := removeInfo.NodeIndex, removeInfo.Level
		left, right := nodeIndex, nodeIndex
		// 计算被移除节点的值。
		lostVal := levelInfos[startLevel][left].PrefixSum - levelInfos[startLevel][left-1].PrefixSum

		for level := startLevel; level < treeHeight; level++ {
			if left > right {
				break
			}

			leftInfo, rightInfo := &levelInfos[level][left], &levelInfos[level][right]
			// 如果节点已被标记为占用，则跳过。
			if leftInfo.IsOccupied != -1 && leftInfo.IsOccupied == rightInfo.IsOccupied {
				break
			}

			// 标记节点为占用。
			leftInfo.IsOccupied, rightInfo.IsOccupied = nodeIndex, nodeIndex

			addVal := 0
			// 计算被移除节点的子节点和。
			if leftInfo.LeftPtr <= rightInfo.RightPtr {
				addVal = levelInfos[level+1][rightInfo.RightPtr].PrefixSum - levelInfos[level+1][leftInfo.LeftPtr-1].PrefixSum
			}

			// 计算新的层和。
			newSum := levelInfos[level][len(levelInfos[level])-1].PrefixSum - lostVal + addVal
			// 如果新的层和不为0，或者该层的节点数不等于总节点数，则更新最大层和。
			if newSum != 0 || (right-left+1 != len(levelInfos[level])-1) {
				maxSum = max(maxSum, newSum)
			}

			// 更新 left 和 right 指针，以及需要减去的值。
			left, right, lostVal = leftInfo.LeftPtr, rightInfo.RightPtr, addVal
		}
	}

	// 返回计算得出的最大层和。
	return maxSum
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
