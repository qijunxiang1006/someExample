package main

import (
	`fmt`
)

type BinarySearchTree struct {
	Key int
	LChild *BinarySearchTree
	RChild *BinarySearchTree
}
func FrontOrderBinarySearchTree(root *BinarySearchTree){
	if root!=nil{
		FrontOrderBinarySearchTree(root.RChild)
		
		fmt.Println(root.Key)
		
		FrontOrderBinarySearchTree(root.LChild)
		
	}
}

func NewBinarySearchTree(key []int)(root *BinarySearchTree) {
	for _,v:=range key{
		if root==nil{
			root=new(BinarySearchTree)
			root.Key=v
			continue
		}
		InsertBinarySearchTree(root,v)
	}
	return root
}
func InsertBinarySearchTree(root *BinarySearchTree,key int)  {
	if root==nil{
		r:=new(BinarySearchTree)
		r.Key=key
		return
	}
	current:=root
	c,p:=FindBinarySearchTree(current,key);if c==nil{
		node:=new(BinarySearchTree)
		node.Key=key
		if node.Key<p.Key{
			p.LChild=node
		}else if node.Key>p.Key{
			p.RChild=node
		}
		root=p
	}
}
func FindBinarySearchTree(root *BinarySearchTree,key int)(*BinarySearchTree,*BinarySearchTree){
	if root==nil{
		return nil,nil
	}
	current:=root
	var p *BinarySearchTree
	for current!=nil{
		if current.Key==key{
			return current,p
		}
		p=current
		if current.Key>key{
			current=current.LChild
		}
		if p.Key<key{
			current=current.RChild
		}
	}
	return nil,p
}
type AVLTree struct {

}
func main(){
	data:=[]int{9,4,6,3,1,2,5,8,7}
	root:=NewBinarySearchTree(data)
	FrontOrderBinarySearchTree(root)
}
