FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY ./html ./html
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

FROM alpine:3.18 as runtime
WORKDIR /
COPY --from=builder /server /server
RUN mkdir html
COPY --from=builder /app/html /html
EXPOSE 8080
ENTRYPOINT ["/server"]
