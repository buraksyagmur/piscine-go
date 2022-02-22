package main

import (
	"fmt"
	"piscine"
)

func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, "Hello")
	piscine.ListPushBack(link, "man")
	piscine.ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}











// type node struct {
// 	data int
// 	next *node
// }

// func insert(head *node, data int) *node {
// 	n := &node{data: data}
// 	if head == nil {
// 		return n
// 	} else {
// 		n.next = head

// 		return n
// 	}
// }

// func Printlist(head *node) {
// 	for head != nil {
// 		fmt.Print(head.data, "->")
// 		head = head.next

// 	}
// 	fmt.Println(nil)
// }

// func main() {
// 	var link *node
// 	link = insert(link, 1)
// 	// link = insert(link, 2)
// 	Printlist(link)
// }
