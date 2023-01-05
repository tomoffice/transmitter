FROM golang:alpine AS builder
WORKDIR /src
COPY ./src .
#build dockerTest to heathcheck
RUN go build dockerTest.go
#build main app to run
RUN go build main.go 

FROM alpine
WORKDIR /App
COPY --from=builder /src/dockerTest .
COPY --from=builder /src/main .
RUN apk update && apk add tzdata
#ENTRYPOINT [ "./main" ]