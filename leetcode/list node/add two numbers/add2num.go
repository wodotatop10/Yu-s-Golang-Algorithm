package main

import (
	"bytes"
	"fmt"
	"strconv"
)

var l1, l2 *ListNode

func main() {
	addBadWay(l1, l2)
	l1.show()
	l2.show()
}

type ListNode struct {
	Num  int
	Next *ListNode
}

func (l *ListNode) add(i int) {
	l.Next = new(ListNode)
	l.Next.Num = i
}

func (l *ListNode) show() int {
	temp := l
	var buf bytes.Buffer
	var i int
	for {
		if temp != nil {
			buf.WriteString(strconv.FormatInt(int64(temp.Num), 10))
			// fmt.Println(temp.Num)
			temp = temp.Next
			i++
		} else {
			break
		}
	}
	fmt.Println(buf.String())
	return i
}

func init() {
	max := 9
	l1 = new(ListNode)
	l2 = new(ListNode)
	l1temp, l2temp := l1, l2
	for i := 9; i > 0; i-- {
		l1temp.Num, l2temp.Num = i, i
		if i != max-1 {
			l1temp.Next, l2temp.Next = new(ListNode), new(ListNode)
			l1temp = l1temp.Next
			l2temp = l2temp.Next
		}

	}
	l1.show()
	l2.show()
}

// func (l *ListNode)

func addBadWay(l1, l2 *ListNode) *ListNode {
	var head *ListNode = new(ListNode)
	list := head
	listnode1, listnode2 := l1, l2
	num := 0
	for {
		if listnode1 != nil || listnode2 != nil {
			var n1, n2 int
			if listnode1 != nil {
				n1 = listnode1.Num

			}
			if listnode2 != nil {
				n2 = listnode2.Num
			}
			list.Num = n1 + n2 + num
			fmt.Println("n1=", n1, ";n2=", n2, ";num=", num)
			num = list.Num / 10
			if num > 0 {
				list.Num = list.Num - 10
			}
			list.show()
			listnode1 = listnode1.Next
			listnode2 = listnode2.Next
			list.Next = new(ListNode)
			list = list.Next
		} else {
			if num > 0 {
				list.Num = num
			}
			break
		}

		if l1.Next == nil && l2.Next == nil {
			if num != 0 {
				list.add(num)
			}
			break
		}
	}

	head.show()
	return head
}
