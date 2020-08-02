# Request

simple golang package for making API requests

## Documentation

Documentation for the project can be found on [pkg.go.dev](https://pkg.go.dev/github.com/michaellindman/request?tab=doc)

## Installation

With a correctly configured golang toolchain:

```sh
go get -u github.com/michaellindman/request
```

## Examples

Requesting and encoding structured data:

```go
type Topics struct {
    TopicList struct {
        Topics []struct {
            ID         int    `json:"id"`
            Title      string `json:"title"`
            PostsCount int    `json:"posts_count"`
            CreatedAt  string `json:"created_at"`
            Slug       string `json:"slug"`
        } `json:"topics"`
    } `json:"topic_list"`
}

func topics() (topics *Topics, err error) {
    headers := map[string]string{
        "Accept":       "application/json",
        "Content-Type": "application/json",
    }
    resp, err := request.API(http.MethodGet, "https://forum.0cd.xyz/t/nixos-vfio-pcie-passthrough/277", headers, nil)
    if err != nil {
        return nil, err
    }
    err = json.Unmarshal(resp.Body, &topics)
    if err != nil {
        return nil, err
    }
    return topics, nil
}
```

Requesting and encoding unstructured data:

```go
users := make(map[string]interface{})
headers := map[string]string{
    "Accept":       "application/json",
    "Content-Type": "application/json",
}

resp, err := API(http.MethodGet, "https://reqres.in/api/users", headers, nil)
if err != nil {
    log.Println(err)
    return
}

err = json.Unmarshal(resp.Body, &users)
if err != nil {
    log.Println(err)
    return
}

fmt.Println(users["data"])
```

Asynchronous API requests:

```go
func topics(urls []string) (topicBody [][]byte, error err) {
    headers := map[string]string{
    "Accept":       "application/json",
    "Content-Type": "application/json",

    ch := make(chan *request.Response)
    var wg sync.WaitGroup

    for _, path := range urls {
        wg.Add(1)
        go request.AsyncAPI(http.MethodGet, path, headers, nil, ch, &wg)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for resp := range ch {
        if resp.Error != nil {
            return err
        }
        topicBody = append(topicBody, resp.Body)
    }
    return topicBody, nil
}
```

Sending requests with data using post, put, etc:

```go
headers := map[string]string{
    "Accept":       "application/json",
    "Content-Type": "application/json",
}

req := `{
    "on": true,
    "bri": 255
}`

_, err := request.API(http.MethodPut, "http://10.0.40.2/api/<apikey>/lights/4/state", headers, bytes.NewBuffer([]byte(req)))
if err != nil {
    log.Println(err)
}
```

## License

```text
MIT License Copyright (c) 2020 Michael Lindman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is furnished
to do so, subject to the following conditions:

The above copyright notice and this permission notice (including the next
paragraph) shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS
OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF
OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
