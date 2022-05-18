package test

import "log"

func Forbase() {
	for i := 0; i < 5; i++ {
		log.Printf("for语句的基本演示：%d", i)

	}
	log.Printf("for语句外是无法打印i的；不在一个作用域")
}
