FROM golang:1.12.9

WORKDIR /src/app

RUN addgroup --system projects && adduser --system projects --ingroup projects

RUN chown -R projects:projects /src/app

USER projects

COPY . .

RUN go get -d -v ./...

