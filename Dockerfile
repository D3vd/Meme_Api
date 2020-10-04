# Build Stage
FROM golang:alpine as build

WORKDIR /src/Meme_Api

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

