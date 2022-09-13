package main

import "fmt"


func main() {

	var map1 map[int]string
	var map2 = make(map[string]int)

	map3 := map[float64]int{
		0.5:0,
		0.3:1,
		0.4:3,
	}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
}