/*
Digital.ai Release API

Testing RiskAssessmentApiService

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

func Test_openapi_RiskAssessmentApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RiskAssessmentApiService GetAssessment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var riskAssessmentId string

		resp, httpRes, err := apiClient.RiskAssessmentApi.GetAssessment(context.Background(), riskAssessmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}