package main

import "fmt"

func main() {

	var a1 [3]int                       // 不初始化默认0值   bool类型不初始化默认false  字符串:""
	fmt.Printf("a1=%T a1=%v\n", a1, a1) //a1=[3]int a1=[0 0 0]
	fmt.Println(a1[0])                  //0

	a2 := make([]int, 3, 4)
	fmt.Printf("a2=%T a2=%v, len(a2)=%d, cap(a2)=%d\n", a2, a2, len(a2), cap(a2)) //a2=[]int a2=[0 0 0], len(a2)=3, cap(a2)=4

	a3 := [...]int{1, 3, 5, 7, 9}
	s1 := a3[0:3]
	fmt.Printf("s1=%T s1=%v\n", s1, s1) //s1=[]int s1=[1 3 5]

}
