# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /gomx_template_server

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY app/. ./app/

RUN CGO_ENABLED=0 GOOS=linux go build -o /build

EXPOSE 8081

CMD ["/build"]
