package graph

import (
	"fmt"
	"testing"
)

func TestNode_PreOrder(t *testing.T) {
	fmt.Println("--------------PreOrder----------------")
	tree.PreOrder()
	fmt.Println()
}

func TestNode_MidOrder(t *testing.T) {
	fmt.Println("--------------MidOrder----------------")
	tree.MidOrder()
	fmt.Println()
}

func TestNode_PostOrder(t *testing.T) {
	fmt.Println("--------------PostOrder----------------")
	tree.PostOrder()
	fmt.Println()
}