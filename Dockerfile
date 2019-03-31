FROM alpine:3.7

RUN apk add --update --no-cache \
    ca-certificates \
    && rm -rf /var/cache/apk/*

COPY bitmap /usr/local/bin/bitmap

WORKDIR /usr/local/var/bitmap

EXPOSE 3000
VOLUME ["/Users/ken/Documents/gopath/src/play/bitmap/data", "/usr/local/var/bitmap"]

CMD ["/usr/local/bin/bitmap", "config.ini"]
