package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"0cd.xyz-go/logger"
	"github.com/fatih/structs"
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

func Header() (headers *Headers, err *logger.HTTPError) {
	resp, err := File("./assets/json/headers.json")
	if err != nil {
		return
	}
	json.Unmarshal(resp, &headers)
	return
}

func Option() (options *Options) {
	resp, err := File("./assets/json/options.json")
	if err != nil {
		return
	}
	json.Unmarshal(resp, &options)
	return
}

func Contact() map[string]interface{} {
	var contact Contacts
	resp, err := File("./assets/json/contacts.json")
	if err != nil {
		m := structs.Map(err)
		return m
	}
	json.Unmarshal(resp, &contact)
	m := structs.Map(contact)
	return m
}
