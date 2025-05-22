# Stage 1: Composer dependencies
FROM composer:latest AS vendor

WORKDIR /app
COPY composer.json composer.lock* ./

RUN composer install \
    --no-interaction \
    --no-plugins \
    --no-scripts \
    --no-dev \
    --prefer-dist \
    && composer clear-cache

# Stage 2 : Frontend dependecnies
FROM node:23.11.1-alpine AS frontend

WORKDIR /app
COPY resources/css ./resources/css
COPY resources/js ./resources/js
COPY package*.json vite.config.js ./

RUN mkdir -p /app/public \
    && npm install \
    && npm run build \
    && npm cache clean --force

# Stage 3: PHP Extension
FROM php:8.3-fpm-alpine

RUN apk add --no-cache libzip \
    && apk add --no-cache --virtual .build-deps libzip-dev $PHPIZE_DEPS \
    && docker-php-ext-install -j$(nproc) pdo_mysql zip opcache \
    && apk del .build-deps \
    && rm -rf /var/cache/apk/*

WORKDIR /var/www/html
COPY . .

RUN rm -rf tests/ resources/js/ resources/css/ node_modules/ README.md \
    package.json package-lock.json phpunit.xml \
    vite.config.js docker-compose.yml nginx.conf

RUN chown -R www-data:www-data /var/www/html

COPY --from=vendor /app/vendor ./vendor
COPY --from=frontend /app/public/build ./public/build
