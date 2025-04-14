FROM golang:1.24.2 as builder

ENV THIS_SERVICE=inventory

WORKDIR /app

COPY go.work go.work.sum ./
COPY pkg/ ./pkg
COPY apps/ ./apps/

RUN go work sync

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/service ./apps/$THIS_SERVICE/cmd/main.go

RUN chmod +x /app/service


FROM gcr.io/distroless/static as runner

WORKDIR /app

COPY --from=builder /app/service /app/service

ENTRYPOINT ["/app/service"]