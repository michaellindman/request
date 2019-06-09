package request

import (
    //"fmt"
    "log"
    "net/http"
    //"context"
    "io/ioutil"
    "encoding/json"
    "time"
    "strings"
)

func url(path string) string {
    return string("https://forum.0cd.xyz/" + path)
}

func Request(path string) (b []byte) {
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
        return body
    }
    return nil
}

func About() (*AutoGen) {
    var about AutoGen
    json.Unmarshal(Request("about"), &about)
    for i := 0; i < len(about.About.Admins); i++ {
        about.About.Admins[i].AvatarTemplate = strings.ReplaceAll(about.About.Admins[i].AvatarTemplate, "{size}", "120")
    }
    return &about
}

func GetTopics(path string) (*TagTopics) {
    var topics TagTopics
    json.Unmarshal(Request(path), &topics)
    return &topics
}

func Topics(path string) {

}