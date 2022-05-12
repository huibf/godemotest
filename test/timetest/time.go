package timetest

import (
	"fmt"
	"math/rand"
	"time"
)

func Timetest() {
	fmt.Println("随机数测试开始：")
	tm1 := time.Now()
	rand.Seed(tm1.UnixNano())
	rd1 := rand.Intn(45) //[0,45)
	fmt.Println("获得随机数：", rd1)
}
