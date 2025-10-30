package parser

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	// Create a temporary recipe file for testing
	content := `
name: test-project
version: 1.0.0
author: Test Author
description: A test project
directories:
  - cmd
  - internal
files:
  - path: cmd/main.go
    content: package main
`
	tmpfile, err := os.CreateTemp("", "recipe.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test parsing the recipe file
	recipe, err := Parse(tmpfile.Name())
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Check the parsed recipe
	if recipe.Name != "test-project" {
		t.Errorf("Expected name to be 'test-project', got '%s'", recipe.Name)
	}
	if recipe.Author != "Test Author" {
		t.Errorf("Expected author to be 'Test Author', got '%s'", recipe.Author)
	}
	if len(recipe.Directories) != 2 {
		t.Errorf("Expected 2 directories, got %d", len(recipe.Directories))
	}
	if len(recipe.Files) != 1 {
		t.Errorf("Expected 1 file, got %d", len(recipe.Files))
	}
}
