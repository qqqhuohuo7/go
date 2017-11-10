package main

import (
	"fmt"
	//	"time"
)

var a = "G"

func n() { print(a) }

func m() {
	a := "O"
	print(a)
}
func fun1(abc string) {
	//	tempa, tempb := 'a', 'b'
	fmt.Println(abc)
}

func main() {
	//	n()
	//	m()
	//	n()
	//	fmt.Println(len("Welcome to the playground!"))

	//	fmt.Println("The time is", time.Now())
	//	fun1("hello")
	//	var i1 = 5
	//	var p1 *int
	//	fmt.Printf("an integer: %d, its location in memory: %p \n", i1, &i1)
	//	fmt.Printf("", p1)
	//	fmt.Println(test05())
	fmt.Println(test06(1, 1))
}
func test01() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	if i1 > 6 {
		fmt.Printf("hh")
	} else { // 无效的
		fmt.Printf("jj")
	}
}
func test02() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}
func test03() {
	for i := 0; i < 100; i++ {
		fmt.Println("Value of i is now:", i)
	}
}
func test04() {
	var num int = 10
	var numx2, numx3 int

	numx2, numx3 = getX2AndX3(num)
	PrintValues(num, numx2, numx3)
	getX2(&num)
	PrintValues(num, numx2, numx3)
}
func PrintValues(num int, numx2 int, numx3 int) {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2(input *int) (x2 int) {
	*input = 2 * (*input)
	// return x2, x3
	return
}
func test05() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

var count = 0

func test06(p1, p2 int) int {
	if count > 8 {
		return p2
	}
	count++
	return test06(p2, p1+p2)

}
