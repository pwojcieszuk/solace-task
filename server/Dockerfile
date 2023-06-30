FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go mod init server

EXPOSE 8080
CMD go mod tidy
