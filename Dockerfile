FROM golang:1.23-bullseye

WORKDIR /app
COPY . .

RUN go build -o /bin/simple_http_request_display ./cmd/simple_http_request_display

CMD ["/bin/simple_http_request_display"]
