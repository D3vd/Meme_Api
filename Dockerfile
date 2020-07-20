# Build Stage
FROM golang:alpine as build

WORKDIR $GOPATH/src/github.com/R3l3ntl3ss/Meme_Api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/Meme_API

# Final Stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/Meme_API /app/

EXPOSE 8080

CMD ./Meme_API

