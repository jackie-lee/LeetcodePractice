# -*- coding:utf8 -*-
#/usr/bin/python
#Given a linked list, swap every two adjacent nodes and return its head.
#
#You may not modify the values in the list's nodes, only nodes itself may be changed.
#
#Example:
#
#Given 1->2->3->4, you should return the list as 2->1->4->3.

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



def swapnodes(lst):
	if None != lst :
		_next = lst._next
	else :
		return lst

	if None == _next:
		return lst

	if None != _next._next:
		lst._next = swapnodes(_next._next)
		_next._next = lst

	return _next

list1 = Node(1)
list1.append(2)
list1.append(3)
list1.append(4)
list1.append(5)

p = swapnodes(list1)
while p:
	print p.data
	p = p._next
