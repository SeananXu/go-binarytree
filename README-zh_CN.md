[English](./README.md) | 简体中文

## 二叉树
二叉树是一个每个节点最多有两个节点的树状数据结构, 分别为左节点和右节点. 这里实现两种二叉树:
[二叉搜索树(Binary Search Tree)](https://en.wikipedia.org/wiki/Binary_search_tree) 和 [平衡二叉树(Adelson-Velsky and Landis Tree)](https://en.wikipedia.org/wiki/AVL_tree).

## 安装
使用 `go get` 指令获取最新代码
```bash
go get github.com/SeananXu/go-binarytree
```
`improt` 依赖:
```go
import "github.com/SeananXu/go-binarytree"
```
使用 `binarytree` 作为包名使用代码

## 例子
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
更多二叉搜索树案例点击[这里](./examples/bst/main.go), 更多平衡二叉树案例点击[这里](./examples/avlt/main.go)
## License

The MIT License (MIT) - see [LICENSE](LICENSE) for more details