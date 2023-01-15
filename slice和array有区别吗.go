package main

import "fmt"

func main() {

	var a1 [3]int // 不初始化默认0值   bool类型不初始化默认false  字符串:""
	var a2 [4]int
	fmt.Printf("a1:%T  a2:%T\n", a1, a2)
	fmt.Println(a1[0])
  // 打印结果：
  // a1:[3]int  a2:[4]int
	// 0
  
	a3 := make([]int, 4, 6)
	fmt.Printf("a3=%T a3=%v, len(a3)=%d, cap(a3)=%d\n", a3, a3, len(a3), cap(a3))
  // 打印结果：
	// a3=[]int a3=[0 0 0 0], len(a3)=4, cap(a3)=6
}
