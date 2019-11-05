/* er cha shu

 */
package main

import (
	"fmt"
)

type jiekou interface {
	Start()
}

type A struct {
	Name string
	jiekou
}

type Node struct {
	Left  *Node
	Val   int
	Right *Node
}

func show(tree *Node) {

	if tree.Left != nil {
		show(tree.Left)
	}
	fmt.Println(tree.Val)
	if tree.Right != nil {
		show(tree.Right)
	}

}

func showNodeTree(tree *Node) {
	if tree == nil {
		fmt.Println("ok")
	} else {
		show(tree)
	}
}

func createNodeTree(tree *Node, a *Node) {

	if tree.Val < a.Val {
		if tree.Left == nil {
			tree.Left = a
		} else {
			createNodeTree(tree.Left, a)
		}

	} else {
		if tree.Right == nil {
			tree.Right = a
		} else {
			createNodeTree(tree.Right, a)
		}

	}
}

func setNode(n1 *Node, data Node) {
	if n1 == nil {
		fmt.Println("create nill")
		n1 = &data
		first = n1
		fmt.Println("create nill ", n1.Val)
	} else {
		fmt.Println(n1.Val)
	}
}

var first *Node

func main() {
	var listData []Node

	var data = []int{1, 3, 10, 7, 80, 99, 2, 6}
	for i := 0; i < len(data); i++ {
		listData = append(listData, Node{Val: data[i]})
		fmt.Println(listData[i].Val)
	}
	first = &listData[0]
	for i := 1; i < len(listData); i++ {

		createNodeTree(first, &listData[i])
	}
	fmt.Println("\nprint")

	showNodeTree(first)

}
