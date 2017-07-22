package main

import "encoding/json"
import "fmt"
import "os"

var settings struct {
  NumRequests int `json:"numRequests"`
  NumError int `json:"errors"`
}
func main() {
  configFile, err := os.Open("config.json")
  if err != nil {
    fmt.Println("ERROR: Opening config file:", err.Error())
  }

  jsonParser := json.NewDecoder(configFile)
  if err = jsonParser.Decode(&settings); err != nil {
    fmt.Println("ERROR: Parsing config file:", err.Error())
  }
  fmt.Printf("%d %d",settings.NumRequests, settings.NumError)
  return
}
