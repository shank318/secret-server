## Secret Server

The secret server can be used to store and share secrets using the random generated URL. But the secret can be read only a limited number of times after that it will expire and won’t be available. The secret may have TTL. After the expiration time the secret won’t be available anymore. 

## Requirements

For building and running the application you need:

- [Go 1.11](https://golang.org/doc/go1.11)
- [Mysql]

## Running the application locally

```shell
go build
go run main.go
```

## API Docs

Swagger --> https://secretserverapi.herokuapp.com/swagger/index.html#/

The app is deployed on Heroku which is considerably slow most of the time. Please run the code locally to check the api's. Also, the mysql database I am using is a free tier Heroku database which allows only 10 connections at a time. 
