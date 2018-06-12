package redirector

import (
  "strings"
)

type Redirector struct {
  WhitelistDomains []string
}

func validateDomainParts(a, b []string) bool {
  if len(a) != len(b) {
    return false
  }

  for i, _ := range a {
    if a[i] == "*" || b[i] == "*" {
      continue
    }

    if a[i] != b[i] {
      return false
    }
  }

  return true
}

func NewRedirector(whitelist string) *Redirector {
  whitelistDomains := strings.Split(whitelist, ",")

  return &Redirector{
    WhitelistDomains: whitelistDomains,
  }
}

func (r *Redirector) ValidateDomain(host string) bool {
  for _, whitelistDomain := range r.WhitelistDomains {
    whitelistDomainParts := strings.Split(whitelistDomain, ".")
    hostParts := strings.Split(host, ".")
    if whitelistDomain == "" {
      return true
    }

    if !validateDomainParts(whitelistDomainParts, hostParts) {
      continue
    }

    return true
  }

  return false
}
