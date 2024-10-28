# echo-clean-arch

## Description

This is an example of implementation of Clean Architecture in Go (Golang) projects.

Rule of Clean Architecture by Uncle Bob

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This project has 4 Domain layer :

- Models Layer
- Repository Layer
- Usecase Layer
- Delivery Layer

#### The diagram:

![golang clean architecture](https://github.com/hoaibn-0332/echo-clean/blob/main/echo-clean.jpg)

The original explanation about this project's structure can read from this medium's post : https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047.
It may be different already, but the concept still the same in application level, also you can see the change log from v1 to current version in Master.

### How To Run This Project

> Make Sure you have run the article.sql in your mysql

Since the project is already use Go Module, I recommend to put the source code in any folder but GOPATH.

#### Run the Testing

```bash
$ make tests
```

#### Run the Applications

Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

# Clone into your workspace
$ git clone https://github.com/hoaibn-0332/echo-clean.git

#move to project
$ cd echo-clean

# copy the example.env to .env
$ cp example.env .env

# Run the application
$ docker-compose up

# The hot reload will running

# Execute the call in another terminal
$ curl localhost:8080/articles
```

### API Documentation
[Swagger](https://hoaibn-0332.github.io/echo-clean/swagger/)

## Dependencies

Echo-Clean is built using the Go programming language and leverages a number of open-source libraries to provide its
core functionality. Some of the key dependencies include:

- [Echo](https://github.com/labstack/echo)
- [ORM](https://entgo.io/)
- [OpenAPI](https://github.com/deepmap/oapi-codegen)
- [Logger](https://github.com/rs/zerolog)