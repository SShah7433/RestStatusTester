FROM golang:latest
RUN mkdir /app
WORKDIR /app
COPY src /app
RUN go get -u github.com/gin-gonic/gin && go build -o RestStatusTester .