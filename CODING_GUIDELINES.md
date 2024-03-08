# Coding Guidelines

## Code

We want our code to be resistant to the passage of time and also to its growth. With the API already stable, the code will be read much more frequently than it is written, which is why we prefer the code to be tedious to write but easy to read.

In this context, the objective of these guidelines is to manage the complexity of our code. Below you will find a list of tools and guides to help you in that regard:

- Follow (as much as you can) the [golang code review guide](https://github.com/golang/go/wiki/CodeReviewComments).
- Follow (as much as you can) the [effective go guide](https://golang.org/doc/effective_go).
- Prepare your [environment to activate our hooks](#environment) that validate code and tests before each commit.
- All new code must contain automated tests that validate the new rules.

The important thing is that you are aware when you make a contribution, and try to follow the same style that is already used in the code. If you have any questions, do not hesitate to ask us!

## Development Tips and Environment Tools

### App Specs

This app is written in GoLang 1.19. In order to contribute to this project, you must provide tested code to keep our coverage as high as possible.

### Environment

To prepare your environment, we recommend running the command below, as it adds our `pre-commit hook` to ensure some rules before pushing code. To include it you must run just once (*or when the script is updated*):

```sh
make hooks
```

### Editor Configuration

To be able to run/debug this application on VSCode, here is a example of a `launch.json`:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceRoot}/cmd/api/",
      "env": {},
      "envFile": [
        "${workspaceFolder}/variables.env"
      ],
      "args": []
    }
  ]
}
```

#### Required environments

Some variables are necessary to run application:

```sh
export APP_PATH="/Users/username/go/src/github.com/<new-project-name>"
export SCOPE="local"
export checksumEnabled="false"
```

### Dependency Management

This app manages its dependencies using `go mod`, in this topic there are some instructions on how to prepare your
environment to have it working.

#### How to set up project to use go modules

When developing using Go modules, it is imperative that you add the following `env var` to your local environment 
(usually in your `~/.profile` or equivalent):

```sh
export GOPRIVATE="github.com/<new-project-name>"
```

If using Go 1.13+ you can run once:

```sh
go env -w GOPRIVATE="github.com/<new-project-name>"
```

Once you've set `GOPRIVATE`, Go modules resolver will use git to clone repositories that match with the value in the env
variable. Go modules clones using https by default, in order to authenticate with Github and be able to clone private 
repositories you usually need to use ssh. In order to force git to use `ssh` on Github, you need to set the following 
lines in your `~/.gitconfig`:

```sh
[url "ssh://git@github.com/"]
    insteadOf = https://github.com/
```

#### Links

* [Using Go Modules](https://blog.golang.org/using-go-modules).
* [Migrating to Go Modules](https://blog.golang.org/migrating-to-go-modules).

### Building, Running and Testing

This repository has some make commands that helps developers to get started to it. To used you MUST have `docker` and 
`docker-compose` installed locally.

To get all support services running, just run `make dev` and it will run all support service containers.

To run the tests, you just need to run `make test`.

To execute the app locally in docker, the command is `make run`

When executing inside a container, the application is configured to perform the migrations in 
`/migrations/mysql/testlocal` automatically at startup.

To execute the process outside the application container, it is necessary that the environment variables are properly 
configured. `cmd/migration/main.go` is a file that performs the migration process, and can be called with the 
`UP | DOWN` parameters.

It can be executed directly using golang:
Eg:

```sh
go run cmd/migration/main.go UP
```

To populate local database with valid data for testing purposes, you can run `make fixtures`. `make fixtures` will load 
the fixtures and generate test data for flight, forecasts and recommendations for next week.

The fixture files are located in `/resources/fixtures` folder. The `cmd/fixtures/app` is responsible for starting the 
application modules invoking the data generation methods in order.

### Running the App Locally with Debugging

To run the application locally with debugging capabilities, follow these steps:

1. Go to `Run > Edit Configurations...` and select `+` to add a new configuration.
2. Select `Go Build` and name it `Build Name of the APP`.
3. In `Environment` variables, add the following:
`APP_PATH=$PWD;configFileName=configFileName=resources/config/local.properties`
4. Click on `EnvFile` tab and add select `+` to add new envfile
5. Choose `variables.env` on project root folder
6. Click on `Apply` and `OK`.
7. Run the APP by clicking the green play button in the top right corner.
8. You should now be able to access the APP at `http://localhost:8080`.

For more detailed instructions and additional configurations, refer to the [Goland guide](docs/guide/dev-env/goland.md).

### How to configure golangci-lint

[Install](https://golangci-lint.run/usage/install/#macos)

[Editor Integration](https://golangci-lint.run/usage/integrations/#editor-integration)

> run  `make lint` to execute linter locally