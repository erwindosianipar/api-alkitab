package utils

import (
	"encoding/csv"
	"os"

	"github.com/sirupsen/logrus"
)

// CSVReader is commons to read csv file
func CSVReader(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		logrus.Fatal("Unable to read input file", err)
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		logrus.Fatal("Unable to parse file as CSV", err)
		return nil, err
	}

	return records, nil
}
