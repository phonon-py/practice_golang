package main

import (
	"fmt"
	"reflect"
)

// スラッシュ二つでコメントアウト
func main()  {
	num := 1
	num01 := 2
	num_01 := 3
	NUM := 4

	num02 := 123
	var num_02 int = 1234567890
	num03 := 1.23
	var num_03 float64 =1.23456789

	str := "######"

	fmt.Println(num)
	fmt.Println(num01)
	fmt.Println(num_01)
	fmt.Println(NUM)
	fmt.Println(str)
	fmt.Println(num02)
	fmt.Println(reflect.TypeOf(num02))
	fmt.Println(num_02)
	fmt.Println(reflect.TypeOf(num_02))
	fmt.Println(num03)
	fmt.Println(reflect.TypeOf(num03))
	fmt.Println(num_03)
	fmt.Println(reflect.TypeOf(num_03))
}
