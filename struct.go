package main

import "fmt"

type struct1 struct {
	dev1 string
	dev2 string
}




type person struct {
	name string
	id int
	department department
}

type department struct {
	depid int
	depname string
}

func main() {
	var structVar1 struct1
	structVar2 := struct1{"srini", "kandu"}
	structVar1 = struct1{"sr", "k"}
	
	fmt.Println(structVar1)
	fmt.Println(structVar2)


	c := person{
		id:1,
		name:"sri",
		department: department{
			depid:2,
			depname:"scan",
		},
	}

	fmt.Println(c);



}