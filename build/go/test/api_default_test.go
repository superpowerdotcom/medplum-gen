/*
Medplum - OpenAPI 3.0

Testing DefaultAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package medplum

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_medplum_DefaultAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultAPIService CreateResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var resourceType string

		resp, httpRes, err := apiClient.DefaultAPI.CreateResource(context.Background(), resourceType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService DeleteResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var resourceType string
		var id string

		httpRes, err := apiClient.DefaultAPI.DeleteResource(context.Background(), resourceType, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService PatchResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var resourceType string
		var id string

		httpRes, err := apiClient.DefaultAPI.PatchResource(context.Background(), resourceType, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ReadResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var resourceType string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.ReadResource(context.Background(), resourceType, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ReadResourceHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var resourceType string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.ReadResourceHistory(context.Background(), resourceType, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService ReadVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var resourceType string
		var id string
		var versionId string

		resp, httpRes, err := apiClient.DefaultAPI.ReadVersion(context.Background(), resourceType, id, versionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService Search", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var resourceType string

		resp, httpRes, err := apiClient.DefaultAPI.Search(context.Background(), resourceType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultAPIService UpdateResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var resourceType string
		var id string

		resp, httpRes, err := apiClient.DefaultAPI.UpdateResource(context.Background(), resourceType, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}