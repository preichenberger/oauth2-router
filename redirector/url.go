package redirector

import (
  "net/url"
)

func CreateUrl(queryValues url.Values) (*url.URL, error) {
	if _, ok := queryValues["state"]; !ok {
    return nil, ValidationError{"Missing state field"}
	}

	stateValues, err := url.ParseQuery(queryValues["state"][0])
	if err != nil {
    return nil, ValidationError{"State field is not a query string"}
	}

	if _, ok := stateValues["redirect"]; !ok {
    return nil, ValidationError{"Query param redirect missing from state"}
	}

	redirectUrl, err := url.ParseRequestURI(stateValues["redirect"][0])
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
