############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

#FROM golang:1.17.2-stretch AS builder
# Install git.
# Git is required for fetching the dependencies.

WORKDIR $GOPATH/src/mypackage/
COPY . ./cisco-webex-bot

WORKDIR $GOPATH/src/mypackage/cisco-webex-bot/
# Fetch dependencies.
# Using go get.
RUN apk update \                                                                                                                                                                                                                        
  && apk add ca-certificates zip wget tar curl gcc musl-dev linux-headers\                                                                                                                                                                                                      
  && update-ca-certificates

RUN go install
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0  GOARCH=386 GOOS=linux go build -o /go/bin/cisco-webex-bot
RUN cd /go/bin/
WORKDIR /go/bin/

############################
# STEP 2 build a small image
############################
#FROM ubuntu 
FROM gcr.io/distroless/static
#FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/cisco-webex-bot /go/bin/cisco-webex-bot

EXPOSE 8080

LABEL version="1.0.0"
LABEL name="cisco-webex-bot"
LABEL maintainer="Andrew Pye"
LABEL description="cisco-webex-bot is a CLI application written in Golang that allows you to interact with Cisco Webex Teams."

ENV WEB_IP=localhost
ENV WEB_PORT=8080


ENTRYPOINT ["/go/bin/cisco-webex-bot","web"]