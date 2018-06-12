package redirector

import (
  "encoding/base64"
  "encoding/json"
  "errors"
  "fmt"
  "net/url"
  "testing"
)

func createStateValue(values map[string]string) (string, error) {
  data, err := json.Marshal(values)
  if err != nil {
    return "", err
  }

  return base64.StdEncoding.EncodeToString(data), nil
}

func TestCreateUrl(t *testing.T) {
  queryValues := url.Values{}
  queryValues.Add("code", "3")
  queryValues.Add("code", "5")

  stateValues := map[string]string{
    "csrf": "test123",
    "redirect": "https://www.github.com/test?callback=3&callback=4",
  }

  state, err := createStateValue(stateValues)
  if err != nil {
    t.Error(err)
  }
  queryValues.Add("state", state)

  redirector := NewRedirector("www.github.com")
  redirectUrl, err := redirector.CreateUrl(queryValues)
  if err != nil {
    t.Error(err)
  }

  expected := fmt.Sprintf("https://www.github.com/test?callback=3&callback=4&code=3&code=5&state=%s",
    url.QueryEscape(state))

  if redirectUrl.String() != expected {
    errMsg := fmt.Sprintf("CreateUrl does not match expected response:\n  Result: %s\nExpected: %s", redirectUrl.String(), expected)
    t.Error(errors.New(errMsg))
  }
}

func TestCreateUrlMissingStateError(t *testing.T) {
  redirector := NewRedirector("")

  _, err := redirector.CreateUrl(url.Values{})
  if err.Error() != "Missing state field" {
    t.Error(errors.New("Missing not found state error"))
  }
}

func TestCreateUrlMissingRedirectError(t *testing.T) {
  stateValues := map[string]string{
    "test": "test",
  }

  queryValues := url.Values{}
  state, err := createStateValue(stateValues)
  if err != nil {
    t.Error(err)
  }
  queryValues.Add("state", state)

  redirector := NewRedirector("")
  _, err = redirector.CreateUrl(queryValues)
  if err.Error() != "Query param redirect missing from state" {
    t.Error(errors.New("Missing query param redirect error"))
  }
}

func TestCreateUrlRedirectParseError(t *testing.T) {
  stateValues := map[string]string{
    "redirect": "http;test",
  }

  queryValues := url.Values{}
  state, err := createStateValue(stateValues)
  if err != nil {
    t.Error(err)
  }
  queryValues.Add("state", state)

  redirector := NewRedirector("")
  _, err = redirector.CreateUrl(queryValues)
  if err.Error() != "Could not parse redirect URL" {
    t.Error(errors.New("Missing could not parse redirect error"))
  }
}
