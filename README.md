# Required


    Mysql
    Redis
    MongoDB
    kafka

# Conf

    config/
    
    
# Run
1. env     `export GOOPS_WORK_DIR=<your project working dir>`
1. import pkg/app/generators/admin.sql

```sql
    source admin.sql
```

2. run migrate

```go
    go run migrate.go  // 数据库初始化
    go run init_permission.go  // casbin权限数据库初始化
```

3. run server
```go
    go run main.go
```

4.run kafka
> docker-compose -f docker-compose.yml up -d run kafka&zoopkeeper service


# TODO LIST

- [x] permission 
- [x] grpc
- [x] message queue
- [x] cron tasks
- [x] mongodb
- [x] common api