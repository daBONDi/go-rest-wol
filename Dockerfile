FROM golang:1.11.2
# Install Dependecies
RUN go get -d github.com/gorilla/handlers
RUN go get -d github.com/gorilla/mux
RUN go get -d github.com/gocarina/gocsv

# Copy all Source Files
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main . 
CMD ["/app/main"]

EXPOSE 8080
