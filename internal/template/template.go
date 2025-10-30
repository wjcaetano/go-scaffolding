package template

import (
	"bytes"
	"text/template"
	"go-scaffolding/internal/parser"
)

func Process(t string, data interface{}) (string, error) {
	tmpl, err := template.New("").Parse(t)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func ProcessRecipe(recipe *parser.Recipe) (*parser.Recipe, error) {
	for i, file := range recipe.Files {
		processedContent, err := Process(file.Content, recipe)
		if err != nil {
			return nil, err
		}
		recipe.Files[i].Content = processedContent
	}

	return recipe, nil
}
