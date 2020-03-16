# -*- coding:utf8 -*-
#/usr/bin/python
#Given an array of non-negative integers, you are initially positioned at the first index of the array.
#
#Each element in the array represents your maximum jump length at that position.
#
#Determine if you are able to reach the last index.
#
#Example 1:
#
#Input: [2,3,1,1,4]
#Output: true
#Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
#Example 2:
#
#Input: [3,2,1,0,4]
#Output: false
#Explanation: You will always arrive at index 3 no matter what. Its maximum
#             jump length is 0, which makes it impossible to reach the last index.
#
#解决思路：
#这道题目是一道典型的回溯类型题目。 思路就是用一个变量记录当前能够到达的最大的索引，我们逐个遍历数组中的元素去更新这个索引。 变量完成判断这个索引是否大于数组下表即可。
#
#解释起来很简单，但只依靠自己能不能想得这样的规律，想出这样的解法？

def jumpgame(steps):
	Max = 0
	Len = len(steps)
	for i in range(Len - 1):
		if Max < i:
			return False
		Max = max(Max, steps[i] + i)
	
	return Max >= Len - 1

arr = [2,2,0,1,4]

print jumpgame(arr)
