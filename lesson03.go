package main

import "fmt"

func main(){
	a := [3]string{"sato", "suzuki", "kimura"}
	a[1] = "tanaka"

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])

	str := ""
	fmt.Println(str)

	b :=[...]string{"nakamura", "osanai", "hayasaka", "gotou"}

	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])
	fmt.Println(b[3])
	
}