FROM alpine:3.12

RUN apk add --no-cache \
    bash

COPY bin/client/client .

RUN chmod +x /client

ENTRYPOINT ["/client"]
