# To-Do REST API Example

A To-Do REST API example with Go + Gin + GORM + MySQL.

## Preparing
* Go 1.16+
* MySQL 8.x

## Setup Database

Create **todo** database:
```SQL
CREATE DATABASE todo CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

Create a user and grant privileges to the database:
```SQL
create user 'todo'@'%' identified by 'todo123';
grant all privileges on todo.* TO 'todo'@'%';
```

## Start Service
```bash
$ go run ./main.go
or
$ go build
$ ./go-todo-api # windows ./go-todo-api.exe
```