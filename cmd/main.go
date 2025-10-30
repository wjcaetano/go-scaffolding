package main

import (
	"flag"
	"fmt"
	"os"

	"go-scaffolding/internal/generator"
	"go-scaffolding/internal/parser"
)

func main() {
	recipePath := flag.String("recipe", "recipes/default.yaml", "Path to the recipe file")
	outputDir := flag.String("output", "output", "Path to the output directory")
	flag.Parse()

	recipe, err := parser.Parse(*recipePath)
	if err != nil {
		fmt.Printf("Error parsing recipe file: %v\n", err)
		os.Exit(1)
	}

	if err := generator.Generate(recipe, *outputDir); err != nil {
		fmt.Printf("Error generating project: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Project generated successfully in %s\n", *outputDir)
}
