package arrtest

import "fmt"

func init() {
	fmt.Println("arrtest包的init函数:arrtest.go")
}

func ArrTest() {
	fmt.Println("数组测试开始：")
	var balance = [5]float32{1000.0, 2.0, 3.4, 7, 50.0}
	fmt.Println("循环输出：")
	for i := 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}
}
