package request

import (
	"io/ioutil"
	"log"
	"os"
)

// File reads data from file
func File(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Println("Error reading request. ", err)
		return nil, err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return byteValue, nil
}
