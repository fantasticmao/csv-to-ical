FROM golang:1.21-alpine

RUN apk add --no-cache tzdata make git

WORKDIR /app
COPY . /app

RUN go mod download && \
    make build-docker && \
    cp ./bin/csv-to-ical-build-docker /bin/csv-to-ical && \
    mkdir /opt/csv-to-ical

CMD csv-to-ical -d /opt/csv-to-ical
