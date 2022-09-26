package main

import "fmt"

/*
注意第一个括号的位置
*/
func swap(a, b int) (int, int) {
	return b, a
}
func max(a, b int) int {
	result := a
	if b > a {
		result = b
	}
	return result
}
func main() {
	var a, b, c, d int = 1, 2, 3, 4
	a, b = swap(a, b)
	c = max(c, d)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
