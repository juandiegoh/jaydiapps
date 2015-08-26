FROM golang
 
ADD . /go/src/github.com/juandiegoh/jaydiapps
RUN go get github.com/gorilla/mux
RUN go install github.com/juandiegoh/jaydiapps
ENTRYPOINT /go/bin/jaydiapps
 
EXPOSE 8080
