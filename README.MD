OAuth2 Router [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/preichenberger/oauth2-router) [![Build Status](https://travis-ci.org/preichenberger/go-gdax.svg?branch=master)](https://travis-ci.org/preichenberger/oauth2-router)
========

Redirects an OAuth2 callback to a URL specified in the state parameter with all query parameters from the original request.
This allows having one endpoint as the destination for different OAuth2 callbacks destinations.

Possible use cases:
- Pull Request builds
- Heroku Review Apps
- any dynamic DNS hostname.

SECURITY: This allows for arbitrary redirects of OAuth2 Authorization Codes. Please understand the implications before using this.

## Install
Grab a release from: https://github.com/preichenberger/oauth2-router/releases

## Start
```bash
$ oauth2-router -port 3000
$ 2018/06/08 12:48:09 Starting OAuth2 Router on port: 3000
```

## Example
Using the following information:
- OAuth2 Router redirect_uri: http://localhost:8080
- Real redirect_uri: http://localhost:5000/random_endpoint?apple=pears
- Authorization code from callback: secretcode
- Original state parameter(URL encoded): csrf%3D1234
- OAuth2 Router state parameter(URL encoded): csrf%3D1234%26redirect%3Dhttp%3A%2F%2Flocalhost%3A5000%2Frandom_endpoint%3Fapple%3Dpears

Setting the state parameter must be implemented by the user's code. It must include a query parameter 'redirect' which is the real redirect_uri that you will be redirecting to.

```bash
curl "http://localhost:8080?code=secretcode&state=csrf%3D1234%26redirect%3Dhttp%3A%2F%2Flocalhost%3A5000%2Frandom_endpoint%3Fapple%3Dpears"
<a href="http://localhost:5000/random_endpoint?apple=pears&amp;code=secretcode&amp;state=csrf%3D1234%26redirect%3Dhttp%3A%2F%2Flocalhost%3A5000%2Frandom_endpoint%3Fapple%3Dpears">Moved Permanently</a>.
```

## Tests
Run tests
```bash
$ go ./...
```
