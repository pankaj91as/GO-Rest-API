FROM alpine:3.14

WORKDIR /usr/src/app

RUN apk update && apk add make

COPY ./build ./build
COPY ./.env .
COPY ./Makefile .

EXPOSE 8080

CMD ["make","rest-server"]
