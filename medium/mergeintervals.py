# -*- coding:utf8 -*-
#/usr/bin/python
#Given a collection of intervals, merge all overlapping intervals.
#
#Example 1:
#
#Input: [[1,3],[2,6],[8,10],[15,18]]
#Output: [[1,6],[8,10],[15,18]]
#Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
#Example 2:
#
#Input: [[1,4],[4,5]]
#Output: [[1,5]]
#Explanation: Intervals [1,4] and [4,5] are considered overlapping.
#NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
#解决思路：
#１.以每一项的第一个元素为键排序
#２.res最后一项的后一个元素与新加入项的前一个元素进行比较
#如果最后项的后元素比新加入项的前元素要小，那么这两项并没有交集，将这个新项加入res
#如果最后项的后元素大于等于新项的前元素，那么我们就比较二者的后元素，选择后元素大的合并项的后元素，将二者合并。

def merge(intervals):
	if len(intervals) <= 1:
		return intervals

	def get_first(a_list):
		return a_list[0]

	intervals.sort(key=get_first)

	res = [intervals[0]]
	for i in range(1, len(intervals)):
		if intervals[i][0] <= res[-1][1]:
			res[-1] = [res[-1][0], max(res[-1][1], intervals[i][1])]
		else:
			res.append(intervals[i])
	return res


iv = [[1,6], [2,5], [8,10]]

print merge(iv)

