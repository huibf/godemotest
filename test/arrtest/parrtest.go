package arrtest

import "fmt"

const MAX int = 3

func init() {
	fmt.Println("arrtest包的init函数：parrtest.go")
}

func Parrtest() {
	fmt.Println("指针数组测试：")

	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d；%x\n", i, *ptr[i], ptr[i])
	}
}
