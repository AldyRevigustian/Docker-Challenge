# Stage 1 : Builder
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ./bin/app ./main.go

# Stage 2 : Runner
FROM scratch AS runner
COPY --from=builder /app/bin/app /app/bin/app
EXPOSE 8181
ENTRYPOINT ["/app/bin/app"]