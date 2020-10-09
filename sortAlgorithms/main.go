package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	l:=10
	a:=make([]int,l)
	for i:=0;i<l;i++{
		a[i]=rand.Intn(l)
	}
	t:=time.Now().UnixNano()
	insertSort(a)
	fmt.Println("insert sort",time.Now().UnixNano()-t)

	a=make([]int,l)
	for i:=0;i<l;i++{
		a[i]=rand.Intn(l)
	}
	t=time.Now().UnixNano()
	bubble(a)
	fmt.Println("bubble sort",time.Now().UnixNano()-t)

	a=make([]int,l)
	for i:=0;i<l;i++{
		a[i]=rand.Intn(l)
	}
	t=time.Now().UnixNano()
	quick(a,0, len(a)-1)
	fmt.Println("quick sort",time.Now().UnixNano()-t)

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