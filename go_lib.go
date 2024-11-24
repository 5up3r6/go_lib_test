package go_lib_test

import "fmt"

func init() {
	fmt.Println("Hello, World! init")
}

// SayHello 打印 Hello, World!
func SayHello() {
    fmt.Println("Hello, World!")
}

// SayGreeting 接受一个名字并打印问候语
func SayGreeting(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

