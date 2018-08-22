FROM golang:alpine as builder

WORKDIR $GOPATH/src/exchange-currency

COPY Gopkg.toml Gopkg.lock ./

RUN apk update && apk add git

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN dep ensure --vendor-only
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
FROM scratch
COPY --from=builder go/src/exchange-currency/main .
COPY --from=builder go/src/exchange-currency/config.yaml .
RUN goose up

CMD ["./main"]
