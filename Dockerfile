FROM golang:1.19-alpine as build

RUN apk add --no-cache \
    build-base \
    gcc

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v
RUN go test -v

RUN go build -o /go/bin/afk

FROM alpine:latest

WORKDIR /app
COPY --from=build /go/src/app/assets ./assets
COPY --from=build /go/bin/afk ./afk

EXPOSE 3000
CMD ["/app/afk"]
