package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				fmt.Println("panic1 recovered")
			}
			if r == "panic2" {
				fmt.Println("panic2 recovered")
			}
		}
	}()

	panic1()
	// panic2()
}

func panic1() {
	panic("panic1")
}

// func panic2() {
// 	panic("panic2")
// }
