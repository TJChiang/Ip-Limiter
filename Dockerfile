FROM golang:1.20-alpine
LABEL authors="nature.chiang"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
