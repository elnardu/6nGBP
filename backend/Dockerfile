FROM golang:alpine

RUN apk upgrade -U && apk add --no-cache glide git && rm -rf /var/cache/*

WORKDIR /go/src/github.com/elnardu/6nGBP/backend/
ADD . .
RUN glide install && apk del glide git

ARG secret
ENV SECRET $secret

RUN go build -o myapp; mkdir /app/; cp myapp /app/
RUN rm -rf /go/src/github.com/elnardu/*

EXPOSE 8080
CMD ["/app/myapp"]
