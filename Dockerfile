FROM golang:alpine AS builder
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN GOOS=linux GOARCH=arm64 go build -o homeworkServer .


FROM alpine AS runner
WORKDIR /app
COPY --from=builder /app/homeworkServer .
EXPOSE 8080
ENTRYPOINT ["./homeworkServer"]