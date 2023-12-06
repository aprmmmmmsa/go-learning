package main

import "fmt"


func mul(a, b int) int {
	return a * b
}
func sub(a, b int) int {
	return a - b
}
func sum(a, b int) int {
	return a + b
}
func dev(a, b int) int {
	return a / b
}
func main() {
	fmt.Printf("Git using")
	fmt.Println(sum(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(mul(10, 2))
	fmt.Println(dev(10, 2))
}
