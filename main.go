package main

import "fmt"

func Hi(name string) string {
	return "Hi " + name
}

//
//func Hello(name string) string {
//	return Hi(name)
//}

func main() {
	answer := Hi("Duncan")
	fmt.Println(answer)
}
