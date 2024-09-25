# Build stage
FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /carehub-backend ./cmd/api/

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=build /carehub-backend ./

EXPOSE 8080
CMD ["./carehub-backend"]
