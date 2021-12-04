package linklist

func DeleteByIndex(l *LinkList, idx int) string {
	pesudoNode := &Node{
		elem: "",
		next: l.head,
	}
	p := pesudoNode
	j := 0
	for p.next != nil && j < idx {
		p = p.next
		j++
	}
	if p.next == nil || j > idx {
		panic("node not exist")
	}
	q := p.next
	p.next = q.next
	elem := q.elem
	q.next = nil
	l.head = pesudoNode.next
	return elem
}

func Append(l *LinkList, e string) {
	pesudoNode := &Node{
		elem: "",
		next: l.head,
	}
	p := pesudoNode
	for p.next != nil {
		p = p.next
	}
	p.next = &Node{
		elem: e,
		next: nil,
	}
	l.head = pesudoNode.next
}

func RecurRevert(l *LinkList) {
	if l.head == nil {
		return
	}

	tail := l.head
	for tail.next != nil {
		tail = tail.next
	}

	pesudoNode := &Node{
		elem: "",
		next: l.head,
	}
	p := pesudoNode

	RRevert(p.next)
	l.head = tail
}

func RRevert(head *Node) {
	if head.next == nil {
		return
	}
	RRevert(head.next)
	head.next.next = head
	head.next = nil
}

func LoopRevert(l *LinkList) {
	if l.head == nil {
		return
	}
	var pre, cur, post *Node
	pre = nil
	cur = l.head

	for cur != nil {
		post = cur.next
		cur.next = pre
		// next
		pre = cur
		cur = post
	}
	l.head = pre
}
