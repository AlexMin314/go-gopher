FROM golang:alpine

RUN apk update
RUN apk add git
RUN apk add bash
WORKDIR /go/src/github.com/AlexMin314/go-gopher/backend

COPY . .

RUN go get ./
RUN go get github.com/codegangsta/gin

EXPOSE 8080

CMD bash start.sh