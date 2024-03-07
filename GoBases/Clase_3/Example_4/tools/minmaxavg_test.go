package tools_test

import (
	"intro-unit-testing/Example_4/tools"
	"testing"
)

// Función de prueba para calcular el mínimo de calificaciones
func TestCalcularMinimo(t *testing.T) {
	calificaciones := []int{80, 90, 85, 95, 88}
	minFunc, _ := tools.Operation("minimum")
	minValue := minFunc(calificaciones...)
	expected := 80.0
	if minValue != expected {
		t.Errorf("El mínimo de calificaciones no es correcto. Se esperaba %f, pero se obtuvo  %.2f", expected, minValue)
	}
}

// Función de prueba para calcular el máximo de calificaciones
func TestCalcularMaximo(t *testing.T) {
	calificaciones := []int{80, 90, 85, 95, 88}
	maxFunc, _ := tools.Operation("maximum")
	maxValue := maxFunc(calificaciones...)
	expected := 95.0
	if maxValue != expected {
		t.Errorf("El máximo de calificaciones no es correcto. Se esperaba %v, pero se obtuvo %.2f", expected, maxValue)
	}
}

// Función de prueba para calcular el promedio de calificaciones
func TestCalcularPromedio(t *testing.T) {
	calificaciones := []int{80, 90, 85, 95, 88}
	avgFunc, _ := tools.Operation("average")
	resultado := avgFunc(calificaciones...)
	expected := 87.6
	if resultado != expected {
		t.Errorf("El promedio de calificaciones no es correcto. Se esperaba %v, pero se obtuvo %.2f", expected, resultado)
	}
}
