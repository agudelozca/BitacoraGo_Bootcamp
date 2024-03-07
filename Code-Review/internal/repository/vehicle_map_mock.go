package repository

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

// NewRepositoryMock is a mock of RepositoryVehicleMap
func NewRepositoryMock() *Mock {
	return &Mock{}
}

type Mock struct {
	mock.Mock
	FuncFindAll                 func() (v map[int]internal.Vehicle, err error)
	FuncFindByColorAndYear      func(color string, fabricationYear int) (v map[int]internal.Vehicle, err error)
	FuncFindByBrandAndYearRange func(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error)
	FuncFindByBrand             func(brand string) (v map[int]internal.Vehicle, err error)
	FuncFindByWeightRange       func(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error)
}

func (m *Mock) FindAll() (v map[int]internal.Vehicle, err error) {
	args := m.Called()
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *Mock) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *Mock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *Mock) FindByBrand(brand string) (v map[int]internal.Vehicle, err error) {
	args := m.Called(brand)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *Mock) FindByWeightRange(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error) {
	args := m.Called(fromWeight, toWeight)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}
