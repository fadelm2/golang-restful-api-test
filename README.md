# golang-Restful-api-Test

## Description

this is golang Restful API test.

## Architecture


![Clean Architecture](architecture.png)

1. External system perform request (http, gRPC, Messaging, etc)
2. The Delivery creates various Model from request data
3. The Delivery calls Use Case, execute it using model data
4. The Use Case create Entity data for the business login
5. The Use Case calls Repository, and execute it using Entity data
6. The Repository Use Entity data to perform databse operation
7. The Repository perform database operation to the database
8. The Use Case create various Model for Gateway it using Model data
9. The Use Case calls gateway, and execute it using Model data
10. The Gateway using Model data to construct it using Model data
11. The Gateway perform request to external system (HTTP, gRPC, Messaging, etc)

## Tech Stack

- Golang : https://github.com/golang/go
- MySQL (Database) : https://github.com/mysql/mysql-server
- Apache Kafka : https://github.com/apache/kafka

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus
- Confluent Kafka Golang : https://github.com/confluentinc/confluent-kafka-go

## Configuration

All configuration is in `config.json` file.

## API Spec

All API Spec is in `api` folder.

## Database Migration

All database migration is in `db/migrations` folder
.
### Installation Container Kafka and Mysql Server with docker 
```Shell
docker compose up
```

```shell
docker compose -f docker.kafka.yml up -d
```
# Create Database

#Docker exec container mysql
```Shell 
docker exec -it mysql8 bash -l
```

```Shell 
mysql -u root -p root
```

#create database
```Shell 
create database golang_test;
```


### Create Migration

```Shell
migrate create -ext sql -dir db/migrations create_table_users
```

```Shell
migrate create -ext sql -dir db/migrations create_table_creditcard
```


### Run Migration

```shell
migrate -database "mysql://root:root@tcp(localhost:3310)/golang_test?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up
```

### Run Appication


### Run unit test


```bash
go test -v ./test/
```


### Run Web Server

```
go run cmd/web/main.go
```


### Run worker

```
go run cmd/worker/main.go
```
