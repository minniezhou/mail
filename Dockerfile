FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o mail ./cmd/api

EXPOSE 54321
ENV WEB_PORT_DEFAULT = 54321

CMD ["./mail"]
