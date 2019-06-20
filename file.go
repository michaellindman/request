package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/structs"
)

// File reads request json file
func File(path string) ([]byte, *HTTPError) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Println("Error reading request. ", err)
		return nil, &HTTPError{Status: "500 Internal Server Error", StatusCode: 500}
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, nil
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

func Header() (headers *Headers, err *HTTPError) {
	resp, err := File("./assets/json/headers.json")
	if err != nil {
		return
	}
	json.Unmarshal(resp, &headers)
	return
}
