FROM golang:1.17.6 AS build
WORKDIR /go/src

COPY go.mod .
COPY go.sum .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

COPY go ./go
COPY main.go .

RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime
COPY --from=build /go/src/openapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]
