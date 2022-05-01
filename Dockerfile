FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app main.go 

FROM golang:1.18-alpine
COPY --from=builder ./app/app .
CMD ["./app"]