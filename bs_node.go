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

// BSNode is node of binary tree.
type BSNode struct {
	value interface{}
	left  *BSNode
	right *BSNode
}

// Add adds values as a new node.
func (n *BSNode) Add(v interface{}, f CompareFunc) (Node, Action) {
	return n.add(v, f)
}

func (n *BSNode) add(v interface{}, f CompareFunc) (*BSNode, Action) {
	if n == nil {
		return &BSNode{value: v}, AddedAction
	}
	var a Action
	if c := f(n.value, v); c > 0 {
		n.left, a = n.left.add(v, f)
	} else if c < 0 {
		n.right, a = n.right.add(v, f)
	} else {
		n.value = v
		// update value, so here is false
		a = UpdatedAction
	}
	return n, a
}

// Remove removes the node that value is v.
func (n *BSNode) Remove(v interface{}, f CompareFunc) (Node, Action) {
	return n.remove(v, f)
}

func (n *BSNode) remove(v interface{}, f CompareFunc) (*BSNode, Action) {
	if n == nil {
		return nil, UnchangedAction
	}
	var a Action
	if c := f(n.value, v); c > 0 {
		n.left, a = n.left.remove(v, f)
	} else if c < 0 {
		n.right, a = n.right.remove(v, f)
	} else {
		if n.right == nil && n.left == nil {
			return nil, UnchangedAction
		} else if n.right == nil {
			return n.left, RemovedAction
		} else if n.left == nil {
			return n.right, RemovedAction
		} else {
			min := n.right.Min()
			n.value = min
			n.right, _ = n.right.remove(min, f)
			a = RemovedAction
		}
	}
	return n, a
}

// Min returns the min value in all node.
func (n *BSNode) Min() interface{} {
	if n.left != nil {
		return n.left.Min()
	}
	return n.value
}

// Max returns the max value in all node.
func (n *BSNode) Max() interface{} {
	if n.right != nil {
		return n.right.Max()
	}
	return n.value
}

// Left returns the left node of current node.
func (n *BSNode) Left() Node {
	return n.left.self()
}

// Right returns the right node of current node.
func (n *BSNode) Right() Node {
	return n.right.self()
}

// self returns the current node.
func (n *BSNode) self() Node {
	// Avoids n is nil, but got node is not nil.
	if n == nil {
		return nil
	}
	return n
}

// Value returns the value of current node.
func (n *BSNode) Value() interface{} {
	if n == nil {
		return nil
	}
	return n.value
}

// Search returns the value in the node
func (n *BSNode) Search(v interface{}, f CompareFunc) (interface{}, bool) {
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
func (n *BSNode) String() string {
	return fmt.Sprintf("value: %v", n.value)
}
