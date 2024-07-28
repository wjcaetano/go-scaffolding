.PHONY: build run test
all: help

help:
	@echo "<project-name> Makefile commands"
	@echo ""
	@echo "build                 Executes the build of all images required to run <project-name> API"
	@echo "down                  Stops all docker containers"
	@echo "run                   Start all containers and keep attached"
	@echo "run_local             Run application in the current terminal"
	@echo "dev                   Starts support services and expose (Use to run app and tests in your IDE)"
	@echo "lint                  Executes local golangci-lint for application"
	@echo "test                  Start containers and run Unit Tests"
	@echo "test_up               Starts all docker containers and detach"
	@echo "test_run              Run Unit Tests"
	@echo "hooks                 Install git hooks locally"
	@echo ""
	@echo "Documentation"
	@echo ""
	@echo "specs_generate        Generate a index.html with the content of our Swagger documentation"
	@echo "specs_serve           Creates a server to render the content of Swagger index.yaml"
	@echo ""
	@echo "guide_serve           Run a local server to render the content of guide documentation"
	@echo ""

build:
	@docker-compose build --no-cache --force-rm --pull

down:
	@docker-compose down

run: down
	@docker-compose up --build

run_local:
	@eval $$(egrep -v '^#' variables.env | xargs) APP_PATH=$$PWD go run cmd/api/*.go

fs:
	@make down

hooks:
	@if [ ! -d "commands/git/pre-commit" ]; then mkdir -p commands/git/pre-commit; fi
	@if [ ! -d "commands/git/pre-push" ]; then mkdir -p commands/git/pre-push; fi
	@rm -rf .git/hooks/*
	@cp -r commands/git/pre-commit/ .git/hooks
	@cp -r commands/git/pre-push/ .git/hooks
	@chmod -R +x .git/hooks/pre-commit
	@chmod -R +x .git/hooks/pre-push

test:
	@docker-compose build --force-rm <project-name>
	@docker-compose run --rm <project-name> /commands/run_test.sh

cleanup:
	@find . -type d -name mocks -exec rm -rf {} \;

mocks: cleanup
	@eval $$(egrep -v '^#' variables.env | xargs) APP_PATH=$$PWD go generate ./...

test_local:
	@eval $$(egrep -v '^#' variables.env | xargs) APP_PATH=$$PWD go test ./... -count 1 -tags=integration -cover -p=1

test_up: down
	@docker-compose up -d

test_run:
	@docker-compose exec <project-name> /commands/test.sh

.PHONY:
specs_generate:
	@echo "Generating OpenAPI documentation from code comments"
	@Swagger generate spec --scan-models --input=docs/specs/template.yaml --output=docs/specs/swagger.yaml
	@echo "Generated swagger.yaml. You may access the docs by running the specs_serve command"

.PHONY:
specs_serve:
	@echo "Remove docker docs api with name $(DOCKER_DOCS_API_CONTAINER)"
	@${DOCKER_EXEC} rm -f -v ${DOCKER_DOCS_API_CONTAINER} || true
	@echo "Running docker docs api with name $(DOCKER_DOCS_API_CONTAINER)"
	@${DOCKER_EXEC} run -d -p ${DOCKER_DOCS_API_CONTAINER_PORT}:8080 --name ${DOCKER_DOCS_API_CONTAINER} -e SWAGGER_JSON=/api/swagger.yaml -v ${PWD}/docs/specs/:/api/ swaggerapi/swagger-ui:v3.25.4
	@echo "Documentation Api can be viewed at http://localhost:$(DOCKER_DOCS_API_CONTAINER_PORT)"

.PHONY:
guide_serve:
	@echo "Remove docker docs guide with name $(DOCKER_DOCS_GUIDE_CONTAINER)"
	@${DOCKER_EXEC} rm -f -v ${DOCKER_DOCS_GUIDE_CONTAINER} || true
	@echo "Running docker docs guide with name $(DOCKER_DOCS_GUIDE_CONTAINER)"
	@${DOCKER_EXEC} run -p ${DOCKER_DOCS_GUIDE_CONTAINER_PORT}:3000 -d --name ${DOCKER_DOCS_GUIDE_CONTAINER} -v ${PWD}/docs/guide:/usr/local/docsify littlstar/docker-docsify:latest
	@echo "Documentation Guide can be viewed at http://localhost:$(DOCKER_DOCS_GUIDE_CONTAINER_PORT)"

.PHONY: lint
lint:
	@make format
	@make lint-code
	@make lint-arch
	@make lint-directories
	@eval $$(egrep -v '^#' variables.env | xargs) APP_PATH=$$PWD go list -e -compiled -test=true -export=false -deps=true -find=false ./... > /dev/null
	@eval $$(egrep -v '^#' variables.env | xargs) APP_PATH=$$PWD golangci-lint run --config=.code_quality/.golangci.yml --new-from-rev=HEAD~1 --fix


.PHONY:
lint-code:
	@echo "Executing golangci-lint"
	@type "golangci-lint" > /dev/null 2>&1 || echo 'Please install golangci-lint: https://github.com/golangci/golangci-lint#install'
	golangci-lint run --config=.golangci-project.yml ./...
	@echo "Executing godepth"
	@type "godepth" > /dev/null 2>&1 || echo 'Please install godepth with command: go install github.com/arthurgustin/godepth@latest'
	find . -name "*.go" -exec godepth -over=4 {} \;

lint-arch:
	@echo "Executing arch-go"
	@type "arch-go" > /dev/null 2>&1 || echo 'Please install arch-go: https://github.com/fdaines/arch-go'
	arch-go

lint-directories:
	@echo "Executing directory-validator"
	commands/directory-validator.sh

GOIMPORTS=goimports

format:
	@echo "Executing goimports"
	@find . -name "*.go" -exec $(GOIMPORTS) -w {} \;