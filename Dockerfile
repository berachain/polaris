# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY *.go ./
LABEL org.opencontainers.image.source=https://github.com/berachain/stargazer
