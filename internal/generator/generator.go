package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"go-scaffolding/internal/parser"
	"go-scaffolding/internal/template"
)

func Generate(recipe *parser.Recipe, outputDir string) error {
	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create directories
	for _, dir := range recipe.Directories {
		dirPath := filepath.Join(outputDir, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
		}
	}

	// Process and create files
	processedRecipe, err := template.ProcessRecipe(recipe)
	if err != nil {
		return fmt.Errorf("failed to process recipe templates: %w", err)
	}

	for _, file := range processedRecipe.Files {
		filePath := filepath.Join(outputDir, file.Path)
		fileDir := filepath.Dir(filePath)

		// Create parent directory if it doesn't exist
		if err := os.MkdirAll(fileDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory for file %s: %w", filePath, err)
		}

		if err := os.WriteFile(filePath, []byte(file.Content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filePath, err)
		}
	}

	return nil
}
