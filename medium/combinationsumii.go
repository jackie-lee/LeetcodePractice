//Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
//Each number in candidates may only be used once in the combination.
//
//Note:
//
//All numbers (including target) will be positive integers.
//The solution set must not contain duplicate combinations.
//Example 1:
//
//Input: candidates = [10,1,2,7,6,1,5], target = 8,
//A solution set is:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
//Example 2:
//
//Input: candidates = [2,5,2,1,2], target = 5,
//A solution set is:
//[
//  [1,2,2],
//  [5]
//]
//
//解题思路：
//考虑使用回溯法解题
//从大到小排序
//从第一个数字开始遍历，若该数字不大于当前目标值，则将其加入到结果数组中
//然后把目标值减去当前数字，并从当前数字开始向后递归寻找下一个满足上述条件的数字
//若到某一步为止目标值为0，则将当前结果数组加入到集合中。
//由于不能使用一个元素多次，所以要建立一个boolean的数组，用于记录每一个是不是都已经被访问了
//从下一个数字开始继续寻找，直到走到数组末尾或者没有不大于目标值的数

package main

import "fmt"

/*
func main(){
	var d2arr [][]int
	d2arr = make([][]int, 3)
	var i, j int
	for i = 0; i < 3; i++ {
		d2arr[i] = make([]int, i + 1)
		for j = 0; j < i + 1; j++ {
			d2arr[i][j] = 0
		}
	}

	fmt.Println(d2arr)
}
*/

func findpath(candidates []int, path []int, res [][]int, target int, begin int, size int) {
	if 0 == target {
		fmt.Println(path)
		res = append(res, path)
	} else {
		for i := begin; i < size; i++ {
			var left_num int = target - candidates[i]
			if left_num < 0 {
				break
			}
			if i > begin && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			findpath(candidates, path, res, left_num, i+1, size)
			path = append(path[:len(path)-1])
		}

	}
}

func combinationSum2(candidates []int, target int, res [][]int) {
	var size int = len(candidates)
	if 0 == size {
		return
	}

	var path []int

	findpath(candidates, path, res, target, 0, size)
}

func main() {
	var candidates = []int{1, 1, 2, 5, 6, 7, 10}

	var result [][]int

	combinationSum2(candidates, 8, result)

	//fmt.Println(result)
}
