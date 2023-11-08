package csv

import (
	"encoding/csv"
	"marckent04/certificate_generator/domain"
	"marckent04/certificate_generator/shared"
	"os"
)

type Parser struct {
	file         *os.File
	Certificates []domain.Certificate
}

func (p *Parser) parse(path string) {
	p.file = shared.LoadFile(path)
	defer p.file.Close()
	reader := csv.NewReader(p.file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	shared.HandleError(err)

	p.Certificates = createCertificates(data)
}
