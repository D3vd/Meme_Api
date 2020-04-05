FROM golang:1.13.5

RUN apt-get update -y

WORKDIR /go/src/github.com/R3l3ntl3ss/Meme_Api

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD [ "go", "run", "." ]
