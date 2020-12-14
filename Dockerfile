FROM golang:1.15 AS build
WORKDIR go/src
COPY . ./
COPY main.go .
RUN mkdir -p /go/src/github.com/books/
COPY ./go /go/src/github.com/books/go
ENV CGO_ENABLED=0
ENV GOARCH=amd64
#ENV GOOS=linux
RUN go get -d -v ./...

RUN go build -o books

FROM scratch AS runtime
COPY --from=build /go/go/src/books ./
EXPOSE 8080/tcp
ENTRYPOINT ["./books"]
