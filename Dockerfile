FROM golang:1.14 AS builder
WORKDIR /khumu
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o khumu-comment main.go

FROM alpine
COPY --from=builder /khumu/khumu-comment /khumu/khumu-comment
ENV KHUMU_HOME /khumu
ENV KHUMU_ENVIRONMENT DEV
CMD ["./khumu/khumu-comment"]