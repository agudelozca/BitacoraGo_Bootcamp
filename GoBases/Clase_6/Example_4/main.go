package main

import (
	"fmt"
)

func main() {
	salary := 1000
	minimumSalary := 150000

	if salary < minimumSalary {
		err := fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
		fmt.Println(err)
	}
}
