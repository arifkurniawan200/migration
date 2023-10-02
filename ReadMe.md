
# Migration Database

This is a service that can be used to perform database migrations.
In this case, two distinct databases are used.


## Authors

- [@arifkurniawan200](https://github.com/arifkurniawan200)


## Environment Variables

**Install Goose:**
```bash
 go install github.com/pressly/goose/v3/cmd/goose@latest
```

**Run Postgres:**
```bash
 sudo systemctl start postgresql
```

**Check Postgres:**
```bash
 sudo systemctl status postgresql
```

**Change App.yaml Configuration:**
```bash
 To run this project, you will need to add the following environment variables to your app.yaml file in folder config

`des_db` : destination database

`src_db` : source database

Change the configuration according to your database 
```



## Run Locally

change app.yaml.example in folder config to app.yaml

install dependencies

```bash
  go mod tidy
```

running database migration

```bash
  go run main.go db:migrate up
```

running seeding database

```bash
  go run main.go db:seeding
```

reset database (delete database and existing data)

```bash
  go run main.go db:migrate reset
```

running api server

```bash
  go run main.go api
```




## API Reference

#### Get all data in table source

```http
  GET /source/
```


#### Get all data in table destination

```http
  GET /destination/
```


#### Update all data in table destination

```http
  PUT /destination/
```


## Tech Stack

**Database:** PostgresSQL

**Framework:** Echo golang

**Migration:** Goose