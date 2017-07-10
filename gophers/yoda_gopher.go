package gophers

import (
  "net/http"
  "io/ioutil"
  "time"
  "fmt"
)

// YodaGopher replies with Yoda transformation of text quotes instead of relaying the original message.
type YodaGopher struct {
	endpoint string
}

func NewYodaGopher() YodaGopher {
	return YodaGopher{
         endpoint: "https://yoda.p.mashape.com/yoda",
	}
}

func ReadMashapeKey() string {
  key, err := ioutil.ReadFile("MashapeKey")
  if err != nil {
    fmt.Println("No MashapeKey file present.")
  }
  return string(key)
}

func (g YodaGopher) TransformMessage(msg string) string {
  client := &http.Client{Timeout: 40* time.Second}
  request, _ := http.NewRequest("GET", g.endpoint, nil)

  request.Header.Add("X-Mashape-Key", "1MkT3DPx72mshEFqgxB06DTVo36ap1WZjvxjsnpxluxXdP5Kf5")
  request.Header.Add("Accept", "text/plain")

  query := request.URL.Query()
  query.Add("sentence", msg)

  request.URL.RawQuery = query.Encode()

  response, err := client.Do(request)

  if err != nil {
    fmt.Println("errored request:", err.Error())
  }

  // Need to take the body of the response.  It is not json, just plain text.
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    panic(err.Error())
  }

  return string(body)
}
