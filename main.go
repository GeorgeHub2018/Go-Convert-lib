package main

import "fmt"

func main() {
	fmt.Println(ToInt64("3"))
	fmt.Println(ToInt64(3.1))
	fmt.Println(ToString(5))
	fmt.Println(ToString(5.2))
	fmt.Println(ToFloat32(3))
	fmt.Println(ToFloat32(3.1))
	fmt.Println(ToFloat32("3"))
}
