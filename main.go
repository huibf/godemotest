package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"godemotest/test"
	"godemotest/test/arrtest"
	"godemotest/test/timetest"
	vartest2 "godemotest/test/vartest"
	"godemotest/utils"
)

func init() {
	fmt.Println("main的init函数：")
}

func main() {

	test.BaseTest()
	test.LogRecoverTest()
	test.LogSettest()

	test.Debug("./debug.log", "测试的显示内容"+time.Now().String())

	filelogger, _ := test.Logfile()
	filelogger = test.LogFileSet(filelogger)
	filelogger.Output(2, "hello  logfile")

	fileloggerD, _ := test.LogFileD()
	fileloggerD = test.LogFileSet(fileloggerD)

	fileloggerD.Println("hello  logfileD")
	fileloggerD.Output(2, "logfiled:输出测试文字")

	log.Print("log print")

	test.Forbase()

	utils.End()

	utils.Count()

	vartest2.VarTest()

	timetest.Timetest()

	timetest.SleepTimeTest()

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
