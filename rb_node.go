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

// RBNode is the node of red–black tree
type RBNode struct {
	color  Color
	value  interface{}
	parent *RBNode
	left   *RBNode
	right  *RBNode
}

func (n *RBNode) Add(v interface{}, f CompareFunc) (Node, Action) {
	o, a := n.add(v, f)
	o.balance()
	return n, a
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

func (n *RBNode) add(v interface{}, f CompareFunc) (*RBNode, Action) {
	if n == nil {
		return &RBNode{value: v, color: Red}, AddedAction
	}
	var a Action
	if c := f(n.value, v); c > 0 {
		n.left, a = n.left.add(v, f)
		return n.left, a
	} else if c < 0 {
		n.right, a = n.right.add(v, f)
		return n.right, a
	} else {
		n.value = v
		// update value, so here is false
		a = UpdatedAction
		return n, a
	}
}

func (n *RBNode) balance() {
	parent, uncle, grandparent := n.pug()
	// CASE1: the current node is root.
	if parent == nil {
		n.color = Black
		return
	}
	// CASE2: The color of parent node is black.
	if n.parent.color == Black {
		return
	}
	// CASE3: The color of parent and uncle are red.
	if uncle != nil && uncle.color == Red {
		parent.color = Black
		uncle.color = Black
		grandparent.color = Red
		grandparent.balance()
		return
	}
	// CASE4: The color of parent is red, but uncle is black
	if parent.isLeft() {
		if !n.isLeft() {
			parent.rotateLeft()
			parent, uncle, grandparent = parent.pug()
		}
		parent.color = Black
		grandparent.color = Red
		grandparent.rotateRight()
	} else {
		if n.isLeft() {
			parent.rotateRight()
			parent, uncle, grandparent = parent.pug()
		}
		parent.color = Black
		grandparent.color = Red
		grandparent.rotateLeft()
	}
}

// rotateRight rotates right node.
func (n *RBNode) rotateRight() *RBNode {
	o := n.left
	n.left = o.right
	if o.right != nil {
		o.right.parent = n
	}
	p := n.parent
	if p != nil {
		if p.isLeft() {
			p.left = o
		} else {
			p.right = o
		}
	}
	o.parent = p
	n.parent = o
	o.right = n
	return o
}

// rotateRight rotates left node.
func (n *RBNode) rotateLeft() *RBNode {
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
	}
	o.parent = p
	n.parent = o
	o.left = n
	return o
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
	panic("implement me")
}

func (n *RBNode) Right() Node {
	panic("implement me")
}

func (n *RBNode) Value() interface{} {
	panic("implement me")
}

func (n *RBNode) Search(v interface{}, f CompareFunc) (interface{}, bool) {
	panic("implement me")
}
