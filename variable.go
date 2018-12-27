package main

import "fmt"

func main() {
	var a1 string
	a1 = "a1 string"
	var a2 int = 10
	a3 := 20

	fmt.Println( a1, a2, a3)


	// 第二类
	var b1, b2, b3 int = 1, 2, 3
	c1, c2, c3 := 1, "string c2", 1.1

	fmt.Println(b1, b2, b3, c1, c2, c3)

	// 第三类
	var (
		d1 int = 2
		d2 float32
	)

	d1 = 1
	d2 = 1.1

	fmt.Println(d1, d2)

	// 指针
	fmt.Println(&d1)

	// 两个变量交换值，前提是类型要一样
	b1, b2 = b2, b1

	fmt.Println(b1, b2)

	// const
	const PI = 3.1415926535897
	e1 := 3 * PI

	fmt.Println(e1)

	// iota
	const (
		f1 = iota
		f2
		f3
		f4 = "f4 string"
		f5 = 10
		f6
	)

	fmt.Println(f1, f2, f3, f4, f5, f6)

	//
	var ptr *int
	ptr = &d1

	fmt.Println(ptr, *ptr)
}
