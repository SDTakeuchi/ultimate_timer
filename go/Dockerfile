FROM golang:1.18.0-stretch

WORKDIR /go/ultimate_timer

RUN export GO111MODULE=auto \
&& apt-get update

COPY ./app .

RUN go mod init ultimate_timer \
&& go get . \
&& go build
# && go install github.com/cosmtrek/air@v1.27.3

EXPOSE 8080

CMD ["go", "run", "main.go", "0.0.0.0:8080"]
# CMD ["air"]