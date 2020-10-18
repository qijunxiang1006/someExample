package main

import (
	`fmt`
	`strings`
)

var migong="#....-...S.-.####-.E...-....."
var move=map[int][]int{
	0:{1,0},
	1:{-1,0},
	2:{0,1},
	3:{0,-1},
}
type StackItem struct {
	x,y int
	move []int
}
type Stack struct {
	data []*StackItem
}
func (s *Stack)push(item *StackItem){
	if s.data==nil{
		s.data=make([]*StackItem,0,1024)
	}
	s.data=append(s.data,item)
}
func (s *Stack)pop()*StackItem{
	if s.data==nil{
		return nil
	}
	item:=s.data[len(s.data)-1]
	s.data=s.data[:len(s.data)-1]
	return item
}
func findWayOut(m [][]byte,x,y,w,h int)(int,int){
	stack:=new(Stack)
	for _,v:=range move{
		stack.push(&StackItem{
			x:    x,
			y:    y,
			move: v,
		})
	}
	m[x][y]='#'
	for {
		item:=stack.pop()
		if item==nil{
			return 0,0
		}else {
			
			if item.x+item.move[0]>=0&&
				item.y+item.move[1]>=0&&
				item.x+item.move[0]<w&&
				item.y+item.move[1]<h{
				if m[item.x+item.move[0]][item.y+item.move[1]]=='.'{
					for _,v:=range move{
						stack.push(&StackItem{
							x:    item.x+item.move[0],
							y:    item.y+item.move[1],
							move: v,
						})
					}
					m[item.x+item.move[0]][item.y+item.move[1]]='#'
				}else if m[item.x+item.move[0]][item.y+item.move[1]]=='E'{
					return item.x+item.move[0],item.y+item.move[1]
				}
				
			}
		}
	}


}

func main() {
	str := strings.Split(migong, "-")
	h := len(str)
	w := len(str[0])
	grid := make([][]byte, h)
	var x, y int
	for i, v := range str {
		for j, k := range v {
			if grid[i] == nil {
				grid[i] = make([]byte, w)
			}
			grid[i][j] = byte(k)
			if k == 'S' {
				x = i
				y = j
			}
		}
	}
	for i, v := range str {
		for j, _ := range v {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println(" ")
	}
	fmt.Println(findWayOut(grid,x,y,w,h))
}
