package main

import (
	"fmt"
	"unsafe"
)

type link interface {
	addNode(p unsafe.Pointer) error
	//delNode(p unsafe.Pointer) error
	isEmpty() bool
	tail() unsafe.Pointer
	traversal()
	//find()
}

type oneDirectionLink struct {
	head *node
}

type node struct {
	data interface{}
	next *node
}

func (l *oneDirectionLink) isEmpty() bool {
	if l.head == nil {
		return true
	}

	return false
}

func (l *oneDirectionLink) tail() unsafe.Pointer {
	p := l.head
	for {
		if p.next != nil {
			p = p.next
			continue
		}

		return unsafe.Pointer(p)
	}
}

func (l *oneDirectionLink) addNode(p unsafe.Pointer) error {
	if l.isEmpty() {
		l.head = (*node)(p)
		l.head.next = nil
		return nil
	}

	t := (*node)(l.tail())
	t.next = (*node)(p)
	(*node)(p).next = nil

	return nil
}

func (l *oneDirectionLink) traversal() {
	p := l.head
	for {
		if p == nil {
			break
		}

		fmt.Println(p.data)
		p = p.next
	}
}

func testInterface(l link) {
	l.traversal()
}

func main() {
	l := oneDirectionLink{}
	i := 10
	for {
		if i <= 0 {
			break
		}
		p := new(node)
		p.data = i
		//fmt.Println(p)
		l.addNode(unsafe.Pointer(p))
		i--
	}

	var li link
	li = &l
	li.traversal()

	//testInterface(&l)
}
