package helper

import (
	"net/url"
	"path"
	"strings"
)

func PlanAddressForHTTPBackend(BackendConfig map[string]interface{}) (string, error) {

	if planAddress, ok := BackendConfig["plan_address"]; ok {
		return planAddress.(string), nil
	}

	address := BackendConfig["address"].(string)
	// As Terraform HTTP backend don't implement workspaces, we need to define a
	// different URL to store the plan file.
	u, err := url.Parse(address)
	if err != nil {
		return "", err
	}

	// Parse the URL and add -plan suffix to the filename
	baseURL, fFullName := path.Split(u.Path)
	fExt := path.Ext(u.Path)
	fName := strings.TrimSuffix(fFullName, fExt)

	fFullName = fName + "-plan" + fExt
	u.Path = path.Join(baseURL, fFullName)
	return u.String(), nil
}
