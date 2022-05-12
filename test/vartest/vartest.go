package vartest

import (
	"fmt"
	"strconv"
)

func VarTest() {
	fmt.Println("开始变量测试：")

	c, ss := 45, 2
	m := c / ss
	n := float64(c) / float64(ss)
	fmt.Println("数字：", m, n)

	oStr := "Hello World"
	fmt.Println("字符串：", oStr)

	fmt.Println("开始数字转字符串：")
	j := 45.34
	strtmp := strconv.Itoa(int(j))
	fmt.Printf("转换前: %T : %d \n 转换后: %T : %s \n", c, c, strtmp, strtmp)
	fmt.Printf("经常使用的输出方式：%T : %#v : %v \n", strtmp, strtmp, strtmp)

	fmt.Printf(`end \n`)
	fmt.Println()

}
