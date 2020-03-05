# Given a string containing just the characters '(',')','{','}','[',']', determin if the input string is valid.

#An input string is valid if:

#Open brackets must be closed by the same of brackets
#open brackets must be closed by the correct order.
#Note that an empty string is also considerd valid

#Example 1:
#Input: "()"
#Output: true

#Example 2:
#Input: "()[]{}"
#Output: true

#Example 3:
#Input: "(]"
#Output: false

#Example 4:
#Input: "([)]"
#Output: failse

#Example 5:
#Input: "{[]}"
#Output: true
#import sys
#for i in range(1, len(sys.argv)):
#	print sys.argv[i]

import sys

#def isValid():
stack = []
map = {"{":"}","[":"]","(":")"} # dist
#for x in sys.argv[1]:
s = ["{","{","]","}"] # list
for x in s:
	print x
	if x in map:
		stack.append(map[x])
	else:
		if len(stack) != 0:
			top_element = stack.pop()
			if x != top_element:
				print False
				break
			else:
				continue
		else:
			print False
			break
print len(stack) == 0

#isValid(sys.argv[1])
