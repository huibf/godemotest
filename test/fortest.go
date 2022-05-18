package test

import (
	"fmt"
	"log"
)

func ForMaptest() {

	mmap := map[string]int{
		"张三": 34,
		"李四": 89,
		"王五": 52,
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}

	for v := range mmap {
		log.Println(v)
	}
}
