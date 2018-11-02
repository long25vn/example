FROM golang:1.11-alpine

LABEL compilerimages="value"

RUN adduser -D dev
USER dev

WORKDIR /home/dev/

CMD ["/bin/cat"]
