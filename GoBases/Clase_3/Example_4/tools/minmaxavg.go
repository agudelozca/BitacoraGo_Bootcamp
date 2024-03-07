package tools

import (
	"errors"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

// Función para realizar cálculos (mínimo, promedio o máximo)
func Operation(opType string) (func(...int) float64, error) {
	switch opType {
	case minimum:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			minValue := nums[0]
			for _, num := range nums[1:] {
				if num < minValue {
					minValue = num
				}
			}
			return float64(minValue)
		}, nil
	case average:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			sum := 0
			for _, num := range nums {
				sum += num
			}
			return float64(sum) / float64(len(nums))
		}, nil
	case maximum:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			maxValue := nums[0]
			for _, num := range nums[1:] {
				if num > maxValue {
					maxValue = num
				}
			}
			return float64(maxValue)
		}, nil
	default:
		return nil, errors.New("operación no definida")
	}
}
