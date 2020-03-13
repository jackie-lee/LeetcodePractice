# -*- coding:utf8 -*-
#/usr/bin/python
#Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
#
#Each number in candidates may only be used once in the combination.
#
#Note:
#
#All numbers (including target) will be positive integers.
#The solution set must not contain duplicate combinations.
#Example 1:
#
#Input: candidates = [10,1,2,7,6,1,5], target = 8,
#A solution set is:
#[
#  [1, 7],
#  [1, 2, 5],
#  [2, 6],
#  [1, 1, 6]
#]
#Example 2:
#
#Input: candidates = [2,5,2,1,2], target = 5,
#A solution set is:
#[
#  [1,2,2],
#  [5]
#]
#
#解题思路：
#考虑使用回溯法解题
#从大到小排序
#从第一个数字开始遍历，若该数字不大于当前目标值，则将其加入到结果数组中
#然后把目标值减去当前数字，并从当前数字开始向后递归寻找下一个满足上述条件的数字
#若到某一步为止目标值为0，则将当前结果数组加入到集合中。
#由于不能使用一个元素多次，所以要建立一个boolean的数组，用于记录每一个是不是都已经被访问了
#从下一个数字开始继续寻找，直到走到数组末尾或者没有不大于目标值的数

#global index # 声明和初始化分行
#index = 0
def findpath(candidates, path, res, target, begin, size):
	if 0 == target:
		#print path
		#global index #使用前再次声明
		#res[index].append(path)
		#res.insert(index, path)
		res.append(path[0:len(path)])
		#index += 1
	else :
		for i in range(begin, size):
			left_num = target - candidates[i]
			if left_num < 0:
				break
			if i > begin and candidates[i] == candidates[i - 1]:
				continue
			path.append(candidates[i])
			findpath(candidates, path, res, left_num, i + 1, size)
			path.pop()

def combinationSum2(candidates, target):
	size = len(candidates)
	if 0 == size :
		return None

	candidates.sort() #还有这种内置函数

	path = []
	res = []
	findpath(candidates, path, res, target, 0, size)

	return res

cdd = [10,1,2,7,6,1,5]

result = combinationSum2(cdd, 8)

print len(result)
print result

