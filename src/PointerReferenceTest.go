package main

import "fmt"

type InterfaceTest interface {
	printData()
}

type StructTest struct {
	data string
}

func (st StructTest) printData() {
	fmt.Println("StructTest" + st.data)
}

type StructTest2 struct {
	data string
}

func (st *StructTest2) printData() {
	fmt.Println("StructTest2" + st.data)
}

//对于一个类来说，其值类型的方法会自动生成相应的指针类型的方法，而指针类型的方法不会自动生成相应的值类型的方法
func main() {
	st := StructTest{data: "seafooler"}
	var it1 InterfaceTest = st
	it1.printData()

	var it2 InterfaceTest = &st
	it2.printData()

	//st2 := StructTest2{data: "seafooler"}
	//var it3 InterfaceTest = st2  // 编译出错
	//it3.printData()

	var it4 InterfaceTest = &st
	it4.printData()
}
