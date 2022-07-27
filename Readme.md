### How To local RUN

You can using **launch.json** in your visual studio code

Change `system.env` with your environment

1. use `go mod tidy`
2. use `go mod download`
3. use `go mod vendor`
4. use command `go run main.go`

### Migration
Migration with `makefile` :
``` 
make run-migrate
```

### Documentation API

You can access the swagger

```
{{url}}/docs/index.html
```

In other documentation in **documentation** directory, and you can use the postman
`postman_collection.json`

#### by **Ahmad Rifai**
