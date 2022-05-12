package test

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func BaseTest() {
	fmt.Println("test包的BaseTest函数")

	fmt.Println("math包的一个简单示例：", math.Exp2(3))

	log.Println("第一个log包的测试")
}

func RecoverTest() {
	log.Println("first 异常处理：")
	if err := recover(); err != nil {
		fmt.Println(err)
	}

}

func LogRecoverTest() {
	defer RecoverTest()

	errInfo := errors.New("自定义的错误信息；发生了一个异常")

	log.Panic(errInfo)
}
