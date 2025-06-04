package main

import (
	"fmt"
	go_pool "go0base/go-pool"
	"go0base/gorm"
	"go0base/test"
)

func main() {
	fmt.Println("----main()....start")
	defer fmt.Println("----main()....end")
	go_pool.Go_Pool()

	s1 := make([]byte, 20)
	s1[1] = 99
	s1[2] = 100
	fmt.Println(s1)
	test.Test_struct()

	// gorm.CreateRecord()
	gorm.Query()

	fmt.Println("Hello, World!")
}
