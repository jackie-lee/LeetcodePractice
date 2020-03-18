#!/usr/bin/python
# -*- coding: UTF-8 -*-
#Given a 2D board and a word, find if the word exists in the grid.
#
#The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.
#
#Example:
#
#board =
#[
#  ['A','B','C','E'],
#  ['S','F','C','S'],
#  ['A','D','E','E']
#]
#
#Given word = "ABCCED", return true.
#Given word = "SEE", return true.
#Given word = "ABCB", return false.
#
#Solution
#This problem does not give start position, or direction restriction, so
#
#1.Scan board, find starting position with matching word first letter
#2.From starting position, DFS (4 (up, down, left, right 4 directions) match word's rest letters
#3.For each visited letter, mark it as visited, here use board[i][j] = '*' to represent visited.
#4.If one direction cannot continue, backtracking, mark start position unvisited, mark board[i][j] = word[start]
#5.If found any matching, terminate
#6.Otherwise, no matching found, return false.

def exist(board, word):
	m = len(board)
	n = len(board[0])

	def dfs(board, r, c, word, index):
		if index == len(word):
			return True
		if r < 0 or r >=m or c < 0 or c >= n or board[r][c] != word[index]:
			return False
		board[r][c] = '*'
		res = dfs(board, r - 1, c, word, index + 1) \
			or dfs(board, r + 1, c, word, index + 1) \
			or dfs(board, r, c - 1, word, index + 1) \
			or dfs(board, r, c + 1, word, index)
		board[r][c] = word[index]
		return res
	
	for r in range(m):
		for c in range(n):
			if board[r][c] == word[0]:
				if dfs(board, r, c, word, 0):
					return True

board = [
	['A', 'B', 'C', 'E'],
	['S', 'F', 'C', 'S'],
	['A', 'D', 'E', 'E']
]

print exist(board, "SFE")
