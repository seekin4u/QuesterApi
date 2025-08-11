FROM golang:1.22-alpine AS build
WORKDIR /app

RUN apk add --no-cache git ca-certificates && update-ca-certificates
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o backend .

FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /app
COPY --from=build /app/backend /app/backend
COPY .env ./
USER nonroot:nonroot
EXPOSE 3000
ENTRYPOINT ["/app/backend"]
