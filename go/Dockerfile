FROM golang:1.18.1-stretch

WORKDIR /go/ultimate_timer/

RUN apt-get update \
&& export GO111MODULE=on

COPY ./app/ .

RUN go mod tidy
# && go build

EXPOSE 8080

CMD ["go", "run", "main.go", "0.0.0.0:8080"]
# CMD ["gin", "--immediate", "run", "main.go", "0.0.0.0:8080"]
