package main

import "fmt"

func main() {
	i := 0

	for i <= 10 {
		fmt.Println(i)
		if i == 10 {
			break
		}
		i++
	}
}