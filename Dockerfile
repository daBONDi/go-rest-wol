FROM golang:alpine3.16 AS builder

LABEL org.label-schema.vcs-url="https://github.com/daBONDi/go-rest-wol" \
      org.label-schema.url="https://github.com/daBONDi/go-rest-wol/blob/master/README.md"

RUN mkdir /app
ADD . /app/
WORKDIR /app

# Install Dependencies
RUN apk update && apk upgrade && \
    apk add --no-cache git && \
    go get -d github.com/gorilla/handlers@v1.5.1 && \
    go get -d github.com/gorilla/mux@v1.8.0 && \
    go get -d github.com/gocarina/gocsv@v0.0.0-20220727205534-7fbf8e1b37fb

# Build Source Files
RUN go build 

# Create 2nd Stage final image
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/pages/index.html ./pages/index.html
COPY --from=builder /app/computer.csv .
COPY --from=builder /app/go-rest-wol .

ENV WOLHTTPPORT=8080
ENV WOLFILE=computer.csv

CMD ["/app/go-rest-wol"]

EXPOSE ${WOLHTTPPORT}
