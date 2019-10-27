package request

import (
	"io/ioutil"
	"log"
	"os"
)

// File reads request json file
func File(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Println("Error reading request. ", err)
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, nil
}
