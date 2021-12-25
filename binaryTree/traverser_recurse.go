package binarytree

import (
	"fmt"
)

func PrevOrderRecurse(tree *BiTree) {
	if tree.head == nil {
		return
	}
	fmt.Println(prevOrder(tree.head))
}

func prevOrder(node *BiTNode) (ctx string) {
	if node == nil {
		return
	}
	ctx = fmt.Sprintf("%s ", string(node.key.(byte)))

	ctx += prevOrder(node.left)
	ctx += prevOrder(node.right)
	return
}

func InOrderRecurse(tree *BiTree) {
	if tree.head == nil {
		return
	}
	fmt.Println(inOrder(tree.head))

}

func inOrder(node *BiTNode) (ctx string) {
	if node == nil {
		return
	}
	ctx += inOrder(node.left)
	ctx += fmt.Sprintf("%s ", string(node.key.(byte)))
	ctx += inOrder(node.right)
	return
}

func PostOrderRecurse(tree *BiTree) {
	if tree.head == nil {
		return
	}
	fmt.Println(postOrder(tree.head))
}

func postOrder(node *BiTNode) (ctx string) {
	if node == nil {
		return
	}
	ctx += postOrder(node.left)
	ctx += postOrder(node.right)
	ctx += fmt.Sprintf("%s ", string(node.key.(byte)))
	return
}
