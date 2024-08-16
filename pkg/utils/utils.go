package utils

import (
	"os"

	"github.com/gocarina/gocsv"
)

func ReadCSV[T any](filePath string) (records []*T, err error) {
	csvFile, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	var data []*T
	if err := gocsv.UnmarshalFile(csvFile, &data); err != nil {
		return nil, err
	}

	return data, nil
}
