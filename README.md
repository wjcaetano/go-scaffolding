# Go Scaffolding

`go-scaffolding` is a tool for generating Go project structures based on YAML recipes. It helps standardize and accelerate the development of Go applications by providing a well-structured and easy-to-follow project layout.

## Project Structure

The project structure is as follows:

```markdown
.
├── cmd
│   └── main.go
├── internal
│   ├── generator
│   ├── parser
│   └── template
├── recipes
│   └── default.yaml
├── web-api
│   ├── ... (existing web-api code)
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Directory Description

- `cmd`: Contains the executable code for the scaffolding tool.
- `internal`: Contains the internal logic of the scaffolding tool.
    - `generator`: Contains the logic for generating the project structure.
    - `parser`: Contains the logic for parsing recipe files.
    - `template`: Contains the logic for processing templates.
- `recipes`: Contains the YAML recipes for generating project structures.
- `web-api`: Contains the original web API application, now as a separate module.

## Usage

To generate a new project, run the following command:

```bash
go run cmd/main.go --recipe recipes/default.yaml
```

This will generate a new project in the `output` directory based on the `default.yaml` recipe.
