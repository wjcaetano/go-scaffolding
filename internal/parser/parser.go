package parser

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Recipe struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Author      string `yaml:"author"`
	Description string `yaml:"description"`
	Directories []string `yaml:"directories"`
	Files       []File `yaml:"files"`
}

type File struct {
	Path    string `yaml:"path"`
	Content string `yaml:"content"`
}

func Parse(path string) (*Recipe, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var recipe Recipe
	err = yaml.Unmarshal(data, &recipe)
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}
