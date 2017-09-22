package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// BiTNode 二叉树的节点
type BiTNode struct {
	data           compare
	lchild, rchild *BiTNode
}

// Tree 二叉树
type Tree struct {
	root *BiTNode
}

type compare interface {
	less(data compare) bool
	equal(data compare) bool
}

type myInt struct {
	number int
}

func (m *myInt) equal(data compare) bool {
	if m.number == data.(*myInt).number {
		return true
	}
	return false
}

func (m *myInt) less(data compare) bool {
	if m.number < data.(*myInt).number {
		return true
	}

	return false
}

func (b *BiTNode) search(data compare) (*BiTNode, bool) {
	if b == nil {
		return nil, false
	}

	if data.equal(b.data) {
		return b, true
	} else if data.less(b.data) {
		return b.lchild.search(data)
	} else {
		return b.rchild.search(data)
	}
}

func (b *BiTNode) insert(data compare) error {
	if b == nil {
		return errors.New("can not insert data to a nil tree")
	}

	if data.equal(b.data) {
		return nil
	} else if data.less(b.data) {
		if b.lchild == nil {
			b.lchild = &BiTNode{data: data}
			return nil
		}
		return b.lchild.insert(data)
	} else {
		if b.rchild == nil {
			b.rchild = &BiTNode{data: data}
			return nil
		}
		return b.rchild.insert(data)
	}
}

func (b *BiTNode) findMax(parent *BiTNode) (*BiTNode, *BiTNode) {
	if b.rchild == nil {
		return b, parent
	}

	return b.rchild.findMax(b)
}

func (b *BiTNode) replaceNode(parent, replacement *BiTNode) error {
	if b == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if b == parent.lchild {
		parent.lchild = replacement
		return nil
	}

	parent.rchild = replacement
	return nil
}

func (b *BiTNode) delete(data compare, parent *BiTNode) error {
	if b == nil {
		return errors.New("not found")
	}

	if data.equal(b.data) {
		if b.lchild == nil && b.rchild == nil {
			b.replaceNode(parent, nil)
			return nil
		}

		if b.lchild == nil {
			b.replaceNode(parent, b.rchild)
			return nil
		}

		if b.rchild == nil {
			b.replaceNode(parent, b.lchild)
		}

		replacement, replParent := b.lchild.findMax(b)
		b.data = replacement.data
		return replacement.delete(replacement.data, replParent)

	} else if data.less(b.data) {
		return b.lchild.delete(data, b)
	} else {
		return b.rchild.delete(data, b)
	}
}

func (tree *Tree) insert(data compare) error {
	if tree.root == nil {
		tree.root = &BiTNode{data: data}
		return nil
	}

	return tree.root.insert(data)
}

func (tree *Tree) search(data compare) (*BiTNode, bool) {
	if tree.root == nil {
		return nil, false
	}

	return tree.root.search(data)
}

func (tree *Tree) traverse(node *BiTNode) {
	if node == nil {
		return
	}
	tree.traverse(node.lchild)
	fmt.Println(node.data)
	tree.traverse(node.rchild)
}

func (tree *Tree) delete(data compare) error {
	if tree == nil {
		return errors.New("empty tree")
	}
	return tree.root.delete(data, tree.root)
}

func getRandomIntSlice(length int) []int {
	s := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		s = append(s, rand.Intn(100))
	}

	return s
}

func main() {
	//a := getRandomIntSlice(20)
	a := []int{15, 74, 14, 52, 37, 80, 92, 91, 61, 92, 44, 28, 72, 64, 56, 60, 84, 89, 50, 97}
	fmt.Println("unsorted : ", a)
	tree := new(Tree)
	for i := 0; i < 20; i++ {
		a1 := new(myInt)
		a1.number = a[i]
		err := tree.insert(a1)
		if err != nil {
			fmt.Println(err)
		}
	}

	tree.traverse(tree.root)

	b, ok := tree.search(&myInt{84})
	if ok {
		fmt.Println("find : ", b.data)
	} else {
		fmt.Println("can not find")
	}

	err := tree.delete(&myInt{84})
	if err != nil {
		fmt.Println(err)
	}

	tree.traverse(tree.root)

	b, ok = tree.search(&myInt{84})
	if ok {
		fmt.Println("find : ", b.data)
	} else {
		fmt.Println("can not find")
	}
}
