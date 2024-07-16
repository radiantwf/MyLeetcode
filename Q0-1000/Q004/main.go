package main

import "fmt"

func main() {
	nums1 := []int{1}
	nums2 := []int{}
	ret := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ret)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Len, nums2Len := len(nums1), len(nums2)
	if nums2Len == 0 {
		nums1, nums2 = nums2, nums1
		nums1Len, nums2Len = nums2Len, nums1Len
	} else if nums1Len != 0 && nums1[0] > nums2[0] {
		nums1, nums2 = nums2, nums1
		nums1Len, nums2Len = nums2Len, nums1Len
	}
	if nums1Len == 0 && nums2Len == 1 {
		return float64(nums2[0])
	}
	left, right := 0, 0
	mid := (nums1Len + nums2Len) / 2
	limitLeft, limitRight := max(0, mid-nums2Len), min(nums1Len, mid)
	for {
		nums1Split := limitLeft + (limitRight-limitLeft)/2
		nums2Split := mid - nums1Split

		if nums1Split > 0 && nums2Split < nums2Len && nums1[nums1Split-1] > nums2[nums2Split] {
			limitRight = nums1Split
			continue
		} else if nums1Split < nums1Len && nums2Split > 0 && nums1[nums1Split] < nums2[nums2Split-1] {
			if limitLeft < nums1Split {
				limitLeft = nums1Split
			} else {
				limitLeft++
			}
			continue
		}

		nums1LeftIndex := nums1Split - 1
		nums1RightIndex := nums1Split
		nums2LeftIndex := nums2Split - 1
		nums2RightIndex := nums2Split

		if nums1LeftIndex < 0 {
			left = nums2[nums2LeftIndex]
		} else if nums2LeftIndex < 0 {
			left = nums1[nums1LeftIndex]
		} else {
			left = max(nums1[nums1LeftIndex], nums2[nums2LeftIndex])
		}
		if nums1RightIndex >= nums1Len {
			right = nums2[nums2RightIndex]
		} else if nums2RightIndex >= nums2Len {
			right = nums1[nums1RightIndex]
		} else {
			right = min(nums1[nums1RightIndex], nums2[nums2RightIndex])
		}
		break
	}
	if (nums1Len+nums2Len)%2 == 0 {
		return float64(left+right) / 2.0
	}
	return float64(right)
}
