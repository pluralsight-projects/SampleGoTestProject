FROM golang:1.12.9-alpine3.9

ENV APP_DIR /SampleGoTestProject

RUN mkdir -p $APP_DIR

WORKDIR ${APP_DIR}