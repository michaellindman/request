package request

import (
	"io/ioutil"
	"log"
	"os"

	"0cd.xyz-go/logger"
)

// File reads request json file
func File(path string) ([]byte, *logger.HTTPError) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Println("Error reading request. ", err)
		return nil, &logger.HTTPError{Status: "500 Internal Server Error", StatusCode: 500}
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, nil
}
