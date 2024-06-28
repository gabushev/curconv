FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o curconv cmd/api/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/curconv .
EXPOSE 3000
ENTRYPOINT ["./curconv"]
