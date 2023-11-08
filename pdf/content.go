package pdf

import (
	"github.com/signintech/gopdf"
	"marckent04/certificate_generator/shared"
)

func (g *Generator) addTextLine(line textLine, y float64) {
	g.pdf.SetXY(0, y)
	err := g.pdf.SetFontWithStyle(line.FontFamily, gopdf.Regular, line.FontSize)
	shared.HandleError(err)
	g.addCenteredText(line.Text)
}

func (g *Generator) addCenteredText(text string) {
	err := g.pdf.CellWithOption(&gopdf.Rect{
		W: g.pageW,
		H: 10,
	},
		text,
		gopdf.CellOption{Align: gopdf.Center})
	shared.HandleError(err)
}

type textLine struct {
	Text, FontFamily string
	FontSize         float64
}

func newTextLine(text, font string, fontSize float64) textLine {
	return textLine{
		Text:       text,
		FontFamily: font,
		FontSize:   fontSize,
	}
}
