package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Value int64
	Left  *Node
	Right *Node
}

func (n *Node) Insert(v int64) error {
	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}
	switch {
	case v == n.Value:
		return nil
	case v < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: v}
			return nil
		}
		return n.Left.Insert(v)
	case v > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: v}
			return nil
		}
		return n.Right.Insert(v)
	}
	return nil
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func (t *Tree) Insert(v int64) error {
	if t.Root == nil {
		t.Root = &Node{Value: v}
		return nil
	}
	return t.Root.Insert(v)
}

func main() {
	data := []int64{5, 1, 2, 3, 4, 6, 7, 8, 9, 10}
	tree := &Tree{}
	for i := 0; i < len(data); i++ {
		err := tree.Insert(data[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Hello World!")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, " , ") })
}
