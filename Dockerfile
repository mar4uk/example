FROM golang:latest

WORKDIR /go/src/github.com/mar4uk/example
COPY . .
RUN go build -o example .

CMD ./example