package gophers

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "net/url"
//"reflect"
  "encoding/json"
)

// YodaGopher replies with Yoda transformation of text quotes instead of relaying the original message.
type YodaGopher struct {
	endpoint string
}

func NewYodaGopher() YodaGopher {
	return YodaGopher{
         //http://funtranslations.com/api  Seems limited to 5 calls per hour
         endpoint: "http://api.funtranslations.com/translate/yoda.json",
	}
}

type YodaAPIResponse struct {
  Success string
  Contents yoda_api_data
}

type yoda_api_data struct {
  Translation string
  Text string
  Translated string
}

func (g YodaGopher) TransformMessage(msg string) string {
  data := url.Values{ "text": { msg} }
  response, err := http.PostForm(g.endpoint, data)

  if err != nil {
    panic(err.Error())
  }

  // Need to take the body of the response and turn into a byte array to Unmarshal
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    panic(err.Error())
  }

  var api_struct = new(YodaAPIResponse)
  err = json.Unmarshal(body, &api_struct)
  if err != nil {
    fmt.Println("Unmarshalling did not go well...: ", err)
  }

  return string(api_struct.Contents.Translated)
}
