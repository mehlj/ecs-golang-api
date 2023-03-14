# mehlj-pipeline

Example project to showcase CI/CD pipelines.

## Quickstart
```
$ docker-compose up -d
```

## API Interaction

### Create
```
$ curl -d '{"Name":"ball","quantity":4}' -X POST localhost:3000/product
```

### Read
```
$ curl localhost:3000/products
```
```
$ curl localhost:3000/product/oranges
```

### Update
```
$ curl -d '{"Name":"oranges","quantity":4}' -X PUT localhost:3000/product
```

### Delete
```
$ curl -d '{"Name":"ball","quantity":4}' -X DELETE localhost:3000/product
```


#### Go dependencies
```
$ go mod init mehlj-pipeline
$ go mod tidy
```

