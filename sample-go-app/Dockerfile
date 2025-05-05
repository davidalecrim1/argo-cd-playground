FROM --platform=arm64 golang:1.24.2-alpine AS builder
WORKDIR /app/
COPY . .
RUN go build -o main .

FROM --platform=arm64 alpine:3.20
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8081
CMD ["./main"]