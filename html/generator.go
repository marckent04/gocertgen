package html

import (
	"fmt"
	"html/template"
	"marckent04/certificate_generator/domain"
	"marckent04/certificate_generator/shared"
	"os"
	"time"
)

type Generator struct {
	templatePath, outputPath string
	certificate              domain.Certificate
}

func (g *Generator) generateTemplate() *template.Template {
	tpl := shared.ReadFile(g.templatePath)

	functions := template.FuncMap{
		"formatDate": func(time time.Time) string {
			return time.Format("02/01/2006")
		},
	}

	t, err := template.New("cert").Funcs(functions).Parse(tpl)
	shared.HandleError(err)

	return t
}

func (g *Generator) Save() {

	t := g.generateTemplate()

	path := fmt.Sprintf("%s/%s.html", g.outputPath, shared.GenerateFileName(g.certificate))
	file, err := os.Create(path)
	shared.HandleError(err)
	defer file.Close()

	err = t.Execute(file, g.certificate)
	shared.HandleError(err)
}

func NewGenerator(certificate domain.Certificate, templatePath, outputPath string) *Generator {
	return &Generator{
		templatePath: templatePath,
		outputPath:   outputPath,
		certificate:  certificate,
	}
}
