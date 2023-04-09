FROM golang:alpine AS build

RUN apk add --no-cache build-base
RUN apk add --no-cache make
RUN apk add --no-cache git

WORKDIR /app
COPY . .
RUN make build

FROM alpine
WORKDIR /app
COPY --from=build /app/build/bin/accounting-daily-tasks /app

ENTRYPOINT [ "/app/accounting-daily-tasks" ][]