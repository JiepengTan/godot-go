package extensionapiparser

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func strictUnmarshal(data []byte, v interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	return dec.Decode(v)
}

// ParseGdextensionApiJson parses gdextension_api.json into a APIJson struct.
func ParseExtensionApiJson(apiStr string) (ExtensionApi, error) {
	// Unmarshal the JSON into our struct.
	var extensionApiJson ExtensionApi
	if err := strictUnmarshal([]byte(apiStr), &extensionApiJson); err != nil {
		return ExtensionApi{}, err
	}

	return extensionApiJson, nil
}

func GenerateExtensionAPI(apiStr, buildConfig string) (ExtensionApi, error) {
	var (
		eapi ExtensionApi
		err  error
	)
	if eapi, err = ParseExtensionApiJson(apiStr); err != nil {
		return ExtensionApi{}, err
	}
	if !eapi.HasBuildConfiguration(buildConfig) {
		return ExtensionApi{}, fmt.Errorf(`unable to find build configuration "%s"`, buildConfig)
	}
	eapi.BuildConfig = buildConfig
	eapi.Classes = eapi.FilteredClasses()

	return eapi, nil
}
