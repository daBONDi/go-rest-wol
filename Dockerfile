FROM golang:alpine AS builder

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

CMD ["/app/main"]

EXPOSE 8080
