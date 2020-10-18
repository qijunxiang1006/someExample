package main

import (
	`fmt`
)

func main()  {
	// s:="aabcbabcaabcaababc"
	    substr:="abcaababc"
	fmt.Println(GetNext(substr))
}
func GetNext(str string)[]int{
	j:=0
	k:=-1
	l:= len(str)
	next:=make([]int, len(str))
	next[0]=-1
	for j<l-1{
		if k==-1||str[j]==str[k]{
			k++
			j++
			next[j]=k
		}else{
			k=next[k]
		}
	}
	return next
}
func StrIndex_KMP(str,substr string)int{
	start:=0
	for ;start< len(str)-len(substr);{
		for j:=0;j< len(substr);j++{
			if str[start]!=substr[j]{
				start=start+j
				break
			}
		}
	}
	return start
	
}