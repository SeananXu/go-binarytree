package main

import (
	"fmt"

	"binarytree"
)

func main() {
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
	tr.Remove(6)
	fmt.Println(tr.String())
}
