package csv

import (
	"fmt"
	"log"
	"marckent04/certificate_generator/domain"
	"time"
)

func createCertificates(data [][]string) (certificates []domain.Certificate) {
	for _, certificateInfo := range data {
		date, err := time.Parse("02-01-2006", certificateInfo[2])
		if err != nil {
			fmt.Println(err)
			log.Fatal(fmt.Errorf("the date format of the row of \"%s\" is incorrect", certificateInfo[0]))
		}

		cert := domain.Certificate{
			Participant: certificateInfo[0],
			Course:      certificateInfo[1],
			DeliveredAt: date,
		}

		certificates = append(certificates, cert)
	}

	return
}

func NewParser(path string) *Parser {
	parser := &Parser{}
	parser.parse(path)
	return parser
}
