FROM golang:1.17.2-alpine AS build

WORKDIR /src/go-web
COPY . /src/go-web
RUN go build -o /bin/go-web

EXPOSE 9090
ENTRYPOINT ["/bin/go-web"]
