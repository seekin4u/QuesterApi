FROM golang:latest as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
#COPY .env ./
#COPY controllers/* ./
#COPY initializers/* ./
#COPY migrate/migrate.go ./
#COPY models/QuestModel.go ./
#COPY main.go ./

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build

FROM scratch

WORKDIR /app
COPY --from=builder /app/.env .
COPY --from=builder /app/QuesterApi .

EXPOSE 3000

CMD ["/app/QuesterApi"]