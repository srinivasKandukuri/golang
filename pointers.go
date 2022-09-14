package main

import "fmt"


func zeroVal(val int){
	val = 0
}

func zerOptr(iptr *int){
	*iptr =0
}

func main(){
	i :=1

	fmt.Println("initial val : ", i)

	zeroVal(i)

	fmt.Println("zeroVal", i)

	zerOptr(&i)
	
	fmt.Println("zerOptr", i)
	fmt.Println("zerOptr", &i)
}