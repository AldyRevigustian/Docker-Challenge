# Stage 1 : Build
FROM alpine:latest AS builder 

RUN apk add --no-cache g++
WORKDIR /app
COPY main.cpp .

RUN g++ -Os -s -static main.cpp -o main

# Stage 2 : Runner
FROM scratch AS runner
COPY --from=builder /app/main /
CMD ["/main"]