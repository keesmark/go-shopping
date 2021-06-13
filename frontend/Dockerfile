FROM node:14.17-alpine

WORKDIR /app
COPY . /app

RUN apk update
RUN npm install -g npm
RUN npm install -g @vue/cli


ENV HOST 0.0.0.0

EXPOSE 3000
