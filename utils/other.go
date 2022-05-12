package utils

import "fmt"

func Poo(tt []int) {

	fmt.Println("poo方法(切片)开始：")
	fmt.Printf("交换前：%T : %#v : %v \t", tt, tt, tt)
	tt[0], tt[len(tt)-1] = tt[len(tt)-1], tt[0]
	//fmt.Println(tt[0])

}

func Boo(tt [6]int) {
	fmt.Println("boo方法(数组)开始：")
	fmt.Printf("交换前：%T : %#v : %v \t", tt, tt, tt)
	tt[0], tt[len(tt)-1] = tt[len(tt)-1], tt[0]
}

func TwoDimensionArray() {

	fmt.Println("多维数组开始：")
	/* 数组 - 5 行 2 列*/
	var a = [][]int{{0, 0}, {1, 2}, {2}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

func ChanlTest() {
	fmt.Println("通道测试开始：")

	ch := make(chan int, 5)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("并发通道测试开始")
	cf := make(chan int)

	go func() {

		cf <- 4
		cf <- 6
		cf <- 8
		close(cf)
	}()

	fmt.Println("循环输出通道：")
	for v := range cf {
		fmt.Println(v)
	}

}
