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

// Node is node of tree, it has two child node that are left and right.
type Node interface {
	// Add adds values as a new node.
	// f func compares the two node.
	Add(v interface{}, f CompareFunc) Node
	// Remove removes the node that value is v.
	Remove(v interface{}, f CompareFunc) Node
	// Min returns the min value in all node.
	Min() interface{}
	// Max returns the max value in all node.
	Max() interface{}
	// Left returns the left node of current node.
	Left() Node
	// Right returns the right node of current node.
	Right() Node
	// Value returns the value of current node.
	Value() interface{}
}

// BSNode is node of binary tree.
type BSNode struct {
	value interface{}
	left  *BSNode
	right *BSNode
}

// Add adds values as a new node.
func (n *BSNode) Add(v interface{}, f CompareFunc) Node {
	return n.add(v, f)
}

func (n *BSNode) add(v interface{}, f CompareFunc) *BSNode {
	if n == nil {
		return &BSNode{value: v}
	}
	if c := f(n.value, v); c > 0 {
		n.left = n.left.add(v, f)
	} else if c < 0 {
		n.right = n.right.add(v, f)
	} else {
		n.value = v
	}
	return n
}

// Remove removes the node that value is v.
func (n *BSNode) Remove(v interface{}, f CompareFunc) Node {
	return n.remove(v, f)
}

func (n *BSNode) remove(v interface{}, f CompareFunc) *BSNode {
	if n == nil {
		return nil
	}
	if c := f(n.value, v); c > 0 {
		n.left = n.left.remove(v, f)
	} else if c < 0 {
		n.right = n.right.remove(v, f)
	} else {
		if n.right == nil && n.left == nil {
			return nil
		} else if n.right == nil {
			return n.left
		} else if n.left == nil {
			return n.right
		} else {
			min := n.right.Min()
			n.value = min
			n.right = n.right.remove(min, f)
		}
	}
	return n
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

type AVLNode struct {
	value interface{}
	left  *AVLNode
	right *AVLNode
	// xheight is a anonymous variable, instead of using height func.
	xheight int
}

// Add adds values as a new node.
func (n *AVLNode) Add(v interface{}, f CompareFunc) Node {
	return n.add(v, f)
}

func (n *AVLNode) add(v interface{}, f CompareFunc) *AVLNode {
	if n == nil {
		return &AVLNode{value: v, xheight: 1}
	}
	if c := f(n.value, v); c > 0 {
		n.left = n.left.add(v, f)
	} else if c < 0 {
		n.right = n.right.add(v, f)
	} else {
		n.value = v
	}
	return n.balance()
}

// Remove removes the node that value is v.
func (n *AVLNode) Remove(v interface{}, f CompareFunc) Node {
	return n.remove(v, f)
}

func (n *AVLNode) remove(v interface{}, f CompareFunc) *AVLNode {
	if n == nil {
		return nil
	}
	if c := f(n.value, v); c > 0 {
		n.left = n.left.remove(v, f)
	} else if c < 0 {
		n.right = n.right.remove(v, f)
	} else {
		if n.right == nil && n.left == nil {
			return nil
		} else if n.right == nil {
			n = n.left
		} else if n.left == nil {
			n = n.right
		} else {
			min := n.right.Min()
			n.value = min
			n.right = n.right.remove(min, f)
		}
	}
	return n.balance()
}

// Min returns the min value in all node.
func (n *AVLNode) Min() interface{} {
	if n.left != nil {
		return n.left.Min()
	}
	return n.value
}

// Max returns the max value in all node.
func (n *AVLNode) Max() interface{} {
	if n.right != nil {
		return n.right.Max()
	}
	return n.value
}

// Left returns the left node of current node.
func (n *AVLNode) Left() Node {
	return n.left.self()
}

// Right returns the right node of current node.
func (n *AVLNode) Right() Node {
	return n.right.self()
}

// self returns the current node.
func (n *AVLNode) self() Node {
	// Avoids n is nil, but got node is not nil.
	if n == nil {
		return nil
	}
	return n
}

// Value returns the value of current node.
func (n *AVLNode) Value() interface{} {
	if n == nil {
		return nil
	}
	return n.value
}

// balance checks and balance the node.
func (n *AVLNode) balance() *AVLNode {
	if n == nil {
		return nil
	}
	n.calculateHeight()
	factor := n.left.height() - n.right.height()
	switch factor {
	case -2:
		if n.right.left.height() > n.right.right.height() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	case 2:
		if n.left.right.height() > n.left.left.height() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	default:
		return n
	}
}

// rotateRight rotates right node.
func (n *AVLNode) rotateRight() *AVLNode {
	dest := n.left
	n.left = dest.right
	dest.right = n
	n.calculateHeight()
	dest.calculateHeight()
	return dest
}

// rotateRight rotates left node.
func (n *AVLNode) rotateLeft() *AVLNode {
	dest := n.right
	n.right = dest.left
	dest.left = n
	n.calculateHeight()
	dest.calculateHeight()
	return dest
}

// calculateHeight calculate the height of current node.
func (n *AVLNode) calculateHeight() {
	n.xheight = 1 + max(n.left.height(), n.right.height())
}

// max returns the max between a and b.
// if a > b, returns a,
// else returns b.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// height returns the height of current node.
func (n *AVLNode) height() int {
	if n == nil {
		return 0
	}
	return n.xheight
}
