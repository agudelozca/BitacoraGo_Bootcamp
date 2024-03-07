package tools

func CalcularProm(notas ...float64) float64 {
	var sumaTotal float64
	var prom float64
	cantidad := len(notas)
	for _, i := range notas {
		if i < 0 {
			return 0
		}
		sumaTotal += i
	}
	prom = sumaTotal / float64(cantidad)

	return prom
}
