package main

import (
	"fmt"
	"jobworker/model"
)

func main() {
	go model.Person1()
	go model.Person2()
	// go model.Person3()
	test := make(chan string)

	go func() {
		fmt.Println("aaaaaaaaaa")

		fmt.Println("rountine 取值")
		str := <-test
		fmt.Println("取到值了：", str)
	}()
	fmt.Println("bbbbbbbbbbbb")

	test <- "hello"
	fmt.Println("cccccccccc")

	for {
	}
}
