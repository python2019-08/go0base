package test

import "fmt"

type sType1 struct {
	a int
	b string
}

func Test_struct() {
	fmt.Println("Test_struct()...start")
	defer fmt.Println("Test_struct()...end")
	// pointer of struct
	var stP *sType1 = &sType1{1, "1"}
	fmt.Println("stP *sType1=", stP)

	// slice of struct
	var stSlice []sType1 = []sType1{
		{1, "1"},
		{2, "2"},
	}
	fmt.Println("stSlice []sType1=", stSlice)

	// slice of struct pointer
	var stpSlice []*sType1 = []*sType1{
		{1, "1"},
		{2, "2"},
	}
	fmt.Println("stpSlice []*sType1", stpSlice)

	// array of struct
	var stArr [4]sType1 = [4]sType1{
		{3, "3"},
		{4, "4"},
	}
	fmt.Println("stArr [4]sType1=", stArr)

	// array of struct pointer
	var stpArr [4]*sType1 = [4]*sType1{
		{3, "3"},
		{4, "4"},
	}
	fmt.Println("stpArr [4]*sType1 =", stpArr)

	var ifVoid01 interface{} = sType1{8, "8void"}
	fmt.Println("ifVoid01=", ifVoid01)

	var ifVoid02 interface{} = &sType1{9, "9void*"}
	fmt.Println("ifVoid02=", ifVoid02)
}

func Test_struct1() {
	fmt.Println("Test_struct1()...start")
	defer fmt.Println("Test_struct1()...end")

	// 结构体指针
	var stP *sType1 = &sType1{1, "1"}
	fmt.Println("stP *sType1=", stP)

	// 结构体数组
	var stSlice [4]sType1 = [4]sType1{
		{1, "1"},
		{2, "2"},
	}
	fmt.Println("stSlice [4]sType1=", stSlice)

	// 结构体指针数组 (正确初始化)
	var stpSlice [4]*sType1 = [4]*sType1{
		&sType1{1, "1"},
		&sType1{2, "2"},
	}
	fmt.Println("stpSlice [4]*sType1=", stpSlice)
}

func Test_struct2() {
	fmt.Println("Test_struct2()...start")
	defer fmt.Println("Test_struct2()...end")

	// 结构体指针
	var stP *sType1 = &sType1{1, "1"}
	fmt.Println("stP *sType1=", stP)

	// 结构体数组
	var stSlice [2]sType1 = [...]sType1{
		{1, "1"},
		{2, "2"},
	}
	fmt.Println("stSlice [4]sType1=", stSlice)

	// 结构体指针数组 (修正方式1: 使用...自动推断长度)
	var stpSlice1 [2]*sType1 = [...]*sType1{
		&sType1{1, "1"},
		&sType1{2, "2"},
	}
	fmt.Println("stpSlice1 [4]*sType1=", stpSlice1)

	// 结构体指针数组 (修正方式2: 使用短变量声明)
	stpSlice2 := [...]*sType1{
		&sType1{1, "1"},
		&sType1{2, "2"},
	}
	fmt.Println("stpSlice2 [4]*sType1=", stpSlice2)

	// 结构体指针数组 (修正方式3: 使用切片而不是数组)
	stpSlice3 := []*sType1{
		&sType1{1, "1"},
		&sType1{2, "2"},
	}
	fmt.Println("stpSlice3 []*sType1=", stpSlice3)

}
