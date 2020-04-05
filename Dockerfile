FROM golang:alpine AS builder

LABEL org.label-schema.vcs-url="https://github.com/daBONDi/go-rest-wol" \
      org.label-schema.url="https://github.com/daBONDi/go-rest-wol/blob/master/README.md"

RUN mkdir /app
ADD . /app/
WORKDIR /app

# Install Dependecies
RUN apk update && apk upgrade && \
    apk add --no-cache git && \
    go get -d github.com/gorilla/handlers && \
    go get -d github.com/gorilla/mux && \
    go get -d github.com/gocarina/gocsv

# Build Source Files
RUN go build -o main .

# Create 2nd Stage final image
FROM alpine
WORKDIR /app
COPY --from=builder /app/pages/index.html ./pages/index.html
COPY --from=builder /app/computer.csv .
COPY --from=builder /app/main .

ARG WOLHTTPPORT=8080
ARG WOLFILE=computer.csv

CMD ["/app/main"]

EXPOSE ${WOLHTTPPORT}

RUN apk add --no-cache samba samba-common-tools