package test

import (
	"fmt"
	"log"
)

func Forbase() {
	for i := 0; i < 5; i++ {
		log.Printf("for语句的基本演示：%d", i)

	}
	log.Printf("for语句外是无法打印i的；不在一个作用域")
}

func ForChan() {
	chl := make(chan int)
	go func() {
		chl <- 34
		chl <- 345
		chl <- 3456
		close(chl)
	}()
	for v := range chl {
		fmt.Println("通道：", v)
	}

}
