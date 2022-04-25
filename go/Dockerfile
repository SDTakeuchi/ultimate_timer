FROM golang:1.18.1-stretch

WORKDIR /go/ultimate_timer/

RUN apt-get update \
&& export GO111MODULE=auto

COPY ./app/ .

# TODO: separate stg env and prd env
RUN go install github.com/codegangsta/gin \
&&  go mod tidy
# && go build

EXPOSE 8080

CMD ["go", "run", "main.go", "0.0.0.0:8080"]
# CMD ["gin", "--immediate", "run", "main.go", "0.0.0.0:8080"]
