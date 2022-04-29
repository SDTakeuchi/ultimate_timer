FROM golang:1.18.1-alpine3.15

WORKDIR /go/ultimate_timer/

RUN apk update \
&&  apk add --no-cache bash \
&&  rm -rf /var/lib/apt/lists/* \
&&  export GO111MODULE=on

COPY ./app/ .

RUN go mod tidy
# && go build

EXPOSE 8080

CMD ["go", "run", "main.go", "0.0.0.0:8080"]
