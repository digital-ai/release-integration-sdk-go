/*
Digital.ai Release API

Testing TaskApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/digital-ai/release-integration-sdk-go/api/release/openapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_TaskApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaskApiService AbortTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.AbortTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService AddAttachments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.AddAttachments(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService AddCondition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.AddCondition(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService AddDependency", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var targetId string
		var taskId string

		resp, httpRes, err := apiClient.TaskApi.AddDependency(context.Background(), targetId, taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService AddTaskTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var containerId string

		resp, httpRes, err := apiClient.TaskApi.AddTaskTask(context.Background(), containerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService AssignTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string
		var username string

		resp, httpRes, err := apiClient.TaskApi.AssignTask(context.Background(), taskId, username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService ChangeTaskType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.ChangeTaskType(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService CommentTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.CommentTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService CompleteTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.CompleteTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService CopyTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.CopyTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService DeleteAttachment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var attachmentId string
		var taskId string

		resp, err := apiClient.TaskApi.DeleteAttachment(context.Background(), attachmentId, taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test TaskApiService DeleteCondition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var conditionId string

		resp, err := apiClient.TaskApi.DeleteCondition(context.Background(), conditionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test TaskApiService DeleteDependency", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dependencyId string

		resp, err := apiClient.TaskApi.DeleteDependency(context.Background(), dependencyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test TaskApiService DeleteTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, err := apiClient.TaskApi.DeleteTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test TaskApiService FailTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.FailTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService GetTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.GetTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService GetTaskVariables", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.GetTaskVariables(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService LockTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, err := apiClient.TaskApi.LockTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test TaskApiService ReopenTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.ReopenTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService RetryTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.RetryTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService SearchTasksByTitle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TaskApi.SearchTasksByTitle(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService SkipTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.SkipTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService StartTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.StartTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService StartTask1", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.StartTask1(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService UnlockTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, err := apiClient.TaskApi.UnlockTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

	})

	t.Run("Test TaskApiService UpdateCondition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var conditionId string

		resp, httpRes, err := apiClient.TaskApi.UpdateCondition(context.Background(), conditionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService UpdateInputVariables", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.UpdateInputVariables(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskApiService UpdateTask", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskApi.UpdateTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
