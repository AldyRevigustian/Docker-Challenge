# Stage builder
FROM node:22-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Stage runner
FROM busybox:latest AS runner
WORKDIR /app
COPY --from=builder /app/dist /app/dist
EXPOSE 8484

CMD ["busybox", "httpd", "-f", "-p", "8484", "-h", "/app/dist"]