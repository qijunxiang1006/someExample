package main

import (
	"fmt"
	"math/rand"
)

func main(){
	l:=10
	a:=make([]int,l)
	for i:=0;i<l;i++{
		a[i]=rand.Intn(l)
	}
	a=[]int{1 ,7 ,7 ,9 ,1 ,8 ,5 ,0 ,6 ,0}
	HeapSort(a)
	fmt.Println(a)
	//t:=time.Now().UnixNano()
	//insertSort(a)
	//fmt.Println("insert sort",time.Now().UnixNano()-t)
	//
	//a=make([]int,l)
	//for i:=0;i<l;i++{
	//	a[i]=rand.Intn(l)
	//}
	//t=time.Now().UnixNano()
	//bubble(a)
	//fmt.Println("bubble sort",time.Now().UnixNano()-t)
	//
	//a=make([]int,l)
	//for i:=0;i<l;i++{
	//	a[i]=rand.Intn(l)
	//}
	//t=time.Now().UnixNano()
	//quick(a,0, len(a)-1)
	//fmt.Println("quick sort",time.Now().UnixNano()-t)

}
func minHeap(root int, end int, c []int)  {
	for {
		var child = 2*root + 1
		//判断是否存在child节点
		if child > end {
			break
		}
		//判断右child是否存在，如果存在则和另外一个同级节点进行比较
		if child+1 <= end && c[child] > c[child+1] {
			child += 1
		}
		if c[root] > c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}
	}
}
func HeapSort(c []int)  {
	var n = len(c)-1
	for root := n / 2; root >= 0; root-- {
		minHeap(root, n, c)
	}
	fmt.Println(c)
	for end := n; end >=0; end-- {
		if c[0]<c[end]{
			c[0], c[end] = c[end], c[0]
			minHeap(0, end-1, c)
		}
	}
}
func insertSort(a []int) {
	for i := 1; i < len(a); i++ {
		temp := a[i]
		j:=i - 1
		for ; j >= 0 && a[j] > temp; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = temp
	}
}
func bubble(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				b := a[i]
				a[i] = a[j]
				a[j] = b
			}
		}
	}
}

func quick(a []int, l, r int) {
	if l < r {
		x := a[l]
		i := l
		j := r
		for i < j {
			for i < j && a[j] >= x {
				j--
			}
			if i < j {
				a[i] = a[j]
			}
			for i < j && a[i] <= x {
				i++
			}
			if i < j {
				a[j] = a[i]
			}
		}
		a[i] = x
		quick(a, l, i-1)
		quick(a, i+1, r)
	}

}