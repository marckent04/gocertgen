package shared

import (
	"fmt"
	"log"
	"marckent04/certificate_generator/domain"
	"os"
	"runtime/debug"
	"strings"
)

func LoadFile(path string) *os.File {
	file, err := os.Open(path)

	HandleError(err)

	return file
}

func ReadFile(path string) string {
	file, err := os.ReadFile(path)

	HandleError(err)

	return string(file)
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(string(debug.Stack()))
	}
}

func GenerateFileName(certificate domain.Certificate) string {
	name := strings.ToLower(strings.ReplaceAll(certificate.Participant, " ", "_"))
	date := certificate.DeliveredAt.Format("02_01_2006")

	return fmt.Sprintf("%s_%s", name, date)
}

func CreateOutputFolderIfNotExists(folder string) {
	_, err := os.Stat(folder)

	if os.IsNotExist(err) {
		err := os.MkdirAll(folder, os.ModePerm)
		HandleError(err)
	}

}
