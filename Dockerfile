FROM golang:1.11-alpine
RUN mkdir /app
COPY httpserve.go /app
WORKDIR /app
RUN go build -o httpserve .
EXPOSE 8080 8888
USER 1001
ENTRYPOINT ["/app/httpserve"]