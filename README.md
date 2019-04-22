# Required


    Mysql
    Redis
    MongoDB
    kafka

# Conf

    config/
    
    
# Run
1. env     `export GOOPS_WORK_DIR=<your project working dir>`
           
2. sql init           
     
     import `pkg/app/generators/admin.sql` to database


3. sql model migrate

via cobra

```go
    ./cli/ops-cli migrate  // 数据库初始化
```

4. run server

```go
    go run main.go
```

5.run kafka

> docker-compose -f docker-compose.yml up -d run kafka&zoopkeeper service


# TODO LIST

- [x] permission 
- [x] grpc
- [x] message queue
- [x] cron tasks
- [x] mongodb
- [x] common api
- [x] command line tool
- [x] crawler
- [ ] error email
