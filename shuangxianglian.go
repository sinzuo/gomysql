/*  xiaoxi fenfa xitong
shuangxiang lianbiao
*/

package main

import (
	"fmt"
)

type Node struct {
	pre  *Node
	Val  int
	next *Node
}

func main() {
	var a = []int{1, 80, 7, 6, 10, 20}

	var p, p1, p2 *Node
	p = &Node{Val: a[0]}
	p2 = p
	for i := 1; i < len(a); i++ {
		p1 = &Node{Val: a[i]}
		for p3 := p; p3 != nil; p3 = p3.next {
			if p1.Val < p3.Val {
				p1.next = p3
				p1.pre = nil
				p3.pre = p1
				p = p1
			}
			if p1.Val > p3.Val && p3.next == nil {
				p3.next = p1
				p1.pre = p3
				p1.next = nil
			}
			if p1.Val > p3.Val && p1.Val < p3.next.Val {
				p4 := p3.next
				p3.next = p1
				p1.pre = p3
				p1.next = p4
				p4.pre = p3
				break
			}
		}
	}
	for p3 := p2; p3 != nil; p3 = p3.next {
		fmt.Println(p3.Val)
	}

	fmt.Println("xiao xi fenfa xitong qidong")
}
