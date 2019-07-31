FROM golang:latest

WORKDIR /go/src/github.com/user/go-gopher/backend

COPY . .

RUN go get ./
RUN go get github.com/codegangsta/gin

EXPOSE 8080

CMD bash start.sh