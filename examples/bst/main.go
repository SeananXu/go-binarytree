package main

import (
	"fmt"

	"binarytree"
)

func main() {
	fmt.Printf("-------------------------------\n------------int\n-------------------------------\n")
	tr := binarytree.NewBSTree(binarytree.IntCompareFunc)
	tr.Add(8)
	tr.Add(5)
	tr.Add(10)
	tr.Add(7)
	tr.Add(6)
	fmt.Println(tr.String())
	fmt.Println("-------------------------------")
	fmt.Printf("min: %v\n", tr.Min())
	fmt.Println("-------------------------------")
	fmt.Printf("max: %v\n", tr.Max())
	fmt.Println("-------------------------------")
	result, ok := tr.Search(7)
	fmt.Printf("search 7: %d, %v\n", result, ok)
	fmt.Println("-------------------------------")
	tr.Remove(6)
	fmt.Println(tr.String())

	fmt.Printf("-------------------------------\n------------kv\n-------------------------------\n")
	tr = binarytree.NewBSTree(binarytree.KVCompareFunc)
	tr.Add(&binarytree.KV{Key: "a", Value: "a"})
	tr.Add(&binarytree.KV{Key: "b", Value: "b"})
	tr.Add(&binarytree.KV{Key: "c", Value: "c"})
	tr.Add(&binarytree.KV{Key: "d", Value: "d"})
	tr.Add(&binarytree.KV{Key: "e", Value: "e"})
	fmt.Println(tr.String())
	fmt.Println("-------------------------------")
	fmt.Printf("min: %v\n", tr.Min())
	fmt.Println("-------------------------------")
	fmt.Printf("max: %v\n", tr.Max())
	fmt.Println("-------------------------------")
	result, ok = tr.Search(&binarytree.KV{Key: "c"})
	fmt.Printf("search c: %v, %v\n", result, ok)
	fmt.Println("-------------------------------")
	tr.Remove(&binarytree.KV{Key: "e"})
	fmt.Println(tr.String())

}
