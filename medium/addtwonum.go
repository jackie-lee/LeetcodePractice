//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

//You may assume the two numbers do not contain any leading zero, except the number 0 itself.

//Example
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

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

func addList(n1 *Node, n2 *Node, carry int) (p *Node) {
	// fmt.Printf("%x %x %v\n", n1, n2, carry)
	if nil != n1 && nil != n2 {
		carry += n1.data + n2.data
	} else if nil != n1 {
		carry += n1.data
	} else if nil != n2 {
		carry += n2.data
	}

	// if 0 != carry % 10 {	// be carefull, i will be a hiding bug
	if 0 != carry {
		p = new(Node)
		p.init(carry % 10)
	}

	if nil != n1 && nil != n2 {
		p.next = addList(n1.next, n2.next, carry / 10)
	} else {
		if nil != n1 {
			p.next = addList(n1.next, nil, carry / 10)
		} else if nil != n2 {
			p.next = addList(nil, n2.next, carry / 10)
		} else {
			if 0 != (carry / 10) {
				p.next = new(Node)
				p.next.init(carry / 10)
			}
		}
	}

	return p
}

func main() {
	var list1 = new(Node)
	list1.init(2)
	list1.append(4)
	list1.append(3)

	var list2 = new(Node)
	list2.init(5)
	list2.append(6)
	list2.append(7)

	var p1 *Node
	p1 = list1

	fmt.Println("list1:")
	for p1 != nil {
		fmt.Println(p1.data)
		p1 = p1.next
	}

	p1 = list2

	fmt.Println("list2:")
	for p1 != nil {
		fmt.Println(p1.data)
		p1 = p1.next
	}


	p1 = addList(list1, list2, 0)

	fmt.Println("result:")
	for p1 != nil {
		fmt.Println(p1.data)
		p1 = p1.next
	}
}
