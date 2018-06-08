FROM golang:1.10.2-alpine3.7 as build
WORKDIR /go/src/github.com/preichenberger/oauth2-router
ADD . ./
RUN go build

FROM alpine:3.7
COPY --from=build /go/src/github.com/preichenberger/oauth2-router/oauth2-router /usr/local/bin/oauth2-router

EXPOSE 8080
ENTRYPOINT ["oauth2-router"]
