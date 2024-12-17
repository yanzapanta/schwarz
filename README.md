# Schwarz Tech Challenge

A Golang service that retrieves, creates and applies coupon.

## Technologies & Packages Used:
- Go version 1.23.4
- Docker
- MySQL Docker Image: mysql:8.0
- Gin Web Framework
- Logrus for Logging

```
├── cmd
│   ├── coupon_service
│       └── main.go
├── internal
|   |── api  
│       ├── entity
│           ├── application_request.go
│           ├── coupon_request.go
│           ├── coupons.go
|       |── api.go
|       |── coupon.go
|   |── config
│       ├── config.go
│       |── db.go
│   ├── repository
│       ├── memdb
│           ├── memdb.go
│   ├── service
│       ├── entity
│           ├── basket.go
│           ├── coupon.go
│       ├── validate
│           ├── coupon.go
│       ├── service_test.go
│       ├── service.go
├── resources
│       ├── db.sql
├── .env
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── makefile
```

## Installation
1. Clone ```schwarz``` repository
2. Run ```go mod tidy```
3. Run ```make up```
4. Localhost: ```http://localhost/```

## Endpoints
- [GET] /api/coupons - retrieves coupons by code
- [POST] /api/create - creates a coupon
- [POST] /api/apply - applies a coupon to a basket

### Note: 
- Database ```coupons``` will be automatically created
- Refer to the ```makefile``` to see more commands

