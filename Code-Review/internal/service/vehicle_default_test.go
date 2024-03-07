package service_test

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

// VehicleMap is a map of vehicles
var VehicleMap = map[int]internal.Vehicle{
	1: {
		Id: 1,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           "Ford",
			Model:           "Fiesta",
			Registration:    "ABC-123",
			Color:           "red",
			FabricationYear: 2010,
			Capacity:        5,
			MaxSpeed:        180,
			FuelType:        "gasoline",
			Transmission:    "manual",
			Weight:          1000,
			Dimensions: internal.Dimensions{
				Height: 1.5,
				Length: 4,
				Width:  1.8,
			},
		},
	},
}

func TestServiceVehicleDefault_FindByColorAndYear(t *testing.T) {
	t.Run("success, vehicles found", func(t *testing.T) {
		// arrange
		rp := repository.NewRepositoryMock()
		rp.On("FindByColorAndYear", "red", 2010).Return(VehicleMap, nil)

		sv := service.NewServiceVehicleDefault(rp)
		// act
		vehicles, err := sv.FindByColorAndYear("red", 2010)
		// assert
		require.NoError(t, err)
		require.Len(t, vehicles, 1)
		rp.AssertExpectations(t)

	})

	t.Run("error - no vehicles", func(t *testing.T) {
		// arrange
		rp := repository.NewRepositoryMock()
		sv := service.NewServiceVehicleDefault(rp)
		rp.On("FindByColorAndYear", "red", 2010).Return(map[int]internal.Vehicle{}, nil)
		// act
		v, err := sv.FindByColorAndYear("red", 2010)
		// assert
		require.Error(t, err)
		require.Len(t, v, 0)
		require.EqualError(t, err, "service: no vehicles")

	})
}

func TestServiceVehicleDefault_FindByBrandAndYearRange(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		rp := repository.NewRepositoryMock()
		rp.On("FindByBrandAndYearRange", "Ford", 2010, 2015).Return(VehicleMap, nil)

		sv := service.NewServiceVehicleDefault(rp)
		// act
		vehicles, err := sv.FindByBrandAndYearRange("Ford", 2010, 2015)
		// assert
		require.NoError(t, err)
		require.Len(t, vehicles, 1)
		rp.AssertExpectations(t)
	})

	t.Run("error - no vehicles", func(t *testing.T) {
		rp := repository.NewRepositoryMock()
		rp.On("FindByBrandAndYearRange", "Ford", 2010, 2015).Return(map[int]internal.Vehicle{}, nil)

		sv := service.NewServiceVehicleDefault(rp)
		// act
		vehicles, err := sv.FindByBrandAndYearRange("Ford", 2010, 2015)
		// assert
		require.Error(t, err)
		require.Len(t, vehicles, 0)
		require.EqualError(t, err, "service: no vehicles")
	})
}

func TestServiceVehicleDefault_AverageMaxSpeedByBrand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rp := repository.NewRepositoryMock()
		rp.On("FindByBrand", "Ford").Return(VehicleMap, nil)

		sv := service.NewServiceVehicleDefault(rp)
		// act
		average, err := sv.AverageMaxSpeedByBrand("Ford")
		// assert
		require.NoError(t, err)
		require.Equal(t, 180.0, average)
		rp.AssertExpectations(t)

	})

	t.Run("error - no vehicles", func(t *testing.T) {
		rp := repository.NewRepositoryMock()
		rp.On("FindByBrand", "Ford").Return(map[int]internal.Vehicle{}, nil)

		sv := service.NewServiceVehicleDefault(rp)
		// act
		average, err := sv.AverageMaxSpeedByBrand("Ford")
		// assert
		require.Error(t, err)
		require.Equal(t, 0.0, average)
		require.EqualError(t, err, "service: no vehicles")
		rp.AssertExpectations(t)

	})
}

func TestServiceVehicleDefault_AverageCapacityByBrand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rp := repository.NewRepositoryMock()
		rp.On("FindByBrand", "Ford").Return(VehicleMap, nil)

		sv := service.NewServiceVehicleDefault(rp)
		// act
		average, err := sv.AverageCapacityByBrand("Ford")
		// assert
		require.NoError(t, err)
		require.Equal(t, 5, average)
		rp.AssertExpectations(t)

	})

	t.Run("error - no vehicles", func(t *testing.T) {
		rp := repository.NewRepositoryMock()
		rp.On("FindByBrand", "Ford").Return(map[int]internal.Vehicle{}, nil)

		sv := service.NewServiceVehicleDefault(rp)
		// act
		average, err := sv.AverageCapacityByBrand("Ford")
		// assert
		require.Error(t, err)
		require.Equal(t, 0, average)
		require.EqualError(t, err, "service: no vehicles")
		rp.AssertExpectations(t)

	})
}

func TestServiceVehicleDefault_SearchByWeightRange(t *testing.T) {
	t.Run("case - query !ok then find all", func(t *testing.T) {
		rp := repository.NewRepositoryMock()
		rp.On("FindAll").Return(VehicleMap, nil)

		sv := service.NewServiceVehicleDefault(rp)
		// act
		vehicles, err := sv.SearchByWeightRange(internal.SearchQuery{}, false)
		// assert
		require.NoError(t, err)
		require.Len(t, vehicles, 1)
		rp.AssertExpectations(t)

	})

	t.Run("case - query ok then find by weight range", func(t *testing.T) {
		rp := repository.NewRepositoryMock()
		rp.On("FindByWeightRange", 1000.0, 2000.0).Return(VehicleMap, nil)

		sv := service.NewServiceVehicleDefault(rp)
		// act
		vehicles, err := sv.SearchByWeightRange(internal.SearchQuery{
			FromWeight: 1000.0,
			ToWeight:   2000.0,
		}, true)
		// assert
		require.NoError(t, err)
		require.Len(t, vehicles, 1)
		rp.AssertExpectations(t)

	})

	t.Run("case - error - no vehicles", func(t *testing.T) {
		rp := repository.NewRepositoryMock()
		rp.On("FindByWeightRange", 1000.0, 2000.0).Return(map[int]internal.Vehicle{}, nil)

		sv := service.NewServiceVehicleDefault(rp)
		// act
		vehicles, err := sv.SearchByWeightRange(internal.SearchQuery{
			FromWeight: 1000.0,
			ToWeight:   2000.0,
		}, true)
		// assert
		require.Error(t, err)
		require.Len(t, vehicles, 0)
		require.EqualError(t, err, "service: no vehicles")
		rp.AssertExpectations(t)

	})

}
