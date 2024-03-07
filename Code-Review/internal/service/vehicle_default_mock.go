package service

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// NewServiceVehicleDefaultMock is a function that returns a new instance of ServiceVehicleDefaultMock
func NewServiceVehicleDefaultMock() *Mock {
	return &Mock{}
}

type Mock struct {
	mock.Mock
	FuncFindByColorAndYear      func(color string, fabricationYear int) (v map[int]internal.Vehicle, err error)
	FuncFindByBrandAndYearRange func(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error)
	FuncAverageMaxSpeedByBrand  func(brand string) (a float64, err error)
	FuncAverageCapacityByBrand  func(brand string) (a int, err error)
	FuncSearchByWeightRange     func(startWeight int, endWeight int) (v map[int]internal.Vehicle, err error)
}

// FindByColorAndYear is a method that returns a map of vehicles that match the color and fabrication year
func (m *Mock) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(color, fabricationYear)
	if m.FuncFindByColorAndYear != nil {
		return m.FuncFindByColorAndYear(color, fabricationYear)
	}
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// FindByBrandAndYearRange is a method that returns a map of vehicles that match the brand and a range of fabrication years
func (m *Mock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(brand, startYear, endYear)
	if m.FuncFindByBrandAndYearRange != nil {
		return m.FuncFindByBrandAndYearRange(brand, startYear, endYear)
	}
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

// AverageMaxSpeedByBrand is a method that returns the average speed of the vehicles by brand
func (m *Mock) AverageMaxSpeedByBrand(brand string) (a float64, err error) {
	args := m.Called(brand)
	if m.FuncAverageMaxSpeedByBrand != nil {
		return m.FuncAverageMaxSpeedByBrand(brand)
	}
	return args.Get(0).(float64), args.Error(1)
}

// AverageCapacityByBrand is a method that returns the average capacity of the vehicles by brand
func (m *Mock) AverageCapacityByBrand(brand string) (a int, err error) {
	args := m.Called(brand)
	if m.FuncAverageCapacityByBrand != nil {
		return m.FuncAverageCapacityByBrand(brand)
	}
	return args.Get(0).(int), args.Error(1)
}

// FindByWeightRange is a method that returns a map of vehicles that match the weight range
func (m *Mock) SearchByWeightRange(query internal.SearchQuery, ok bool) (v map[int]internal.Vehicle, err error) {
	args := m.Called(query, ok)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}
