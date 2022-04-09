FROM golang:1.17.8-alpine3.15

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

COPY servertemplate ./
COPY resources ./resources

EXPOSE 8080

CMD ["servertemplate"]