package LRUCache

type Node struct {
	Next  *Node
	Prev  *Node
	Value interface{}
}

type List struct {
	Head *Node
	Tail *Node
}

func NewList() *List {
	return &List{}
}

func (l *List) PushBack(v interface{}) {
	newTail := &Node{
		Value: v,
	}
	if l.Head == nil {
		l.Head = newTail
		l.Tail = l.Head
		return
	}
	l.Tail.Next = newTail
	newTail.Prev = l.Tail
	l.Tail = newTail
}

func (l *List) PushFront(v interface{}) {
	newHead := &Node{
		Value: v,
	}
	if l.Head == nil {
		l.Head = newHead
		l.Tail = l.Head
		return
	}
	newHead.Next = l.Head
	l.Head.Prev = newHead
	l.Head = newHead
}

func (l *List) PopFront() interface{} {
	value := l.Head.Value
	n := l.Head.Next
	l.Head = n
	return value
}

func (l *List) PopBack() {
	n := l.Tail.Prev
	n.Next = nil
	l.Tail = n
}

func (l *List) Delete(v interface{}) {
	if l.Head.Value == v {
		l.PopFront()
	} else if l.Tail.Value == v {
		l.PopBack()
	} else {
		cur := l.Head
		for cur != nil {
			if cur.Value == v {
				prev := cur.Prev
				next := cur.Next
				prev.Next = next
				next.Prev = prev
				return
			}
			cur = cur.Next
		}
	}
}
