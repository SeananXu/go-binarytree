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

import "fmt"

// RBNode is the node of red–black tree
type RBNode struct {
	color  Color
	value  interface{}
	parent *RBNode
	left   *RBNode
	right  *RBNode
}

func (n *RBNode) Add(v interface{}, f CompareFunc) (Node, Action) {
	if n == nil {
		return &RBNode{color: Black, value: v}, AddedAction
	}
	o, a := n.add(v, f)
	return o.balance(n), a
}

func (n *RBNode) isLeft() bool {
	return n == n.parent.left
}

func (n *RBNode) isRight() bool {
	return n == n.parent.right
}

func (n *RBNode) pug() (parent *RBNode, uncle *RBNode, grandparent *RBNode) {
	if n.parent != nil {
		parent = n.parent
		if grandparent = parent.parent; grandparent != nil {
			if parent.isLeft() {
				uncle = grandparent.right
			} else {
				uncle = grandparent.left
			}
		}
	}
	return
}

func (n *RBNode) add(v interface{}, f CompareFunc) (o *RBNode, a Action) {
	for node := n; ; {
		if c := f(node.value, v); c > 0 {
			if node.left == nil {
				node.left = &RBNode{color: Red, value: v, parent: node}
				return node.left, AddedAction
			}
			node = node.left
		} else if c < 0 {
			if node.right == nil {
				node.right = &RBNode{color: Red, value: v, parent: node}
				return node.right, AddedAction
			}
			node = node.right
		} else {
			node.value = v
			return node, UpdatedAction
		}
	}
}

func (n *RBNode) balance(root *RBNode) *RBNode {
	parent, uncle, grandparent := n.pug()
	// CASE1: the current node is root.
	if parent == nil {
		n.color = Black
		return n
	}
	// CASE2: The color of parent node is black.
	if n.parent.color == Black {
		return nil
	}
	// CASE3: The color of parent and uncle are red.
	if uncle != nil && uncle.color == Red {
		parent.color = Black
		uncle.color = Black
		grandparent.color = Red
		return grandparent.balance(root)
	}
	// CASE4: The color of parent is red, but uncle is black
	if parent.isLeft() {
		if !n.isLeft() {
			root = parent.rotateLeft(root)
			parent, uncle, grandparent = parent.pug()
		}
		parent.color = Black
		grandparent.color = Red
		return grandparent.rotateRight(root)
	} else {
		if n.isLeft() {
			root = parent.rotateRight(root)
			parent, uncle, grandparent = parent.pug()
		}
		parent.color = Black
		grandparent.color = Red
		return grandparent.rotateLeft(root)
	}
}

// rotateRight rotates right node.
func (n *RBNode) rotateRight(root *RBNode) *RBNode {
	o := n.left
	if o.right != nil {
		o.right.parent = n
	}
	n.left = o.right
	p := n.parent
	if p != nil {
		if p.isLeft() {
			p.left = o
		} else {
			p.right = o
		}
	} else {
		root = o
	}
	o.parent = p
	n.parent = o
	o.right = n
	return root
}

// rotateRight rotates left node.
func (n *RBNode) rotateLeft(root *RBNode) *RBNode {
	o := n.right
	n.right = o.left
	if o.left != nil {
		o.left.parent = n
	}
	p := n.parent
	if p != nil {
		if p.isLeft() {
			p.left = o
		} else {
			p.right = o
		}
	} else {
		root = o
	}
	o.parent = p
	n.parent = o
	o.left = n
	return root
}

func (n *RBNode) String() string {
	return fmt.Sprintf("value: %v, color: %s, parent: %v", n.value, n.color, n.parent.Value())
}

func (n *RBNode) Remove(v interface{}, f CompareFunc) (Node, Action) {
	panic("implement me")
}

func (n *RBNode) Min() interface{} {
	panic("implement me")
}

func (n *RBNode) Max() interface{} {
	panic("implement me")
}

func (n *RBNode) Left() Node {
	return n.left.self()
}

func (n *RBNode) Right() Node {
	return n.right.self()
}

// self returns the current node.
func (n *RBNode) self() Node {
	// Avoids n is nil, but got node is not nil.
	if n == nil {
		return nil
	}
	return n
}

func (n *RBNode) Value() interface{} {
	if n == nil {
		return nil
	}
	return n.value
}

func (n *RBNode) Search(v interface{}, f CompareFunc) (interface{}, bool) {
	panic("implement me")
}
