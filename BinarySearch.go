package main

// BinarySearchTree 二叉查找树
type BinarySearchTree struct {
	Root *BinarySearchTreeNode // 树根节点
}

// BinarySearchTreeNode 二叉查找树节点
type BinarySearchTreeNode struct {
	Value int64                 // 值
	Times int64                 // 值出现的次数
	Left  *BinarySearchTreeNode // 左子树
	Right *BinarySearchTreeNode // 右字树
}

// NewBinarySearchTree 初始化一个二叉查找树
func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

func (tree *BinarySearchTree) Add(value int64) {
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}
	tree.Root.Add(value)
	return
}

func (node *BinarySearchTreeNode) Add(value int64) {

}
