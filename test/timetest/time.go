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

func SleepTimeTest() {
	sleepTime := 1*time.Hour + 10*time.Minute + 5*time.Second
	//time.Sleep(sleepTime)
	fmt.Printf("时间延时的设置；如睡眠1小时10分5秒：%#v", sleepTime)
}
