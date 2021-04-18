# Golang Boilerplate Project using Echo Framework
The fastest way to build a restful api with [Echo](https://github.com/labstack/echo/) with a structured project that using **PostgreSQL** db and **JWT** base authentication middleware.

This project ships following features as default
* ORM Integration
* Easy Database Migration
* Authentication Using Jwt
* Easy dotenv Management
* Easy to Mock all Interfaces
* CORS Configuration

 Inspired from [Gin boilerplate](https://github.com/Massad/gin-boilerplate).
 
[![License](https://img.shields.io/github/license/triaton/go-echo-boilerplate)](https://github.com/triaton/go-echo-boilerplate/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/triaton/go-echo-boilerplate)](https://github.com/triaton/go-echo-boilerplate/blob/master/go.mod)
[![DB Version](https://img.shields.io/badge/DB-PostgreSQL--latest-blue)](https://github.com/triaton/go-echo-boilerplate/blob/master/go.mod)
[![Build Status](https://travis-ci.com/triaton/go-echo-boilerplate.svg?branch=master)](https://travis-ci.org/triaton/go-echo-boilerplate) [![Go Report Card](https://goreportcard.com/badge/github.com/triaton/go-echo-boilerplate)](https://goreportcard.com/report/github.com/triaton/go-echo-boilerplate)

## Configured with
- [gorm](https://github.com/jinzhu/gorm): ORM library for Golang
- [jwt-go](https://github.com/dgrijalva/jwt-go): JSON Web Tokens (JWT) as middleware
- [mockery](https://github.com/vektra/mockery): Go test mock library
- [godotenv](https://github.com/joho/godotenv): go dotenv library
- Go Modules
- Built-in **CORS Middleware**
- Built-in **RequestID Middleware**
- Feature **PostgreSQL 12**
- Environment support
- Unit test

### Installation

```
$ go get github.com/triaton/go-echo-boilerplate
```

```
$ cd $GOPATH/src/github.com/triaton/go-echo-boilerplate
```

```
$ go mod init
```

```
$ go install
```

## Running Your Application

Rename .env.example to .env and place your database credentials and jwt secret key

```
$ mv .env.example .env
$ go run main.go
```

## Building Your Application

```
$ go build -v
```

```
$ ./go-echo-boilerplate
```

## Testing Your Application

```
$ go test -v ./...
```

### Generate Code Coverage
```
$ chmod +x generate-test-coverage.sh
$ ./generate-test-coverage.sh
```
This will generate `cover.html` with detailed coverage result.

### Make mock interfaces with mockery
```
$ mockery --all --keep-tree
```
For more information please visit this [link](https://github.com/vektra/mockery).

### Add db model, controller and service
To add a controller
* Add `*_controller.go` file
* Define controller struct, define route entries, define route handlers and request dto interfaces. Please refer `auth/auth_controller.go`.<br>
* Add controller object to `routes/api.go`
```
...
func DefineApiRoute(e *echo.Echo) {
	controllers := []common.Controller{
		auth.AuthController{},
		blogs.BlogsController{},
		// add newly added controller here 
	}
	var routes []common.Route
...
```
To add a service
* Add `*_service.go` file.
* Define service interface, define service methods as well as singleton methods.
* Please don't forget to add `set` method. It will be used when mocking service methods.

To add db models
* add directory `models` and define `*.model.go` file.
* Define model struct and include `models.Base`. Please refer `blogs/models/blog.model.go`.

### Add migration
Edit `database/migrations.go` to add a new migration. For detailed guide, please visit this [link](http://gorm.io/docs/migration.html).

## Import Postman Collection (API's)

Download [Postman](https://www.getpostman.com/) -> Import -> Import From Link
https://www.getpostman.com/collections/401dc48dd6e9b15cc287

## Contribution

You are welcome to contribute to keep it up to date and always improving!

## License

(The MIT License)

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
