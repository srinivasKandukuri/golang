package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empid  int
	salary int
	pf     int
}

type Contract struct {
	empid  int
	salary int
}

type FreeLancer struct {
	empid        int
	salary       int
	workinghours int
}

func (p Permanent) CalculateSalary() int {
	return p.salary + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.salary
}

func (f FreeLancer) CalculateSalary() int {

	return f.salary * f.workinghours
}

func totalExpensives(s []SalaryCalculator) {

	totalExpensives := 0

	for _, c := range s {
		totalExpensives += c.CalculateSalary()
	}

	fmt.Printf("Total Expense Per Month $%d", totalExpensives)

}

func main() {

	p := Permanent{empid: 1, salary: 20, pf: 10}
	c := Contract{empid: 2, salary: 30}

	f := FreeLancer{empid: 3, salary: 10, workinghours: 40}

	emp := []SalaryCalculator{p, c, f}

	totalExpensives(emp)

}
