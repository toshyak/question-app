FROM golang:1.11.5-alpine as builder
LABEL stage=intermediate
WORKDIR /go/src/github.com/toshyak/question-app/
ENV GO111MODULE=on
RUN apk add --no-cache git
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . /go/src/github.com/toshyak/question-app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o question-app .

FROM alpine:latest
WORKDIR /go/
COPY --from=builder /go/src/github.com/toshyak/question-app/question-app .
EXPOSE 8080
CMD ["./question-app"]