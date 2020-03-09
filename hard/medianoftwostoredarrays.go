/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

解题思路1：
二分法，找到了A[n/2] 和 B[n/2]来比较，
1、如果他们相等，那样的话，我们的搜索结束了，因为答案已经找到了A[n/2]就肯定是排序后的中位数了。
2、如果我们发现B[n/2] > A[n/2]，说明什么，这个数字应该在 A[n/2]->A[n]这个序列里面， 或者在 B[1]-B[n/2]这里面。 或者，这里的或者是很重要的， 我们可以说，我们已经成功的把问题变成了在排序完成的数组A[n/2]-A[n]和B[0]-B[n/2]里面找到合并以后的中位数， 显然递归是个不错的选择了。
3、如果B[n/2] < A[n/2]呢？显然就是在A[0]-A[n/2]和B[n/2]-B[n]里面寻找了。
*/

package main

import "fmt"

func medianofarrays(arr1 []int, len1 int, arr2 []int, len2) (median float32) {
	mid1 := len1 / 2
	mid2 := len2 / 2

	var mid int
	if mid1 <= midb {
		mid = mid1
	} else {
		mid = mid2
	}

	if 1 == len1 {
		if 0 == len2 % 2 {
			if arr1[0] > b[mid2 - 1] {
				return b[mid2]
			} else if arr1[0] < b[mid2 - 1]
				return b[mid2 - 1]

			return arr1[0]
		} else {
			return b[mid2]
		}
	}


	if 1 == len2 {
		if 0 == len1 % 2 {
			if arr2[0] > arr1[mid1 - 1] {
				return arr1[mid1]
			} else if arr2[0] < arr1[mid1 - 1]
				return arr1[mid1 - 1]

			return arr2[0]
		} else {
			return arr1[mid1]
		}
	}

	if arr1[mid1] == arr2[mid2] {
		return arr1[mid1]
	} else if arr1[mid1] < arr2[mid2] {
		return medianofarrays(arr1[mid1], len1 - mid, arr2, len2 - mid)
	} else {
		retrun medianofarrays(arr1, len1 - mid, arr2[mid2], len2 - mid)
	}
	}
	return median
}

func main () {
}
