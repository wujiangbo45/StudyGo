package main

import (
	"fmt"
)

func main() {
	test()
}

func sw(a int, b int) (int, int) {
	var tmp int
	tmp = a
	a = b
	b = tmp
	return a, b
}

func test() {
	var c = 2
	var d = 3
	//var e *int
	//e = &d
	//f := c
	fmt.Println(sw(c, d))
	fmt.Printf("a=%d,b=%d\n", c, d)

	swPoint(&c, &d)

	fmt.Printf("a=%d,b=%d\n", c, d)
	//if c == 3 {
	//	print(c,f)
	//}
}

func swPoint(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
