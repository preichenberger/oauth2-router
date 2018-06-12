OAuth2 Router [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/preichenberger/oauth2-router) [![Build Status](https://travis-ci.org/preichenberger/oauth2-router.svg?branch=master)](https://travis-ci.org/preichenberger/oauth2-router)
========

Redirects an OAuth2 callback to a URL specified in the state parameter with all query parameters from the original request.
This allows having one endpoint as the destination for different OAuth2 callbacks destinations. Allows for whitelisting redirect domains.

Possible use cases:
- Pull Request builds
- Heroku Review Apps
- any dynamic DNS hostname.

SECURITY: This allows for arbitrary redirects of OAuth2 Authorization Codes. Please understand the implications before using this.

## Install

### Binary
Grab a release from: https://github.com/preichenberger/oauth2-router/releases

### Docker
https://hub.docker.com/r/preichenberger/oauth2-router

### Heroku
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Start
```bash
$ oauth2-router -port 3000 -whitelist localhost,*.google.com
$ 2018/06/08 12:48:09 Starting OAuth2 Router on port: 3000
```

## Client Implementation
For the following info:
- Whitelist: localhost
- OAuth2 Router redirect_uri: http://localhost:8080
- Real redirect_uri: http://localhost:5000/random_endpoint?apple=pears

1. Add the OAuth2 Router redirect_uri: http://localhost:8080 to your OAuth2 provider (Google, Facebook, Github, etc...)
2. Create a JSON with any date to be passed in the OAuth2 state parameter. Add a redirect parameter, with the real redirect_uri
```json
{
	"csrf": "39adsijdfiaj",
	"redirect": "http://localhost:5000/random_endpoint?apple=pears"
}
```
3. Base64 encode the json and set it as the state parameter in the OAuth2 implementation
4. Use OAuth2

## Tests
Run tests
```bash
$ go ./...
```
