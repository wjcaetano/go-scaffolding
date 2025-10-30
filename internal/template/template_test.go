package template

import (
	"go-scaffolding/internal/parser"
	"testing"
)

func TestProcess(t *testing.T) {
	data := struct {
		Name string
	}{
		Name: "test-project",
	}

	result, err := Process("Hello, {{.Name}}!", data)
	if err != nil {
		t.Fatalf("Process() error = %v", err)
	}

	if result != "Hello, test-project!" {
		t.Errorf("Expected 'Hello, test-project!', got '%s'", result)
	}
}

func TestProcessRecipe(t *testing.T) {
	recipe := &parser.Recipe{
		Name: "test-project",
		Files: []parser.File{
			{
				Path:    "README.md",
				Content: "This is a {{.Name}}",
			},
		},
	}

	processedRecipe, err := ProcessRecipe(recipe)
	if err != nil {
		t.Fatalf("ProcessRecipe() error = %v", err)
	}

	expectedContent := "This is a test-project"
	if processedRecipe.Files[0].Content != expectedContent {
		t.Errorf("Expected content to be '%s', got '%s'", expectedContent, processedRecipe.Files[0].Content)
	}
}
