package launcher

import (
	"encoding/json"
	"go-mclauncher/api"
	"go-mclauncher/entities/launcher"
)

func X19LoginOTP(sAuth *launcher.MPaySAuthToken) (*api.RequestConfig, error) {
	// Construct the request parameters
	sAuthJson, err := json.Marshal(sAuth)
	if err != nil {
		return nil, err
	}

	return &api.RequestConfig{
		URL:      "/login-otp",
		HostName: api.CoreServerURL,
		Method:   "POST",
		Params:   sAuthJson,
	}, nil
}
