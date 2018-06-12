package redirector

import (
  "errors"
  "testing"
)

func TestValidateDomain(t *testing.T) {
  r := NewRedirector("")
  if !r.ValidateDomain("pizza") {
    t.Error(errors.New("Empty whitelist failed"))
  }
}

func TestValidateDomainWildcard(t *testing.T) {
  r := NewRedirector("google.com,apple.com,*.github.com")
  if !r.ValidateDomain("api.github.com") {
    t.Error(errors.New("Wildcard domain validation failed"))
  }
}

func TestValidateDomainNested(t *testing.T) {
  r := NewRedirector("google.com,*.api.apple.com,*.github.com")
  if !r.ValidateDomain("test.api.apple.com") {
    t.Error(errors.New("Nested domain validation failed"))
  }
}

func TestValidateDomainFail(t *testing.T) {
  r := NewRedirector("google.com,*.api.apple.com,*.github.com")
  if r.ValidateDomain("test.api.github.com") {
    t.Error(errors.New("Validate domain did not fail"))
  }
  if r.ValidateDomain("api.apple.com") {
    t.Error(errors.New("Validate domain did not fail"))
  }
  if r.ValidateDomain("test.google.com") {
    t.Error(errors.New("Validate domain did not fail"))
  }
}
