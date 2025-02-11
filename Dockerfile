FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/app/main"]
