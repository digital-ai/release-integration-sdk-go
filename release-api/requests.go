package release_api

import (
	"github.com/xebialabs/go-remote-runner-wrapper/http"
)

type ReleasesReleaseIdVariablesGetRequest struct {
	client    *http.HttpClient
	releaseId string
}

type ReleasesReleaseIdVariablesPostRequest struct {
	client    *http.HttpClient
	releaseId string
	variable  *Variable1
}

type ReleasesReleaseIdVariablesPutRequest struct {
	client    *http.HttpClient
	releaseId string
	variable  *Variable
}

type ReleasesVariableIdGetRequest struct {
	client     *http.HttpClient
	variableId string
}

type ReleasesVariableIdPutRequest struct {
	client     *http.HttpClient
	variableId string
	variable   *Variable
}

type ReleasesVariableIdDeleteRequest struct {
	client     *http.HttpClient
	variableId string
}
