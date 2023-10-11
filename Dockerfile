FROM golang:1.21.1-alpine AS build

WORKDIR /orders
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o orders

FROM alpine:latest

WORKDIR /orders
COPY --from=build /orders/orders /orders/orders

EXPOSE 8094

CMD ["./orders"]
