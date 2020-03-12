# -*- coding:utf8 -*-
#/usr/bin/python
#
#Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
#
#Note: You may not slant the container and n is at least 2.
#
#The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
#
#
#
#Example:
#
#Input: [1,8,6,2,5,4,8,3,7]
#Output: 49
#
#贪心法思路
#我们可以把盛水的容器的两块板分为左板、右板。假设左板高度为h，且比右板低，两块板之间的距离为w，则此时最多能装水，此时我们尝试移动隔板。如果将左板向右移，那么可能使容积变大，例如，左板的右边板子高h1（还是比右板低），此时最多装水，有可能比大；如果将右板向左移，由于水的高度不能高于左板，所以容积最多为，肯定比小。当然，同样的道理，如果右边的板子比左边的低，那就向左移动右边的板子。
#
#基于上面的假设，我们只要把两块隔板依次向中间靠拢，就可以求出最大的容积。

def maxArea(height):
	max_area = 0
	left = 0
	right = len(height) - 1
	while right > left:
		max_area = max(max_area, min(height[left], height[right]) * (right - left))
		if height[right] > height[left]:
			left += 1
		else:
			right -= 1
	return max_area

arr = (1,8,6,2,5,4,8,3,7)

print maxArea(arr)
