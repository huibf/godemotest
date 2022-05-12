package arrtest

import "fmt"

func SliceTest() {
	fmt.Println("切片测试开始：")
	arr := []int{3, 6, 8, 24, 75}
	s := arr[:]
	for k := range s {
		fmt.Println(s[k])
	}

	fmt.Println("make切片的默认值：")
	s1 := make([]int, 3)
	for k := range s1 {
		fmt.Println(s1[k])
	}

	slice1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	s2 := slice1[3:7:10]
	fmt.Printf("切片输出的值：%#v \n", s2)
	fmt.Printf("长度：%d; 容量：%d \n", len(s2), cap(s2))
	fmt.Println()
}
