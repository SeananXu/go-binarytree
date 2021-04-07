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

// LLRBNode is the node of left-leaning red–black tree.
type LLRBNode struct {
	color Color
	value interface{}
	left  *LLRBNode
	right *LLRBNode
}

// isRed returns whether the color of current node is red.
func (n *LLRBNode) isRed() bool {
	if n == nil {
		return false
	}
	return n.color.IsRed()
}

// Add adds values as a new node.
func (n *LLRBNode) Add(v interface{}, f CompareFunc) (Node, Action) {
	o, a := n.add(v, f)
	o.color = Black
	return o, a
}

func (n *LLRBNode) add(v interface{}, f CompareFunc) (*LLRBNode, Action) {
	if n == nil {
		return &LLRBNode{color: Red, value: v}, AddedAction
	}
	var a Action
	if c := f(n.value, v); c > 0 {
		n.left, a = n.left.add(v, f)
	} else if c < 0 {
		n.right, a = n.right.add(v, f)
	} else {
		n.value = v
		a = UpdatedAction
	}
	return n.balance(), a
}

func (n *LLRBNode) balance() *LLRBNode {
	if n == nil {
		return n
	}
	if n.right.isRed() && !n.left.isRed() {
		n = n.rotateLeft()
	}
	if n.left.isRed() && n.left.left.isRed() {
		n = n.rotateRight()
	}
	if n.left.isRed() && n.right.isRed() {
		n.flipColors()
	}
	return n
}

func (n *LLRBNode) rotateRight() *LLRBNode {
	out := n.left
	n.left = out.right
	out.right = n

	out.color = n.color
	n.color = Red
	return out
}

func (n *LLRBNode) rotateLeft() *LLRBNode {
	o := n.right
	n.right = o.left
	o.left = n

	o.color = n.color
	n.color = Red
	return o
}

func (n *LLRBNode) flipColors() {
	n.color = !n.color
	n.left.color = !n.left.color
	n.right.color = !n.right.color
}

func (n *LLRBNode) Remove(v interface{}, f CompareFunc) (Node, Action) {
	panic("implement me")
}

// Min returns the min value in all node.
func (n *LLRBNode) Min() interface{} {
	if n.left != nil {
		return n.left.Max()
	}
	return n.value
}

// Max returns the max value in all node.
func (n *LLRBNode) Max() interface{} {
	if n.right != nil {
		return n.right.Max()
	}
	return n.value
}

// Left returns the left node of current node.
func (n *LLRBNode) Left() Node {
	return n.left.self()
}

// Right returns the right node of current node.
func (n *LLRBNode) Right() Node {
	return n.right.self()
}

// self returns the current node.
func (n *LLRBNode) self() Node {
	// Avoids n is nil, but got node is not nil.
	if n == nil {
		return nil
	}
	return n
}

// Value returns the value of current node.
func (n *LLRBNode) Value() interface{} {
	if n == nil {
		return nil
	}
	return n.value
}

// Search returns the value in the node.
func (n *LLRBNode) Search(v interface{}, f CompareFunc) (interface{}, bool) {
	if n == nil {
		return nil, false
	}
	if c := f(n.value, v); c > 0 {
		return n.left.Search(v, f)
	} else if c < 0 {
		return n.right.Search(v, f)
	} else {
		return n.value, true
	}
}

// String returns node as string.
func (n *LLRBNode) String() string {
	return fmt.Sprintf("value: %v, color: %s", n.value, n.color)
}
