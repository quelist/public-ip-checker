package ipfinder
//A requesting packages.

import (
     "time"
     "net/http"
     "io/ioutil"
)
func MakeRequest(url string) string{
  timeout := time.Duration(5 * time.Second)
   client := http.Client{
        Timeout: timeout,
        }
  resp, err := client.Get(url)
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "Connection Problem" 
  }
  return string(body)
}