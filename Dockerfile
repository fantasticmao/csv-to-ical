FROM golang:1.21-alpine3.19 AS build

RUN apk add --no-cache tzdata make git

WORKDIR /app
COPY . /app

RUN go mod download && make build-docker

FROM alpine:3.19

COPY --from=build /app/bin/csv-to-ical-build-docker /bin/csv-to-ical

RUN mkdir /opt/csv-to-ical

CMD csv-to-ical -d /opt/csv-to-ical
