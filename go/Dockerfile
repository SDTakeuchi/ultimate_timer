# 1st stage build for prod env
FROM golang:1.18.1-alpine3.15 as builder

WORKDIR /go/ultimate_timer/

RUN apk update \
&&  apk add --no-cache bash gcc musl-dev \
&&  rm -rf /var/lib/apt/lists/* \
&&  export GO111MODULE=on

COPY ./app/ .

RUN go mod tidy \
&&  go build -o /go_app .

# prod env; runs built image
FROM alpine as prod

COPY --from=builder /go_app .
COPY ./app/.env .

EXPOSE 8080

ENTRYPOINT [ "/go_app" ]


# dev env; live reload enabled
FROM golang:1.18.1-alpine3.15 as dev

WORKDIR /go/ultimate_timer/

RUN apk update \
&&  apk add --no-cache bash gcc musl-dev \
&&  rm -rf /var/lib/apt/lists/* \
&&  export GO111MODULE=on

COPY ./app/ .

RUN go mod tidy

EXPOSE 8080

CMD ["./start.sh"]
