package graph

import "fmt"

//二叉树的一个结点
type Node struct {
	Self int
	Left *Node
	Right *Node
}

var (
	tree = Node{
		'A',
		&Node{
			'B',
			&Node{
				'D',
				nil,
				&Node{'E', nil, nil},
			},
			&Node{
				'F',
				&Node{'G', nil, nil},
				nil,
			},
		},
		&Node{
			'C',
			nil,nil,
		},
	}
)

//前序遍历
func (tree *Node) PreOrder(){
	if tree != nil {
		fmt.Printf("%c\t", tree.Self)
		tree.Left.PreOrder()
		tree.Right.PreOrder()
	}
}

func (tree *Node) MidOrder() {
	if tree != nil {
		tree.Left.MidOrder()
		fmt.Printf("%c\t", tree.Self)
		tree.Right.MidOrder()
	}
}


func (tree *Node) PostOrder() {
	if tree != nil {
		tree.Left.PostOrder()
		tree.Right.PostOrder()
		fmt.Printf("%c\t", tree.Self)
	}
}