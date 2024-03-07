package tools_test

import (
	"intro-unit-testing/Example_3/tools"
	"testing"
)

func TestCalcularSalarioCategoriaA(t *testing.T) {
	minutosTrabajados := 160
	categoria := "A"
	resultado := tools.CalcuSalary(minutosTrabajados, categoria)
	expected := float64(160 * 3000 / 60) // 160 minutos * $3000/hora * 40 horas / 60 minutos
	if resultado != expected {
		t.Errorf("El salario calculado para la categoría A no es correcto. Se esperaba %v, pero se obtuvo %v", expected, resultado)
	}
}

// Función de prueba para calcular el salario de la categoría "B"
func TestCalcularSalarioCategoriaB(t *testing.T) {
	minutosTrabajados := 160
	categoria := "B"
	resultado := tools.CalcuSalary(minutosTrabajados, categoria)
	expected := float64(160 * 1500 / 60) // 160 minutos * $1500/hora * 40 horas / 60 minutos
	if resultado != expected {
		t.Errorf("El salario calculado para la categoría B no es correcto. Se esperaba %v, pero se obtuvo %v", expected, resultado)
	}
}

// Función de prueba para calcular el salario de la categoría "C"
func TestCalcularSalarioCategoriaC(t *testing.T) {
	minutosTrabajados := 160
	categoria := "C"
	resultado := tools.CalcuSalary(minutosTrabajados, categoria)
	expected := float64(160 * 1000 * 40 / 60) // 160 minutos * $1000/hora * 40 horas / 60 minutos
	if resultado != expected {
		t.Errorf("El salario calculado para la categoría C no es correcto. Se esperaba %v, pero se obtuvo %v", expected, resultado)
	}
}
