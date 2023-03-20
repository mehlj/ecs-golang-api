# mehlj-pipeline

Golang REST API deployed to AWS

## API Interaction

### Create
```
$ curl -d '{"Name":"ball","quantity":4}' -X POST localhost/product
```

### Read
```
$ curl localhost:3000/products
```

### Update
```
$ curl -d '{"Name":"ball","quantity":5}' -X PUT localhost/product
```

### Delete
```
$ curl -d '{"Name":"ball","quantity":5}' -X DELETE localhost/product
```

#### Go dependencies
```
$ go mod init mehlj-pipeline
$ go mod tidy
```