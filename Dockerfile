FROM golang:alpine3.19

ENV APPLICATION_PACKAGE=./cmd/go-scaffolding
ENV APP_PATH=/app
ENV GO_TEST_FLAGS="-tags=integration -p=1"
ENV SCOPE=local

RUN apk update && apk add --no-cache bash

RUN apk add nodejs npm

RUN	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.56.2/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

RUN go install github.com/arthurgustin/godepth@latest
RUN go install github.com/fdaines/arch-go@v1.4.0
RUN npm -g install directory-validator

ADD ./commands/*.sh /commands/
ADD .golangci-project.yml /commands/
ADD arch-go.yml /commands/
ADD .directoryvalidator.json /commands/
RUN chmod a+x /commands/*.sh
ADD ./resources /resources
ADD . .

RUN chmod +x /commands/*.sh

ENV GO_RACE_DETECTION_ENABLED="true"