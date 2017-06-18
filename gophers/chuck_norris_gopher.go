package gophers

import (
  "net/http"
//"io/ioutil"
//"fmt"
  "encoding/json"
)

// ChuckNorrisGopher replies with Chuck Norris quotes instead of relaying the original message.
type ChuckNorrisGopher struct {
	endpoint string
}

func NewChuckNorrisGopher() ChuckNorrisGopher {
	return ChuckNorrisGopher{
		endpoint: "http://api.icndb.com/jokes/random",
	}
}

type ChuckNorrisAPIResponse struct {
  Response string
  Value struct {
    Id int
    Joke string }
}

func (g ChuckNorrisGopher) TransformMessage(msg string) string {
  response, err := http.Get(g.endpoint)
  if err != nil {
    panic(err.Error())
  }

  // Need to take the body of the response and turn into a byte array
//body, err := ioutil.ReadAll(response.Body)
//if err != nil {
//  panic(err.Error())
//}

//var api_struct = new(ChuckNorrisAPIResponse)
//err = json.Unmarshal(body, &api_struct)
//if err != nil {
//  fmt.Println("Unmarshalling did not go well...: ", err)
//}
  var api_struct = new(ChuckNorrisAPIResponse)
  err = json.NewDecoder(response.Body).Decode(&api_struct)
  return string(api_struct.Value.Joke)
	// Helpful links:
	// * https://golang.org/pkg/net/http/#example_Get
	// * https://golang.org/pkg/encoding/json/#example_Decoder_Decode_stream
}
