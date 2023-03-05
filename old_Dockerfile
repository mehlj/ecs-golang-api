FROM golang:1.15-buster

WORKDIR /opt/mehlj-pipeline/api/

COPY api/* ./

RUN go mod download

RUN go build main.go sql.go

CMD ["./main"]