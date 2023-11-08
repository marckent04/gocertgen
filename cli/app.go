package cli

import (
	"flag"
	"fmt"
	"log"
)

type Params struct {
	OutputDir, Format, CsvPath, TemplatePath string
}

func GetParams() Params {
	format, csv, output, htmlTemplate :=
		flag.String("format", "html", "output file format"),
		flag.String("path", "", "csv file path"),
		flag.String("output", "outputs", "output directory"),
		flag.String("template", "", "html template path")

	flag.Parse()

	params := Params{
		CsvPath:      *csv,
		OutputDir:    *output,
		Format:       *format,
		TemplatePath: *htmlTemplate,
	}

	validateParams(params)

	return params
}

func validateParams(params Params) {
	var errors []string

	if params.Format != "pdf" && params.Format != "html" {
		errors = append(errors, fmt.Sprintf("%s format is not supported", params.Format))
	}

	if params.CsvPath == "" {
		errors = append(errors, "csv path is required")
	}

	if params.Format == "html" && params.TemplatePath == "" {
		errors = append(errors, "HTML template path is required")
	}

	if len(errors) != 0 {
		log.Fatal(errors)
	}

}
