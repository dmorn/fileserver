FROM golang:1.16.0-alpine3.13

ARG GITLAB_LOGIN
ARG GITLAB_TOKEN
ARG GO111MODULE=on
ARG CGO_ENABLED=0

WORKDIR /app

RUN apk add --no-cache git make

COPY go.mod .
# COPY go.sum .

# If you see an "unknown revision error", chances are you're not passing
# correct/valid credentials in the step above.
RUN go mod download && go mod verify

# Copy other packages if needed. Avoid copying the entire directory at this
# stage, otherwise each change will make the builder stage to recompile.
COPY cmd ./cmd

COPY Makefile ./

RUN make all
