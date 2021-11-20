# How to build:
# docker build -t web-api-golang -f Dockerfile .

# This is a recipe for go binaries in a scratch docker container with up-to-date tls certs and timezone data.

## We'll choose the incredibly lightweight
## Go alpine image to work with
FROM golang as golang
## We create an /app directory in which
## we'll put all of our project code
RUN mkdir /app
ADD . /app
WORKDIR /app
## We want to build our application's binary executable
RUN make all

FROM alpine:latest as alpine
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
# -0 means no compression.  Needed because go's
# tz loader doesn't handle compressed data.
RUN zip -q -r -0 /zoneinfo.zip .

## the lightweight scratch image we'll run our application within
#FROM busybox:1.28.4 AS production
FROM scratch AS production
#FROM alpine:latest as production
ARG app_name
WORKDIR /app
## We have to copy the output from our
## builder stage to our production stage
COPY --from=golang /app/main main
ENV ZONEINFO /zoneinfo.zip
COPY --from=alpine /zoneinfo.zip /
# the tls certificates:
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
## we can then kick off our newly compiled
## binary exectuable!!
CMD ["./main"]
