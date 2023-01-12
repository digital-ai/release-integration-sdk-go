package release_api

func (req ReleasesReleaseIdVariablesGetRequest) Execute() ([]Variable, error) {
	return req.ReleasesReleaseIdVariablesGet()
}

func (req ReleasesReleaseIdVariablesPostRequest) Execute() (*Variable, error) {
	return req.ReleasesReleaseIdVariablesPost()
}

func (req ReleasesReleaseIdVariablesPutRequest) Execute() (*Variable, error) {
	return req.ReleasesReleaseIdVariablesPut()
}

func (req ReleasesVariableIdGetRequest) Execute() (*Variable, error) {
	return req.ReleasesVariableIdGet()
}

func (req ReleasesVariableIdPutRequest) Execute() (*Variable, error) {
	return req.ReleasesVariableIdPut()
}

func (req ReleasesVariableIdDeleteRequest) Execute() error {
	return req.ReleasesVariableIdDelete()
}
