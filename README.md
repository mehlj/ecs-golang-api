# mehlj-pipeline

Golang REST API deployed to AWS

## API Interaction

### Create
```bash
curl -d '{"Name":"ball","quantity":4}' -X POST localhost/product
```

### Read
```bash
curl localhost/products
```
```bash
curl localhost/product?name=ball
```

### Update
```bash
curl -d '{"Name":"ball","quantity":5}' -X PUT localhost/product
```

### Delete
```bash
curl -d '{"Name":"ball","quantity":5}' -X DELETE localhost/product
```

#### Go dependencies
```bash
go mod init mehlj-pipeline
go mod tidy
```