package linktable

import "reflect"

type Node struct {
	prev, next *Node
	linkTable  *LinkTable
	Val        any
}

func (n *Node) Next() *Node {
	if n.linkTable == nil {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.linkTable == nil {
		return nil
	}
	return n.prev
}

type LinkTable struct {
	head, tail *Node
	len        int
	valType    any
}

func Create() *LinkTable {
	linkTable := &LinkTable{
		head:    nil,
		tail:    nil,
		len:     0,
		valType: nil,
	}
	return linkTable
}

func (l *LinkTable) Front() *Node {
	return l.head
}

func (l *LinkTable) Back() *Node {
	return l.tail
}

func (l *LinkTable) Len() int {
	return l.len
}

func (l *LinkTable) PushFront(v any) *Node {
	n := &Node{
		prev:      nil,
		next:      nil,
		linkTable: l,
		Val:       v,
	}
	if l.head == nil && l.tail == nil {
		l.head = n
		l.tail = n
		l.valType = v
	} else {
		if reflect.TypeOf(v) != reflect.TypeOf(l.valType) {
			return nil
		}
		n.next = l.head
		l.head.prev = n
		l.head = n
	}
	l.len++
	return n
}

func (l *LinkTable) PushBack(v any) *Node {
	n := &Node{
		prev:      nil,
		next:      nil,
		linkTable: l,
		Val:       v,
	}
	if l.head == nil && l.tail == nil {
		l.head = n
		l.tail = n
		l.valType = v
	} else {
		if reflect.TypeOf(v) != reflect.TypeOf(l.valType) {
			return nil
		}
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
	l.len++
	return n
}

func (l *LinkTable) Remove(n *Node) any {
	if n == nil {
		return nil
	}

	if n.linkTable == l {
		switch {
		case l.head == n && l.tail == n:
			l.head = nil
			l.tail = nil
		case l.head == n:
			l.head = l.head.next
			n.next.prev = nil
		case l.tail == n:
			l.tail = l.tail.prev
			n.prev.next = nil
		default:
			n.prev.next = n.next
			n.next.prev = n.prev
		}
		n.prev = nil
		n.next = nil
		n.linkTable = nil
		l.len--
	}
	return n.Val
}

func (l *LinkTable) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
	l.valType = nil
}
