FROM golang:1.19-alpine as build

RUN apk add --no-cache make upx

WORKDIR /go/src/app
COPY . .

RUN go mod download
#RUN go vet -v
#RUN go test -v

RUN make && mv afk /go/bin/afk

FROM scratch

COPY --from=build /go/bin/afk ./afk

EXPOSE 3000
CMD ["/afk"]