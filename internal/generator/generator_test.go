package generator

import (
	"go-scaffolding/internal/parser"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
	recipe := &parser.Recipe{
		Directories: []string{"test-dir"},
		Files: []parser.File{
			{
				Path:    "test-file.txt",
				Content: "test-content",
			},
		},
	}

	outputDir, err := os.MkdirTemp("", "test-output")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(outputDir)

	if err := Generate(recipe, outputDir); err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	// Check if the directory was created
	dirPath := filepath.Join(outputDir, "test-dir")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		t.Errorf("Expected directory '%s' to be created", dirPath)
	}

	// Check if the file was created
	filePath := filepath.Join(outputDir, "test-file.txt")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Expected file '%s' to be created", filePath)
	}

	// Check the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "test-content" {
		t.Errorf("Expected file content to be 'test-content', got '%s'", string(content))
	}
}
