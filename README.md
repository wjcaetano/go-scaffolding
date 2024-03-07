# Go Project Model Documentation

This document describes the standard structure of a Go project, which has been designed to be easy to understand and maintain, while following the best practices of package-oriented design in Go.

## Project Structure

Here is the directory structure of the project:

```
.
├── cmd
│   └── api
│       ├── main.go
│       └── modules
├── app
│   ├── $domain_name*
│   │   ├── entity.go
│   │   ├── repository
│   │   │   └── $technology
│   │   ├── service
│   │   │   └── $context
│   │   └── usecase
│   │       └── $usecase
│   └── $another_domain_name*
├── test
│   ├── integration
│   └── mock
├── internal
│   ├── config
│   ├── fury
│   ├── httpclient
│   ├── metrics
│   ├── publisher
│   └── server
├── pkg
│   ├── apperror
│   ├── customdate
│   ├── http
│   └── validator
├── resources
│   └── config
├── Dockerfile
├── Dockerfile.runtime
├── Makefile
├── README.md
├── go.mod
└── go.sum
```

## Directory Description

- `cmd`: Contains the executable code, which can be a web application, a CLI, etc.
- `app`: Contains the application code and should be organized into domain packages.
- `test`: Contains the test suites, helpers and also the integration tests for the application.
- `internal`: Contains all the infrastructure code, separated as packages for generic functionalities.
- `pkg`: Contains code that is intended to be used by other projects.
- `resources`: Contains configuration files and other resources necessary for the project.

## Description of Domain Packages (`$domain_name*`)

Each domain should be self-contained and represent something that makes sense for the domain and purpose of the application. Within each domain package, we have:

- `entity.go`: Defines the domain entities.
- `repository`: Responsible for retrieving and recording data, no matter the source of these data.
- `service`: Responsible for implementing the application logic that may or may not be based on data from the repositories.
- `usecase`: Aims to remove the complexity and coupling of the service layer, building features that can apply logic by orchestrating two or more services.

## Conclusion

This project structure is just a suggestion and may not be suitable for all projects. However, it provides a good starting point for creating Go projects that are easy to understand and maintain.
