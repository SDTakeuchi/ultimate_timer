FROM golang:1.17.8-stretch

WORKDIR /go/src/app

RUN export GO111MODULE=auto \
&& apt-get update

COPY ./app/main.go .

RUN go mod init ultimate_timer \
&& go get . \
&& go build

EXPOSE 8080

CMD ["go", "run", "main.go", "0.0.0.0:8080"]