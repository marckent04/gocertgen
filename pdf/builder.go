package pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"marckent04/certificate_generator/shared"
	"path"
	"strings"
)

func (g *Generator) initPdf() {
	g.pdf = gopdf.GoPdf{}
	g.pageH = gopdf.PageSizeA4.W
	g.pageW = gopdf.PageSizeA4.H

	g.pdf.Start(gopdf.Config{PageSize: gopdf.Rect{
		W: g.pageW,
		H: g.pageH,
	}})

	g.addPdfFonts()

}

func (g *Generator) addPdfFonts() {

	fonts := []font{
		newFont("Regular", "Roboto-Regular.ttf"),
		newFont("Bold", "Roboto-Bold.ttf"),
		newFont("Italic", "Roboto-lightItalic.ttf"),
	}

	for _, currentFont := range fonts {
		err := g.pdf.AddTTFFont(currentFont.Name, path.Join(g.packageDir, "fonts", currentFont.FileName))
		shared.HandleError(err)
	}
}

func (g *Generator) addPdfHeader() {
	err := g.pdf.SetFont("Regular", "", 40)
	shared.HandleError(err)
	g.pdf.AddHeader(func() {
		g.pdf.SetY(g.pageH / 5)
		g.addCenteredText("Certificate of Completion")
	})
}

func (g *Generator) addPdfFooter() {
	imagePath := path.Join(g.packageDir, "images", "signature.png")

	imageSize := &gopdf.Rect{
		W: 80,
		H: 80,
	}

	x := g.pageW - (imageSize.W + 55)
	y := g.pageH - (imageSize.H + 55)

	g.pdf.AddFooter(func() {
		err := g.pdf.Image(imagePath, x, y, imageSize)
		shared.HandleError(err)
	})
}

func (g *Generator) addPdfBody() {
	y := g.pageH / 3

	lines := []textLine{
		newTextLine("This Certificate is Presented to", "Regular", 20),
		newTextLine(strings.ToUpper(g.cert.Participant), "Bold", 30),
		newTextLine(fmt.Sprintf("For participation in the %s", strings.ToUpper(g.cert.Course)), "Regular", 25),
		newTextLine(fmt.Sprintf("Date: %s", g.cert.DeliveredAt.Format("02/01/2006")), "Italic", 15),
	}

	for _, line := range lines {
		g.addTextLine(line, y)
		y += 50
	}
}

func (g *Generator) addPdfBackground() {
	imagePath := path.Join(g.packageDir, "images", "frame.png")

	imageSize := &gopdf.Rect{
		W: g.pageW,
		H: g.pageH,
	}

	err := g.pdf.Image(imagePath, 0, 0, imageSize)
	shared.HandleError(err)
}

type font struct {
	Name, FileName string
}

func newFont(name, fileName string) font {
	return font{
		FileName: fileName,
		Name:     name,
	}
}
