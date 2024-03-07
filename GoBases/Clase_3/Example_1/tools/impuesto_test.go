package tools_test

import (
	"intro-unit-testing/Example_1/tools"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcularImpuestoMenorA50k(t *testing.T) {
	// Arrange
	//Act
	salario := 40000.0
	impuesto := tools.SalaryImpuesto(salario)
	//Assert
	require.Equal(t, 0.0, impuesto)
}
func TestCalcularImpuestoMayorA50k(t *testing.T) {
	salario := 60000.0
	impuesto := tools.SalaryImpuesto(salario)
	if impuesto != salario*0.17 {
		t.Errorf("Para salario por encima de $50,000, el impuesto calculado no es correcto. Se esperaba %v, pero se obtuvo %v", salario*0.17, impuesto)
	}
}
func TestCalcularImpuestoMayorA150k(t *testing.T) {
	salario := 160000.0
	impuesto := tools.SalaryImpuesto(salario)
	if impuesto != salario*0.27 {
		t.Errorf("Para salario por encima de $150,000, el impuesto calculado no es correcto. Se esperaba %v, pero se obtuvo %v", salario*0.27, impuesto)
	}
}
