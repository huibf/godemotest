package main

import (
	"fmt"
	"os"

	"hello/utils"

	"hello/test"
	"hello/test/arrtest"
	"hello/test/timetest"
	vartest2 "hello/test/vartest"
)

func init() {
	fmt.Println("main的init函数：")
}

func main() {

	os.Exit(1)

	utils.Count()

	test.BaseTest()

	vartest2.VarTest()

	timetest.Timetest()

	//系统参数(flag包也可以获取命令参数)
	fmt.Println("命令行参数数量：", len(os.Args))
	for _, arg := range os.Args {
		fmt.Println("值：", arg)
	}

	arrtest.ArrTest()

	arrtest.SliceTest()

	utils.TwoDimensionArray()
	fmt.Println()

	// 切片
	p := []int{2, 3, 5, 7, 11, 13}
	utils.Poo(p)
	fmt.Println("参数交换后的结果：", p) // [13 3 5 7 11 2]

	// 数组
	b := [...]int{2, 3, 5, 7, 11, 13}
	utils.Boo(b)
	//poo(b)  //参数类型不一致
	fmt.Println("参数交换后的结果：", b) // [2 3 5 7 11 13]

	arrtest.Parrtest()

	utils.ChanlTest()

}
