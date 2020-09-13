FROM node:14-alpine

RUN apk update && apk upgrade

COPY . /build/apiGateway
WORKDIR /build/apiGateway

RUN yarn install --production=false

EXPOSE 5003
ENTRYPOINT yarn start


