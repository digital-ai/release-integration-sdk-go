/*
Digital.ai Release API

Testing EnvironmentReservationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/xebialabs/go-remote-runner-wrapper/api/release/openapi"
	"testing"
)

func Test_openapi_EnvironmentReservationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EnvironmentReservationApiService AddApplication", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentReservationId string

		resp, err := apiClient.EnvironmentReservationApi.AddApplication(context.Background(), environmentReservationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test EnvironmentReservationApiService CreateReservation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EnvironmentReservationApi.CreateReservation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentReservationApiService DeleteEnvironmentReservation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentReservationId string

		resp, err := apiClient.EnvironmentReservationApi.DeleteEnvironmentReservation(context.Background(), environmentReservationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test EnvironmentReservationApiService GetReservation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentReservationId string

		resp, httpRes, err := apiClient.EnvironmentReservationApi.GetReservation(context.Background(), environmentReservationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentReservationApiService SearchReservations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EnvironmentReservationApi.SearchReservations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentReservationApiService UpdateReservation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentReservationId string

		resp, httpRes, err := apiClient.EnvironmentReservationApi.UpdateReservation(context.Background(), environmentReservationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}