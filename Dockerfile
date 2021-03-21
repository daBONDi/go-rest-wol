FROM golang:1.16-alpine AS builder

LABEL org.label-schema.vcs-url="https://github.com/daBONDi/go-rest-wol" \
      org.label-schema.url="https://github.com/daBONDi/go-rest-wol/blob/master/README.md"

WORKDIR /app

ADD . .

# Install Dependecies
RUN apk update && apk upgrade && \
    apk add --no-cache git

# Build Source Files
RUN go build -o go-rest-wol . 

# Create 2nd Stage final image
FROM alpine

WORKDIR /app

COPY --from=builder /app/pages/index.html ./pages/index.html
COPY --from=builder /app/computer.csv ./computer.csv
COPY --from=builder /app/go-rest-wol .

ENV WOLHTTPPORT=8080
ENV WOLFILE=computer.csv

EXPOSE ${WOLHTTPPORT}

ENTRYPOINT ["/app/go-rest-wol"]