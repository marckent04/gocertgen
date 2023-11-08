package main

import (
	"marckent04/certificate_generator/cli"
	"marckent04/certificate_generator/csv"
	"marckent04/certificate_generator/html"
	"marckent04/certificate_generator/pdf"
	"marckent04/certificate_generator/shared"
)

func main() {

	params := cli.GetParams()

	csvParser := csv.NewParser(params.CsvPath)

	shared.CreateOutputFolderIfNotExists(params.OutputDir)

	for _, certificate := range csvParser.Certificates {
		var saver Saver

		if params.Format == "html" {
			saver = html.NewGenerator(certificate, params.TemplatePath, params.OutputDir)
		} else if params.Format == "pdf" {
			saver = pdf.NewGenerator(certificate, params.OutputDir)
		}

		saver.Save()

	}
}

type Saver interface {
	Save()
}
