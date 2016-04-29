FROM golang:1.6

WORKDIR /workspace
ADD . /workspace

# {{{ このへんはちょっといろいろあって...
RUN echo "nameserver 8.8.8.8" >> /etc/resolv.conf
RUN go get golang.org/x/text/encoding
RUN go get golang.org/x/crypto/ssh/terminal
RUN go get google.golang.org/appengine
RUN go get google.golang.org/cloud/storage
RUN go get github.com/golang/glog
RUN go get github.com/otiai10/marmoset
# }}}

EXPOSE 8080
ENTRYPOINT go run main.go
