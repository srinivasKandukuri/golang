package main

import "fmt"

func main() {
	nums := []int{1,2,3,4}

	sum :=0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum", sum)

	for i, num := range nums {
		sum += num +i

		fmt.Println("sum, num, i %v %v %v", sum, num, i)
	}

	
	strs := map[string]string{"a":"apple", "b":"banana"}

	for k,v := range strs {
		fmt.Println("%s -> %s\n", k,v)
	}


	for k := range strs {

		fmt.Println("key", k)
	}

	for i,c := range "go"{
		fmt.Println(i,c)
	}
}