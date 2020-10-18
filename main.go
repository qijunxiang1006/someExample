package main

import (
	`fmt`
)

var move map[int][]int =map[int][]int{0: []int{1,0},
						1:[]int{0,1},
						2:[]int{-1,0},
						3:[]int{0,-1}, }
				

func moveNext(grid [][]byte,x,y,r,c int)bool{
	if x<r && x>=0 && y<c && y>=0 && grid[x][y]!=2{
		if grid[x][y]==1{
			grid[x][y]=2
			for _,v:=range move{
				moveNext(grid,x+v[0],y+v[1],r,c)
			}
			return true
		}else{
			return false
		}
	}
	return false
}

func solve( grid [][]byte ) int {
	num:=0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if moveNext(grid,i,j,len(grid),len(grid[i])){
				num++
			}
		}
	}
	return num
}
type ListNode struct{
	Val int
	Next *ListNode
}
func ReverseList( pHead *ListNode ) *ListNode {
	if pHead==nil{
		return nil
	}
	if pHead.Next==nil{
		return pHead
	}
	curr:=pHead
	var pre *ListNode
	next:=curr.Next
	curr.Next=pre
	pre=curr
	curr=next
	for curr.Next!=nil{
		next=curr.Next
		curr.Next=pre
		pre=curr
		curr=next
	}
	return curr
}
func isValid( str string ) bool {
	var couple=map[byte]byte{
		'(':')',
		'[':']',
		'{':'}',
	}
	stack:=make([]byte,0,1024)
	push:=func(b byte){
		if stack!=nil{
			stack=append(stack,b)
		}
	}
	popup:=func()byte{
		if len(stack)>=1{
			top:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			return top
		}
		return 0
	}
	for i:=0;i< len(str);i++{
		_,exist:=couple[str[i]];if exist{
			push(str[i])
		}else {
			if couple[popup()]!=str[i]{
				return false
			}
		}
	}
	return true
}
func main()  {
	
	// data:=[][]byte{{1,1,0,0,0},{0,1,0,1,1},{0,0,0,1,1},{0,0,0,0,0},{0,0,1,1,1}}
	// var pre *ListNode=&ListNode{
	// 	0,
	// 	nil,
	// }
	// node:=pre
	// for i:=1;i<6;i+=1{
	// 	curr:=&ListNode{
	// 		i,
	// 		nil,
	// 	}
	// 	pre.Next=curr
	// 	pre=curr
	// }
	// fmt.Println(ReverseList(node).Val)
	fmt.Println(isValid("([]{)}"))
	// s:=newMyStack()
	// s.push('a')
	// s.push('b')
	// s.push('c')
	// fmt.Println(s.popup())
	// fmt.Println(s.data)
}