package main

import (
 "bytes"
 "encoding/json"
 "net/http"
)

func main() {
  yourData := `{"hello": "world"}`
  
  jsonData, _ := json.Marshal(yourData)
  if err != nil {
    // Handle the error
  }
  
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
  if err != nil {
    // Handle the error
  }
  
  req.Header.Set("Content-Type", "application/json")
  
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    // Handle the error
  }
  defer resp.Body.Close()
}

// https://medium.com/@samix.ys/how-to-send-json-in-a-http-post-request-in-go-f8114f385892
