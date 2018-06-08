package redirector

import (
  "encoding/base64"
  "encoding/json"
  "net/url"
)

func CreateUrl(queryValues url.Values) (*url.URL, error) {
  var stateValues map[string]string

	if _, ok := queryValues["state"]; !ok {
    return nil, ValidationError{"Missing state field"}
	}

	stateQuery, err := base64.StdEncoding.DecodeString(queryValues["state"][0])
	if err != nil {
		return nil, ValidationError{"Could not base64 decode state value"}
	}

  if err := json.Unmarshal(stateQuery, &stateValues); err != nil {
    return nil, ValidationError{"Could not json decode state value"}
  }

	if _, ok := stateValues["redirect"]; !ok {
    return nil, ValidationError{"Query param redirect missing from state"}
	}

  redirect, err := url.QueryUnescape(stateValues["redirect"])
	if err != nil {
    return nil, ValidationError{"Could not URL decode redirect value"}
	}

	redirectUrl, err := url.ParseRequestURI(redirect)
	if err != nil {
    return nil, ValidationError{"Could not parse redirect URL"}
	}

	redirectValues := redirectUrl.Query()
	for key, values := range queryValues {
		for _, value := range values {
			redirectValues.Add(key, value)
		}
	}

	redirectUrl.RawQuery = redirectValues.Encode()
  return redirectUrl, nil
}
