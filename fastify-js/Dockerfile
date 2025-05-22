# Stage 1 : Build
FROM node:22-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

# Stage 2 : Runner
FROM node:22-alpine AS runner
WORKDIR /app
COPY --from=build /app .
RUN npm install --omit=dev
EXPOSE 8080

CMD ["node", "server.js"]