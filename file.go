package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/structs"
)

// File reads request json file
func File(path string) (b []byte) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func Contact() map[string]interface{} {
	var contact Contacts
	json.Unmarshal(File("./json/contacts.json"), &contact)
	m := structs.Map(contact)
	return m
}

func Header() (headers Headers) {
	json.Unmarshal(File("./json/headers.json"), &headers)
	return
}
