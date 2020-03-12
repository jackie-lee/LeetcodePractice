//
//iven n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
//
//ote: You may not slant the container and n is at least 2.
//
//he above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
//
//
//
//xample:
//
//Input: [1,8,6,2,5,4,8,3,7]
//Output: 49
//
// sulotion as containmostwater.py

package main

import(
	"fmt"
	"math"
	)

func maxArea(height []int) (max_area int) {
	max_area = 0
	left := 0
	right := len(height) - 1
	for right > left {
		max_area = int(math.Max(float64(max_area), math.Min(float64(height[left]), float64(height[right])) * float64((right - left))))
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}

	return max_area
}

func main() {
	var arr = []int{1,8,6,2,5,4,8,3,7}

	fmt.Println(maxArea(arr))
}
