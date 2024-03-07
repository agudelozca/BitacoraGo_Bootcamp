package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

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

func TestRepositoryVehicle_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		rp := repository.NewRepositoryReadVehicleMap(VehicleMap)
		// act
		vehicles, err := rp.FindAll()
		// assert
		require.NoError(t, err)
		require.Len(t, vehicles, 1)
	})
}

func TestRepositoryVehicle_FindByColorAndYear(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		rp := repository.NewRepositoryReadVehicleMap(VehicleMap)
		// act
		rp.FindAll()
		vehicles, err := rp.FindByColorAndYear("red", 2010)
		// assert
		require.NoError(t, err)
		require.Len(t, vehicles, 1)
	})
}

func TestRepositoryVehicle_FindByBrandAndYearRange(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		rp := repository.NewRepositoryReadVehicleMap(VehicleMap)
		// act
		vehicles, err := rp.FindByBrandAndYearRange("Ford", 2010, 2015)
		// assert
		require.NoError(t, err)
		require.Len(t, vehicles, 1)
	})

	t.Run("error", func(t *testing.T) {
		// arrange
		rp := repository.NewRepositoryReadVehicleMap(VehicleMap)
		// act
		vehicles, _ := rp.FindByBrandAndYearRange("Ford", 2015, 2010)
		// assert
		require.Len(t, vehicles, 0)
	})
}

func TestRepositoryVehicle_FindByBrand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		rp := repository.NewRepositoryReadVehicleMap(VehicleMap)
		// act
		vehicles, err := rp.FindByBrand("Ford")
		// assert
		require.NoError(t, err)
		require.Len(t, vehicles, 1)
	})
}

func TestRepositoryVehicle_FindByWeightRange(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		rp := repository.NewRepositoryReadVehicleMap(VehicleMap)
		// act
		vehicles, err := rp.FindByWeightRange(1000, 2000)
		// assert
		require.NoError(t, err)
		require.Len(t, vehicles, 1)
	})

	t.Run("error", func(t *testing.T) {
		// arrange
		rp := repository.NewRepositoryReadVehicleMap(VehicleMap)
		// act
		vehicles, _ := rp.FindByWeightRange(2000, 1000)
		// assert
		require.Len(t, vehicles, 0)
	})
}
