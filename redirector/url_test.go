package redirector

import (
  "errors"
  "fmt"
  "net/url"
  "testing"
)

func TestCreateUrl(t *testing.T) {
  csrfString := "test313"
  redirectUrlString := "https://www.github.com/test?callback=3&callback=4"

  queryValues := url.Values{}
  queryValues.Add("code", "3")
  queryValues.Add("code", "5")

  stateQueryValues := url.Values{}
  stateQueryValues.Add("csrf", csrfString)
  stateQueryValues.Add("redirect", redirectUrlString)
  println(stateQueryValues.Encode())
  queryValues.Add("state", stateQueryValues.Encode())

  redirectUrl, err := CreateUrl(queryValues)
  if err != nil {
    t.Error(err)
  }

  expected := fmt.Sprintf("https://www.github.com/test?callback=3&callback=4&code=3&code=5&state=%s",
    url.QueryEscape(stateQueryValues.Encode()))
  if redirectUrl.String() != expected {
    errMsg := fmt.Sprintf("CreateUrl does not match expected response:\nResult:%s\nExpected: %s", redirectUrl.String(), expected)
    t.Error(errors.New(errMsg))
  }
}

func TestCreateUrlMissingStateError(t *testing.T) {
  _, err := CreateUrl(url.Values{})
  if err.Error() != "Missing state field" {
    t.Error(errors.New("Missing not found state error"))
  }
}

func TestCreateUrlMissingRedirectError(t *testing.T) {
  queryValues := url.Values{}
  queryValues.Add("state", "test=3")

  _, err := CreateUrl(queryValues)
  if err.Error() != "Query param redirect missing from state" {
    t.Error(errors.New("Missing query param redirect error"))
  }
}

func TestCreateUrlRedirectParseError(t *testing.T) {
  queryValues := url.Values{}
  queryValues.Add("state", "redirect=http/test")

  _, err := CreateUrl(queryValues)
  if err.Error() != "Could not parse redirect URL" {
    t.Error(errors.New("Missing could not parse redirect error"))
  }
}
