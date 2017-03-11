package main

import (
     "ipfinder"
     "fmt"
     )

func main() {
  url :="http://ipecho.net/plain"

  //Passing url to MakeRequest
  ip := ipfinder.MakeRequest(url)

  // Priting IP
  fmt.Println(ip)
}