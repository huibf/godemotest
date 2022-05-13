package utils

import (
	"fmt"
	"log"
	"time"
)

func Count() {
	fmt.Println("utils包Count函数。。。")
}

func End() {
	defer func() {
		log.Println("调用了没有呀；当然没有了")
	}()

	time.Sleep(5 * time.Second)
	log.Fatalln("人为的；结束了") //Fatalln会调用 os.Exit(1)执行退出
}
