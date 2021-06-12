FROM golang:alpine as builder
ARG PROJECT_NAME=den-arango
WORKDIR /go/src/${PROJECT_NAME}

COPY go.mod .
COPY go.sum .
RUN go get -d -v .

COPY . .
RUN go build -v -o /app .

FROM alpine as app
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /app .
EXPOSE 8800

CMD ["/app"]