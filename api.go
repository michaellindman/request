package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/structs"
)

func url(path string) string {
	return string("https://forum.0cd.xyz/" + path)
}

// Request sends GET to path
func Request(path string) []byte {
	req, err := http.NewRequest("GET", url(path), nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", "57dc92f574f7f400e1d12670940c19930a1f85b78485ca25ac96d96590dc7f99")
	req.Header.Set("Api-Username", "Michael")

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//fmt.Printf("%s %s %d\n", resp.Request.Method, url("about.json"), resp.StatusCode)

	if resp.StatusCode == 200 {
		/*var result map[string]interface{}
		json.Unmarshal([]byte(body), &result)
		return result*/
		return body
	}
	return nil
}

// About gets json data from about page
func About() map[string]interface{} {
	var about AutoGen
	json.Unmarshal(Request("about"), &about)
	for i := 0; i < len(about.About.Admins); i++ {
		about.About.Admins[i].AvatarTemplate = strings.ReplaceAll(about.About.Admins[i].AvatarTemplate, "{size}", "120")
	}
	//return about
	m := structs.Map(about)
	return m
}

/*// GetTopics gets topic list from tag
func GetTopics(path string) *TagTopics {
	var topics TagTopics
	json.Unmarshal(Request(path), &topics)
	return &topics
}*/

// Topics n/a
func Topics(path string) {
}
