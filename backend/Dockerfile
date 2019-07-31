
FROM golang:latest

WORKDIR /go/src/github.com/AlexMin314/go-gopher/backend

# COPY . .
ADD . .

RUN go get ./
RUN go get github.com/codegangsta/gin

EXPOSE 8080
EXPOSE 3000

# CMD go run server.go
CMD bash start.sh