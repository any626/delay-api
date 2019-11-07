FROM golang:alpine

WORKDIR /app
COPY . /app

RUN go build

EXPOSE 8080

CMD ["./delayApi"]