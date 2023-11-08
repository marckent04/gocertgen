package pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"marckent04/certificate_generator/domain"
	"marckent04/certificate_generator/shared"
	"os"
	"path"
)

type Generator struct {
	pdf                   gopdf.GoPdf
	cert                  domain.Certificate
	pageH, pageW          float64
	packageDir, outputDir string
}

func (g *Generator) Save() {
	g.addPdfHeader()
	g.addPdfFooter()
	g.pdf.AddPage()
	g.addPdfBackground()
	g.addPdfBody()

	fileName := fmt.Sprintf("%s.pdf", shared.GenerateFileName(g.cert))

	err := g.pdf.WritePdf(path.Join(g.outputDir, fileName))
	shared.HandleError(err)
}

func NewGenerator(certificate domain.Certificate, outputDir string) *Generator {
	currentDir, err := os.Getwd()
	shared.HandleError(err)

	generator := &Generator{}

	generator.packageDir = path.Join(currentDir, "pdf")
	generator.outputDir = outputDir
	generator.initPdf()
	generator.cert = certificate

	return generator
}
