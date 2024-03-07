package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/service"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandlerVehicle_FindByColorAndYear(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.FindByColorAndYear()
		s.On("FindByColorAndYear", "red", 2010).Return(map[int]internal.Vehicle{
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
		}, nil)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/color/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("color", "red")
		ctx.URLParams.Add("year", "2010")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusOK, w.Code)
		expectBody := `{
			"message": "vehicles found",
			"data": {
				"1": {
					"Id": 1,
					"Brand": "Ford",
					"Model": "Fiesta",
					"Registration": "ABC-123",
					"Color": "red",
					"FabricationYear": 2010,
					"Capacity": 5,
					"MaxSpeed": 180,
					"FuelType": "gasoline",
					"Transmission": "manual",
					"Weight": 1000,
					"Height": 1.5,
					"Length": 4,
					"Width": 1.8
				}
			}
		}`

		require.JSONEq(t, expectBody, w.Body.String())
		s.AssertExpectations(t)
		s.AssertNumberOfCalls(t, "FindByColorAndYear", 1)
	})

	t.Run("case error, year invalid in request", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.FindByColorAndYear()
		s.On("FindByColorAndYear", "red", 2010).Return(map[int]internal.Vehicle{}, nil)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/color/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("color", "red")
		ctx.URLParams.Add("year", "201a")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusBadRequest, w.Code)
		expectBody := `{
			"status": "Bad Request",
			"message": "invalid year"
		}`

		require.JSONEq(t, expectBody, w.Body.String())

	})
}

func TestHandlerVehicle_FindByBrandAndYearRange(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.FindByBrandAndYearRange()
		s.On("FindByBrandAndYearRange", "Ford", 2010, 2015).Return(map[int]internal.Vehicle{
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
		}, nil)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/brand/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "Ford")
		ctx.URLParams.Add("start_year", "2010")
		ctx.URLParams.Add("end_year", "2015")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusOK, w.Code)
		expectBody := `{
			"message": "vehicles found",
			"data": {
				"1": {
					"Id": 1,
					"Brand": "Ford",
					"Model": "Fiesta",
					"Registration": "ABC-123",
					"Color": "red",
					"FabricationYear": 2010,
					"Capacity": 5,
					"MaxSpeed": 180,
					"FuelType": "gasoline",
					"Transmission": "manual",
					"Weight": 1000,
					"Height": 1.5,
					"Length": 4,
					"Width": 1.8
				}
			}
		}`
		require.JSONEq(t, expectBody, w.Body.String())
		s.AssertExpectations(t)
		s.AssertNumberOfCalls(t, "FindByBrandAndYearRange", 1)
	})
	t.Run("case error, start_year invalid in request", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.FindByBrandAndYearRange()
		s.On("FindByBrandAndYearRange", "Ford", 2010, 2015).Return(map[int]internal.Vehicle{}, nil)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/brand/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "Ford")
		ctx.URLParams.Add("start_year", "201a")
		ctx.URLParams.Add("end_year", "2015")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusBadRequest, w.Code)
		expectBody := `{
			"status": "Bad Request",
			"message": "invalid start_year"
		}`
		require.JSONEq(t, expectBody, w.Body.String())
	})

	t.Run("case error, end_year invalid in request", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.FindByBrandAndYearRange()
		s.On("FindByBrandAndYearRange", "Ford", 2010, 2015).Return(map[int]internal.Vehicle{}, nil)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/brand/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "Ford")
		ctx.URLParams.Add("start_year", "2010")
		ctx.URLParams.Add("end_year", "201a")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusBadRequest, w.Code)
		expectBody := `{
			"status": "Bad Request",
			"message": "invalid end_year"
		}`
		require.JSONEq(t, expectBody, w.Body.String())
	})

}

func Test_handler_AverageMaxSpeedByBrand(t *testing.T) {
	t.Run("case - success", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.AverageMaxSpeedByBrand()
		s.On("AverageMaxSpeedByBrand", "Ford").Return(180.0, nil)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/average_speed/brand/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "Ford")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusOK, w.Code)
		expectBody := `{
			"message": "average max speed found",
			"data": 180
		}`
		require.JSONEq(t, expectBody, w.Body.String())
		s.AssertExpectations(t)
		s.AssertNumberOfCalls(t, "AverageMaxSpeedByBrand", 1)
	})

	t.Run("case error, vehicle not found", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.AverageMaxSpeedByBrand()
		s.On("AverageMaxSpeedByBrand", "Ford").Return(0.0, internal.ErrServiceNoVehicles)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/average_speed/brand/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "Ford")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusNotFound, w.Code)
		expectBody := `{
			"status": "Not Found",
			"message": "vehicles not found"
		}`
		require.JSONEq(t, expectBody, w.Body.String())
		s.AssertExpectations(t)
		s.AssertNumberOfCalls(t, "AverageMaxSpeedByBrand", 1)
	})
}

