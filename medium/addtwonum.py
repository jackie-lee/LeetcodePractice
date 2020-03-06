# -*- coding:utf8 -*-
#/usr/bin/python
#You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

#You may assume the two numbers do not contain any leading zero, except the number 0 itself.

#Example
#Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
#Output: 7 -> 0 -> 8
#Explanation: 342 + 465 = 807.


class Node(object):

	def __init__(self, value, pnext = None): # watch out, do not forget the "pnext" input
		self.data = value
		self._next = pnext

	def __repr__(self):
		return str(self.data)

	def append(self, dataOrNode):
		item = None
		if isinstance(dataOrNode, Node):
			item = dataOrNode
		else:
			item = Node(dataOrNode)

		node = self
		while node._next:
			node = node._next
		node._next = item

list1 = Node(2)
list1.append(4)
list1.append(3)

p1 = list1
print("l1:")
while p1:
	print(p1.data)
	p1 = p1._next

list2 = Node(5)
list2.append(6)
list2.append(4)

p2 = list2
print("l2:")
while p2:
	print(p2.data)
	p2 = p2._next


def addTwoList(n1, n2, carry):
	carry += (0 if (None == n1) else n1.data) + (0 if (None == n2) else n2.data)
	ret = Node(carry % 10)
	
	#ret = addTwoList(n1._next, n2._next, carry / 10)
	t1 = n1._next
	t2 = n2._next
	carry = carry / 10
	#ret = addTwoList(t1, t2, carry)
	print t1.data, t2.data
	
	#addTwoList(t1, t1, carry)

	return ret

result = addTwoList(list1, list2, 0)

p3 = result
print("result:")
while p3:
	print(p3.data)
	p3 = p3._next
