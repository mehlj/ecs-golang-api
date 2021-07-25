FROM golang:1.15-buster

WORKDIR /opt/mehlj-pipeline/api/

COPY api/* /opt/mehlj-pipeline/api/

RUN go get -u github.com/gorilla/mux &&\
    go build main.go

CMD ["./main"]
