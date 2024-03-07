package main

import (
	"fmt"
)

type SalaryError struct {
	msg string
}

func (e *SalaryError) Error() string {
	return e.msg
}

func calculateSalary(hoursWorked, hourlyRate int) (float64, error) {
	if hoursWorked < 0 || hoursWorked < 80 {
		return 0, &SalaryError{"Error: the worker cannot have worked less than 80 hours per month"}
	}

	salary := float64(hoursWorked * hourlyRate)

	if salary >= 150000 {
		salary *= 0.9 // Aplicar descuento del 10% si el salario es igual o superior a $150,000
	}

	return salary, nil
}

func main() {
	hoursWorked := 100
	hourlyRate := 1500

	salary, err := calculateSalary(hoursWorked, hourlyRate)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Printf("Salary: $%.2f\n", salary)
	}
}
