FROM golang:1.5
EXPOSE 8080
COPY main.go .
RUN go get github.com/otiai10/marmoset
ENTRYPOINT go run main.go
