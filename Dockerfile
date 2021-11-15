FROM golang:alpine
LABEL maintainer="Ángel Píñar <angle@correo.ugr.es>"
ENV CGO_ENABLED 0

WORKDIR /app/test

RUN apk update --no-cache && \
    apk add --no-cache make git && \
    adduser -D commandFTL
    
USER commandFTL

ENV GOPATH=/home/commandFTL/go

CMD make test