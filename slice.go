package main

import "fmt"


func updateSlice(slc []int, pos int, val int) []int {
	if(pos >= len(slc)){
		fmt.Println("slice index out of bound")
	} else {
		slc[pos] = val
	}
	return slc
}

func main() {
	
	slice := []string{"srini", "vas"}
	fmt.Println(slice);
	slice = append(slice,"dev");

	fmt.Println("this is slice 1 : %v", slice)

	slc := []int{1,2,3,4,5}
	fmt.Println(slc);
	slc = updateSlice(slc, 2,200)
	fmt.Println(slc);

}