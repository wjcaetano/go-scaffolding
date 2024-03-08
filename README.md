# Go Project Model Documentation

The `go-scaffolding` project is a template for future `Golang` projects. Its main objective is to standardize and accelerate the development of Go applications by providing a well-structured and easy-to-follow project layout.

## Project Structure

The project structure is as follows:

```markdown
.
├── CODING_GUIDELINES.md
├── CONTRIBUTING.md
├── Dockerfile
├── Dockerfile.runtime
├── Makefile
├── README.md
├── app
│   └── domain
│       ├── entity.go
│       ├── entrypoint
│       │   └── rest
│       │       └── search
│       │           └── dto.go
│       ├── repository
│       │   ├── http
│       │   │   └── dto.go
│       │   └── sql
│       │       └── dto.go
│       ├── service
│       │   └── rules
│       │       ├── servicename.go
│       │       └── servicename_test.go
│       └── usecase
│           └── rules
│               ├── usecasename.go
│               └── usecasename_test.go
├── arch-go.yml
├── cmd
│   ├── api
│   │   ├── main.go
│   │   └── modules
│   │       └── domain.go
│   └── modules
├── commands
│   ├── Dockerfile.directory-validator
│   ├── create_sql_user.sh
│   ├── directory-validator.sh
│   ├── execute_before_test.sh
│   ├── git
│   │   ├── pre-commit
│   │   └── pre-push
│   ├── linter.sh
│   ├── run_local.sh
│   └── run_test.sh
├── docker-compose.yaml
├── docs
│   ├── guide
│   └── specs
├── go.mod
├── internal
│   ├── config
│   ├── db
│   ├── httpclient
│   ├── metrics
│   ├── publisher
│   └── server
├── pkg
│   ├── apperror
│   ├── http
│   │   ├── errorhandler
│   │   ├── parser
│   │   └── render
│   └── validator
├── resources
│   └── config
│       └── local.properties
├── test
│   └── integration
└── variables.env
```

## Directory Description

- `app/domain`: This directory contains the business logic of the application. It is divided into several subdirectories:
    - `entity.go`: This file contains the definition of the entities used in the application.
    - `entrypoint`: This directory contains the HTTP handlers of the application.
      - `rest`: This directory contains the HTTP handlers of the application.
        - `search`: This directory contains the HTTP handlers of the application.
          - `dto.go`: This file contains the definition of the DTOs used in the application.
    - `repository`: This directory contains the interfaces of the repositories used in the application.
      - `http`: This directory contains the HTTP repositories used in the application.
        - `dto.go`: This file contains the definition of the DTOs used in the application.  
        - `sql`: This directory contains the SQL repositories used in the application.
          - `dto.go`: This file contains the definition of the DTOs used in the application.
    - `service`: This directory contains the services used in the application.
      - `rules`: This directory contains the rules used in the application.
        - `servicename.go`: This file contains the definition of the service used in the application.
        - `servicename_test.go`: This file contains the tests of the service used in the application.
    - `usecase`: This directory contains the use cases used in the application.
      - `rules`: This directory contains the rules used in the application.
          - `usecasename.go`: This file contains the definition of the use case used in the application.
          - `usecasename_test.go`: This file contains the tests of the use case used in the application.
- `cmd`: Contains the executable code, which can be a web application, a CLI, etc.
- `commands`: Contains various scripts for tasks such as linting, testing, and setting up the development environment.
  docs: Contains the project documentation.
- `internal`: Contains all the infrastructure code, separated as packages for generic functionalities. Each package has the following subdirectories:
    - `config`: Contains the application configuration.
    - `db`: Contains the database configuration and connection.
    - `httpclient`: Contains the HTTP client configuration and connection.
    - `metrics`: Contains the metrics configuration and connection.
    - `publisher`: Contains the publisher configuration and connection.
    - `server`: Contains the HTTP server configuration and connection.
- `pkg`: Contains all the generic code, separated as packages for generic functionalities. Each package has the following subdirectories:
    - `apperror`: Contains the application error definition.
    - `http`: Contains the HTTP generic code.
    - `validator`: Contains the validator generic code.
- `resources`: Contains the application resources, such as configuration files.
- `test`: Contains the application tests.
- `variables.env`: Contains the application environment variables.
- `CODING_GUIDELINES.md`: Contains the coding guidelines for the application.
- `CONTRIBUTING.md`: Contains the contributing guidelines for the application.
- `Dockerfile`: Contains the Dockerfile for the application.
- `Dockerfile.runtime`: Contains the Dockerfile for the application runtime.
- `Makefile`: Contains the Makefile for the application.
- `README.md`: Contains the README for the application.
- `arch-go.yml`: Contains the Arch-go configuration for the application.
- `go.mod`: Contains the Go module for the application.
- `docker-compose.yaml`: Contains the Docker Compose configuration for the application.
- `variables.env`: Contains the application environment variables.

## How to use

To use this project as a template, you can do the following:

1. Clone the repository:
```bash 
git clone
```
2. Change the remote URL to the new repository:
```bash
git remote set-url origin <new-repository-url>
```

3. Change the project name in the `go.mod` file:
``` 
module <new-project-name>
```
4. Change the project name in the `arch-go.yml` file:
```yaml
name: <new-project-name>
```
5. Change the project name in the `Makefile` file:
```makefile
PROJECT_NAME=<new-project-name>
```
6. Change the project name in the `Dockerfile` file:
```dockerfile
LABEL org.opencontainers.image.title="<new-project-name>"
```
7. Change the project name in the `Dockerfile.runtime` file:
```dockerfile
LABEL org.opencontainers.image.title="<new-project-name>"
```
8. Change the project name in the `README.md` file:
```markdown
# <new-project-name> Documentation
```
9. Change the project name in the `CODING_GUIDELINES.md` file:
```markdown
# <new-project-name> Coding Guidelines
```
10. Change the project name in the `CONTRIBUTING.md` file:
```markdown
# <new-project-name> Contributing Guidelines
```
11. Change the project name in the `docker-compose.yaml` file:
```yaml
<new-project-name>:
```
12. Change the project name in the `variables.env` file:
```env
PROJECT_NAME=<new-project-name>
```
