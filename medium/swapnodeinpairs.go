/*
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
*/

package main

import "fmt"

type Node struct {
        data int
        next *Node
}

func (self *Node) init(data int) {
        self.data = data
        self.next = nil
}

func (self *Node) append(data int) {
        var n = new(Node)
        n.init(data)

        p := self
        for p.next != nil {
                p = p.next
        }

        p.next = n
}

func swapnode(list *Node) (rlist *Node) {
        var p1 *Node = list

	var p2 *Node
	if nil != p1 {
		p2 = p1.next
	}

	if nil != p2 {
		p1.next = p2.next
		p2.next = p1
		rlist = p2
	} else {
		rlist = p1
	}

	if nil != p1.next {
		var head *Node = p1
		head.next = swapnode(p1.next)
	}

	return rlist
}

func main () {
        var list1 = new(Node)
        list1.init(1)
        list1.append(2)
        list1.append(3)
        list1.append(4)
        list1.append(5)

        var p *Node = swapnode(list1)

        for p != nil {
                fmt.Println(p.data)
                p = p.next
        }

}
