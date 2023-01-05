FROM golang:alpine AS builder
WORKDIR /src
COPY ./src .
RUN go build dockerTest.go

FROM alpine
WORKDIR /App
COPY --from=builder /src/dockerTest .
RUN apk update && apk add tzdata
#ENTRYPOINT [ "./exec" ]

