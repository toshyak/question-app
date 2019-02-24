FROM golang:1.11.5-alpine as builder
WORKDIR /go/src/github.com/toshyak/question-app/
COPY . /go/src/github.com/toshyak/question-app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o question-app .

FROM alpine:latest
WORKDIR /go/
COPY --from=builder /go/src/github.com/toshyak/question-app/question-app .
EXPOSE 8080
CMD ["./question-app"]