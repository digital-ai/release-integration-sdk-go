/*
Digital.ai Release API

Testing FolderApiService

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

func Test_openapi_FolderApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FolderApiService AddFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.FolderApi.AddFolder(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService CreateFolderVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.FolderApi.CreateFolderVariable(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService DeleteFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, err := apiClient.FolderApi.DeleteFolder(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test FolderApiService DeleteFolderVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string
		var variableId string

		resp, err := apiClient.FolderApi.DeleteFolderVariable(context.Background(), folderId, variableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test FolderApiService Find", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FolderApi.Find(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService GetFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.FolderApi.GetFolder(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService GetFolderPermissions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FolderApi.GetFolderPermissions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService GetFolderTeams", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.FolderApi.GetFolderTeams(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService GetFolderTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.FolderApi.GetFolderTemplates(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService GetFolderVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string
		var variableId string

		resp, httpRes, err := apiClient.FolderApi.GetFolderVariable(context.Background(), folderId, variableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService IsFolderOwner", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.FolderApi.IsFolderOwner(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService List", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.FolderApi.List(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService ListRoot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FolderApi.ListRoot(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService ListVariableValues", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.FolderApi.ListVariableValues(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService ListVariables", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.FolderApi.ListVariables(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService Move", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, err := apiClient.FolderApi.Move(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test FolderApiService MoveTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string
		var templateId string

		resp, err := apiClient.FolderApi.MoveTemplate(context.Background(), folderId, templateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test FolderApiService RenameFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, err := apiClient.FolderApi.RenameFolder(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test FolderApiService SearchReleasesFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.FolderApi.SearchReleasesFolder(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService SetFolderTeams", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string

		resp, httpRes, err := apiClient.FolderApi.SetFolderTeams(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FolderApiService UpdateFolderVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId string
		var variableId string

		resp, httpRes, err := apiClient.FolderApi.UpdateFolderVariable(context.Background(), folderId, variableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}