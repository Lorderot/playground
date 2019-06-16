FROM golang:1.11
WORKDIR /go/src/github.com/Lorderot/playground/
COPY . .
RUN go install -v github.com/Lorderot/playground/src/server

CMD ["server"]