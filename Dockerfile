FROM golang:1.15-buster

WORKDIR /opt/mehlj-pipeline/api/

COPY api/* /opt/mehlj-pipeline/api/

RUN go get -u github.com/gorilla/mux &&\
    go get -u github.com/mattn/go-sqlite3 &&\
    go build main.go sql.go

CMD ["./main"]
