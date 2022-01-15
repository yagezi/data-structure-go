package binarytree

import (
	"fmt"

	"github.com/yagezi/data-structure-go/queue"
	"github.com/yagezi/data-structure-go/stack"
)

func LevelOrderQueue(tree *BiTree) {
	if tree.head == nil {
		return
	}
	str := ""
	q := queue.NewQueue()
	q.Enqueue(tree.head)
	for !q.IsEmpty() {
		n := q.Dequeue().(*BiTNode)
		str += fmt.Sprintf("%s ", string(n.key.(byte)))
		if n.left != nil {
			q.Enqueue(n.left)
		}
		if n.right != nil {
			q.Enqueue(n.right)
		}
	}
	fmt.Println(str)
}

func PreOrderStack(tree *BiTree) {
	str := ""
	s := stack.NewStack()
	cur := tree.head

	for !s.IsEmpty() || cur != nil {
		if cur != nil {
			s.Push(cur)
			str += fmt.Sprintf("%s ", string(cur.key.(byte)))
			cur = cur.left
		} else {
			n := s.Pop()
			cur = n.(*BiTNode).right
		}
	}
	fmt.Println(str)
}

func InOrderStack(tree *BiTree) {
	str := ""
	s := stack.NewStack()
	cur := tree.head

	for !s.IsEmpty() || cur != nil {
		if cur != nil {
			s.Push(cur)
			cur = cur.left
		} else {
			n := s.Pop()
			cur = n.(*BiTNode)
			str += fmt.Sprintf("%s ", string(cur.key.(byte)))
			cur = cur.right
		}
	}
	fmt.Println(str)
}

func PostOrderStack(tree *BiTree) {
	str := ""
	s := stack.NewStack()
	flags := stack.NewStack()
	cur := tree.head

	for !s.IsEmpty() || cur != nil {
		if cur != nil {
			s.Push(cur)
			flags.Push(cur)
			cur = cur.left
		} else {
			n := s.Pop()
			cur = n.(*BiTNode)
			str += fmt.Sprintf("%s ", string(cur.key.(byte)))
			cur = cur.right
		}
	}
	fmt.Println(str)
}
