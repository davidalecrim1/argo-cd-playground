FROM --platform=arm64 golang:1.24.2-alpine

WORKDIR /app/

COPY . .

RUN go build -o main .

EXPOSE 8081

CMD [ "./main" ]