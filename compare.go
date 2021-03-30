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

// CompareFunc compares between a and b.
// if a > b, returns 1
// if a == b, returns 0
// if a < b, returns -1
type CompareFunc func(a, b interface{}) int

// KV provides simple entry.
type KV struct {
	Key   string
	Value interface{}
}

// CompareTo compares between a and b.
func (a *KV) CompareTo(b *KV) int {
	if a == nil {
		return -1
	}
	if b == nil {
		return 1
	}
	return strings.Compare(a.Key, b.Key)
}

func (a *KV) String() string {
	return fmt.Sprintf("k: %s, v: %v", a.Key, a.Value)
}

var (
	// IntCompareFunc compares between a and b, a and b is int type.
	IntCompareFunc CompareFunc = func(a, b interface{}) int {
		switch {
		case a.(int) > b.(int):
			return 1
		case a.(int) == b.(int):
			return 0
		default:
			return -1
		}
	}

	// StringCompareFunc compares between a and b, a and b is string type.
	StringCompareFunc CompareFunc = func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	}

	// KVCompareFunc compares between a and b, a and b is KV type.
	KVCompareFunc CompareFunc = func(a, b interface{}) int {
		return a.(*KV).CompareTo(b.(*KV))
	}
)
