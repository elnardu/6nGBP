FROM golang:alpine

RUN apk upgrade -U && apk add glide git && rm -rf /var/cache/*

WORKDIR /go/src/github.com/elnardu/6nGBP/backend/
ADD glide.yaml .
ADD glide.lock .
RUN glide install && apk del glide git

ARG secret
ARG dbconnect
ENV SECRET $secret
ENV DBCONNECT $dbconnect

ADD . .
RUN go build -o myapp; mkdir /app/; cp myapp /app/
RUN rm -rf /go/src/github.com/elnardu/*

EXPOSE 8080
CMD ["/app/myapp"]
