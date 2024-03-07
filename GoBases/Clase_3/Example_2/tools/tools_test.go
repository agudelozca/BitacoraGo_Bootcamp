package tools_test

import (
	"intro-unit-testing/Example_2/tools"
	"testing"
)

func TestCalcularPromedio(t *testing.T) {
	// Caso exitoso
	resultado := tools.CalcularProm(80, 90, 85, 95, 88)
	expected := 87.6
	if resultado != expected {
		t.Errorf("El promedio calculado no es correcto. Se esperaba %v, pero se obtuvo %v", expected, resultado)
	}

	// Caso sin notas
	resultado = tools.CalcularProm()
	expected = 0
	if resultado != expected {
		t.Errorf("El promedio calculado para ningún alumno debería ser 0, pero se obtuvo %v", resultado)
	}
}