func Test_handler_AverageCapacityByBrand(t *testing.T) {
	t.Run("case - success", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.AverageCapacityByBrand()
		s.On("AverageCapacityByBrand", "Ford").Return(5, nil)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/average_capacity/brand/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "Ford")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusOK, w.Code)
		expectBody := `{
			"message": "average capacity found",
			"data": 5
		}`
		require.JSONEq(t, expectBody, w.Body.String())
		s.AssertExpectations(t)
		s.AssertNumberOfCalls(t, "AverageCapacityByBrand", 1)
	})

	t.Run("case error, vehicle not found", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.AverageCapacityByBrand()
		s.On("AverageCapacityByBrand", "Ford").Return(0, internal.ErrServiceNoVehicles)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/average_capacity/brand/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("brand", "Ford")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusNotFound, w.Code)
		expectBody := `{
			"status": "Not Found",
			"message": "vehicles not found"
		}`
		require.JSONEq(t, expectBody, w.Body.String())
		s.AssertExpectations(t)
		s.AssertNumberOfCalls(t, "AverageCapacityByBrand", 1)
	})
}

func Test_handler_SearchByWeightRange(t *testing.T) {
	t.Run("case - success", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.SearchByWeightRange()
		s.On("SearchByWeightRange", mock.AnythingOfType("internal.SearchQuery"), mock.AnythingOfType("bool")).Return(
			map[int]internal.Vehicle{
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
			}, nil)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/weight/", nil)
		query := r.URL.Query()
		query.Add("weight_min", "1000")
		query.Add("weight_max", "2000")
		r.URL.RawQuery = query.Encode()
		w := httptest.NewRecorder()
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusOK, w.Code)
		expectBody := `{
			"message": "vehicles found",
			"data": {
				"1": {
					"Id": 1,
					"Brand": "Ford",
					"Model": "Fiesta",
					"Registration": "ABC-123",
					"Color": "red",
					"FabricationYear": 2010,
					"Capacity": 5,
					"MaxSpeed": 180,
					"FuelType": "gasoline",
					"Transmission": "manual",
					"Weight": 1000,
					"Height": 1.5,
					"Length": 4,
					"Width": 1.8
				}
			}
		}`
		require.JSONEq(t, expectBody, w.Body.String())

		s.AssertExpectations(t)
		s.AssertNumberOfCalls(t, "SearchByWeightRange", 1)
	})

	t.Run("case error, weight_min is not a number", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.SearchByWeightRange()
		s.On("SearchByWeightRange", mock.AnythingOfType("internal.SearchQuery"), mock.AnythingOfType("bool")).Return(
			map[int]internal.Vehicle{}, nil)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/weight/", nil)
		query := r.URL.Query()
		query.Add("weight_min", "abc")
		query.Add("weight_max", "2000")
		r.URL.RawQuery = query.Encode()
		w := httptest.NewRecorder()
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusBadRequest, w.Code)
		expectBody := `{
			"status": "Bad Request",
			"message": "invalid weight_min"
		}`
		require.JSONEq(t, expectBody, w.Body.String())

	})

	t.Run("case error, weight_max is not a number", func(t *testing.T) {
		// arrange
		s := service.NewServiceVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(s)
		h := hd.SearchByWeightRange()
		s.On("SearchByWeightRange", mock.AnythingOfType("internal.SearchQuery"), mock.AnythingOfType("bool")).Return(
			map[int]internal.Vehicle{}, nil)

		//request
		r := httptest.NewRequest(http.MethodGet, "/vehicles/weight/", nil)
		query := r.URL.Query()
		query.Add("weight_min", "1000")
		query.Add("weight_max", "abc")
		r.URL.RawQuery = query.Encode()
		w := httptest.NewRecorder()
		// act
		h(w, r)
		// assert
		require.Equal(t, http.StatusBadRequest, w.Code)
		expectBody := `{
			"status": "Bad Request",
			"message": "invalid weight_max"
		}`
		require.JSONEq(t, expectBody, w.Body.String())

	})

}
