package main

import (
	"fmt"
)

func main() {
	// [[2,1,1],[1,2,1]]
	circles := [][]int{{2, 1, 1}, {1, 2, 1}}
	X, Y := 3, 3
	fmt.Println(canReachCorner(X, Y, circles))
}

// 解题思路   （题目有错误，测试case限制了，圆心坐标在矩形内）
// 如果所有圆组合的图形，覆盖面连通了矩形的左上边框与右下边框，则矩形的左下角与右下角不可达
// 依次判断每个圆是否相交或相切（交点与切点需要在矩形上或内，且不能在矩形右上角，与左下角。由于圆心在矩形内，所以暂不检查）
// 依次判断每个圆是否与矩形的左上角或右下角相交或相切，且不能在矩形右上角，与左下角
// 并查集 方式判断【左+上线->圆->右+下线】的连通性

func canReachCorner(X int, Y int, circles [][]int) bool {
	// 初始化并查集，左上线位n，右下线位n+1
	fa := make([]int, len(circles)+2)
	for i := range fa {
		fa[i] = i
	}
	// 查找
	var find func(x int) int
	find = func(x int) int {
		if x != fa[x] {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	var union func(x, y int)
	union = func(x, y int) {
		fa[find(x)] = find(y)
	}
	// 判断是否连通
	var isConnected func(x, y int) bool
	isConnected = func(x, y int) bool {
		return find(x) == find(y)
	}
	circlesLen := len(circles)
	for i, circle := range circles {
		// if isRectangleInsideCircle(circle, X, Y) {
		// 	return false
		// }
		// if isRectangleOutsideCircle(circle, X, Y) {
		// 	continue
		// }
		if isCircleIntersectRectangle(circle, X, Y) == 0x11 {
			return false
		} else if isCircleIntersectRectangle(circle, X, Y) == 0x1 {
			union(i, circlesLen)
		} else if isCircleIntersectRectangle(circle, X, Y) == 0x10 {
			union(i, circlesLen+1)
		}
		for j := i + 1; j < circlesLen; j++ {
			if isIntersectionInsideRectangle(circle, circles[j], X, Y) {
				union(i, j)
			}
		}
	}
	return !isConnected(circlesLen, circlesLen+1)
}

func isCircleIntersectRectangle(circle []int, X, Y int) int {
	ret := 0
	// 检查圆是否与矩形的左上边界相交或相切
	cx, cy, cr := circle[0], circle[1], circle[2]
	if (cx-cr <= 0 && cx >= 0) || (cx+cr >= 0 && cx < 0) ||
		(cy+cr >= Y && cy <= Y) || (cy-cr >= Y && cy > Y) {
		ret |= 0x1
	}
	// 检查圆是否与矩形的右下边界相交或相切
	if (cx+cr >= X && cx <= X) || (cx-cr <= X && cx > X) ||
		(cy-cr <= 0 && cy >= 0) || (cy+cr >= 0 && cy < 0) {
		ret |= 0x10
	}
	return ret
}
func isIntersectionInsideRectangle(c1, c2 []int, X, Y int) bool {
	return (c1[0]-c2[0])*(c1[0]-c2[0])+(c1[1]-c2[1])*(c1[1]-c2[1]) <= (c1[2]+c2[2])*(c1[2]+c2[2])
}

// func isIntersectionInsideRectangle(c1, c2 []int, X, Y int) bool {
// 	d := math.Sqrt(float64((c1[0]-c2[0])*(c1[0]-c2[0]) + (c1[1]-c2[1])*(c1[1]-c2[1])))
// 	radiusSum := float64(c1[2] + c2[2])
// 	radiusDiff := math.Abs(float64(c1[2] - c2[2]))
// 	if d > radiusSum || d < radiusDiff {
// 		return false
// 	}
// 	a := (float64(c1[2]*c1[2]) - float64(c2[2]*c2[2]) + d*d) / (2 * d)
// 	h := math.Sqrt(float64(c1[2]*c1[2]) - a*a)
// 	x2 := float64(c1[0]) + a*(float64(c2[0])-float64(c1[0]))/d
// 	y2 := float64(c1[1]) + a*(float64(c2[1])-float64(c1[1]))/d
// 	rx := -(float64(c2[1]) - float64(c1[1])) * (h / d)
// 	ry := (float64(c2[0]) - float64(c1[0])) * (h / d)
// 	if isPointInsideRectangle([]float64{x2 + rx, y2 + ry}, X, Y) {
// 		return true
// 	} else {
// 		return isPointInsideRectangle([]float64{x2 - rx, y2 - ry}, X, Y)
// 	}
// }

// func isPointInsideRectangle(point []float64, X, Y int) bool {
// 	return point[0] >= 0 && point[0] <= float64(X) && point[1] >= 0 && point[1] <= float64(Y) && !(point[0] == float64(X) && point[1] == float64(Y)) && !(point[0] == 0 && point[1] == 0)
// }

// func _isPointInsideCircle(circle []int, x, y int) bool {
// 	distance := math.Sqrt(float64((circle[0]-x)*(circle[0]-x) + (circle[1]-y)*(circle[1]-y)))
// 	return distance <= float64(circle[2])
// }
// func _isPointOnsideCircle(circle []int, x, y int) bool {
// 	distance := math.Sqrt(float64((circle[0]-x)*(circle[0]-x) + (circle[1]-y)*(circle[1]-y)))
// 	return distance >= float64(circle[2])
// }

// func isRectangleInsideCircle(circle []int, X, Y int) bool {
// 	// 检查矩形的四个顶点是否都在圆内
// 	return _isPointInsideCircle(circle, 0, 0) &&
// 		_isPointInsideCircle(circle, X, 0) &&
// 		_isPointInsideCircle(circle, 0, Y) &&
// 		_isPointInsideCircle(circle, X, Y)
// }

// func isRectangleOutsideCircle(circle []int, X, Y int) bool {
// 	// 检查矩形的四个顶点是否都在圆外
// 	return _isPointOnsideCircle(circle, 0, 0) &&
// 		_isPointOnsideCircle(circle, X, 0) &&
// 		_isPointOnsideCircle(circle, 0, Y) &&
// 		_isPointOnsideCircle(circle, X, Y) &&
// 		(circle[0] < 0 || circle[0] > X || circle[1] < 0 || circle[1] > Y)
// }
