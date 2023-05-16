FROM golang:1.20 as build

WORKDIR /service

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o boilerplate

FROM alpine:3.17

RUN apk add --no-cache ca-certificates

WORKDIR /service

COPY --from=build /service/boilerplate .
COPY --from=build /service/config.yaml .

RUN ln -s /service/boilerplate /usr/local/bin/boilerplate

ENTRYPOINT ["core-service"]
