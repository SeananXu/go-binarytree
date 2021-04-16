/*
MIT License
Copyright (c) 2021 Seanan Xu
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package binarytree implements a binary tree.
// A binary tree is a tree data structure in which each node has at most two children,
// which are referred to as the left child and the right child.
package binarytree

import (
	"fmt"
	"strings"
)

// Tree is a binary tree.
type Tree struct {
	root        Node
	size        int64
	compareFunc CompareFunc
}

// Add adds v.
func (t *Tree) Add(v interface{}) {
	n, a := t.root.Add(v, t.compareFunc)
	if a.Added() {
		t.size++
	}
	t.root = n
}

// Remove removes v.
func (t *Tree) Remove(v interface{}) {
	n, a := t.root.Remove(v, t.compareFunc)
	if a.Removed() {
		t.size--
	}
	t.root = n
}

// Min returns the min value of tree.
func (t *Tree) Min() interface{} {
	return t.root.Min()
}

// Max returns the max value of tree.
func (t *Tree) Max() interface{} {
	return t.root.Max()
}

// Search returns the value in the tree
func (t *Tree) Search(v interface{}) (interface{}, bool) {
	return t.root.Search(v, t.compareFunc)
}

// Size returns the size of the tree.
func (t *Tree) Size() int64 {
	return t.size
}

// interval returns string of tree.
func (t *Tree) String() string {
	builder := &strings.Builder{}
	stringify(t.root, builder, "--[", "      ", 0)
	return builder.String()
}

// StringFor returns string of tree by specifying trunk and interval.
func (t *Tree) StringFor(trunk string, interval string) string {
	builder := &strings.Builder{}
	stringify(t.root, builder, trunk, interval, 0)
	return builder.String()
}

func stringify(n Node, builder *strings.Builder, trunk string, interval string, level int) {
	if n != nil {
		prefix := ""
		for i := 0; i < level; i++ {
			prefix += interval
		}
		prefix += trunk
		level++
		stringify(n.Right(), builder, trunk, interval, level)
		fmt.Printf(prefix+"%v\n", n)
		stringify(n.Left(), builder, trunk, interval, level)
	}
}

// NewBSTree returns a new Binary Search Tree.
func NewBSTree(f CompareFunc) *Tree {
	return &Tree{root: (*BSNode)(nil), compareFunc: f}
}

// NewAVLTree returns a new Adelson-Velsky and Landis Tree.
func NewAVLTree(f CompareFunc) *Tree {
	return &Tree{root: (*AVLNode)(nil), compareFunc: f}
}

func NewRBTree(f CompareFunc) *Tree {
	return &Tree{root: (*RBNode)(nil), compareFunc: f}
}

func NewLLRBTree(f CompareFunc) *Tree {
	return &Tree{root: (*LLRBNode)(nil), compareFunc: f}
}
