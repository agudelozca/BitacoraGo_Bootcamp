package tools

func SalaryImpuesto(salario float64) float64 {
	if salario > 150000 {
		return salario * 0.27
	} else if salario > 50000 {
		return salario * 0.17
	}
	return 0
}
