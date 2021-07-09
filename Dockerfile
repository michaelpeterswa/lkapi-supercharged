FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get -d -v .
RUN go build .

EXPOSE 8091
CMD ./lkapi-supercharged