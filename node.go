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
	Add(v interface{}, f CompareFunc) (Node, Action)
	// Remove removes the node that value is v.
	Remove(v interface{}, f CompareFunc) (Node, Action)
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
	// Search returns the value in the node
	Search(v interface{}, f CompareFunc) (interface{}, bool)
}

// Action returns the type of node operation.
type Action int8

const (
	// UnchangedAction means the node operation nothing changes.
	UnchangedAction Action = iota
	// AddedAction means the node operation added a new node.
	AddedAction
	// UpdatedAction means the node operation updated a node.
	UpdatedAction
	// RemovedAction means the node operation removed a exist node.
	RemovedAction
)

// Added judges whether the action is AddedAction.
func (a Action) Added() bool {
	return a == AddedAction
}

// Updated judges whether the action is UpdatedAction.
func (a Action) Updated() bool {
	return a == UpdatedAction
}

// Removed judges whether the action is RemovedAction.
func (a Action) Removed() bool {
	return a == RemovedAction
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

// Color means the color of red-black tree.
type Color bool

const (
	// Red means the color of node is red.
	Red Color = true
	// Black means the color of node is black.
	Black Color = false
)

// IsRed predicates whether the color is red.
func (c Color) IsRed() bool {
	return c == Red
}

// String returns the name of color.
func (c Color) String() string {
	switch c {
	case Black:
		return "black"
	case Red:
		return "red"
	default:
		panic("no such color")
	}
}
