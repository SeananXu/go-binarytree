English | [简体中文](./README-zh_CN.md)

## Binary Tree
a binary tree is a tree data structure in which each node has at most two children, which are referred to as the left child and the right child. 
here implements two types of tree: [BST(Binary Search Tree)](https://en.wikipedia.org/wiki/Binary_search_tree) and [AVLT(Adelson-Velsky and Landis Tree)](https://en.wikipedia.org/wiki/AVL_tree).

## Install
Using `go get` command to get the latest version
```bash
go get github.com/SeananXu/go-binarytree
```
Import it with:
```go
import "github.com/SeananXu/go-binarytree"
```
and use `binarytree` as the package name inside the code.

## Example
#### Initialization
```go
// Binary Search Tree
t := binarytree.NewBSTree(binarytree.IntCompareFunc)

// Adelson-Velsky and Landis Tree
t := binarytree.NewAVLTree(binarytree.IntCompareFunc)
```

#### Basic Operations
```go
// adds value
t.Add(1)

// removes value
t.Remove(1)

// min value
min := t.Min()

// max value
max := t.Max()

// get string of tree
str := t.String()

...
```
more BST cases click [here](./examples/bst/main.go), AVLT cases click [here](./examples/avlt/main.go)
## License

The MIT License (MIT) - see [LICENSE](LICENSE) for more details