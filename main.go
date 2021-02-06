package main

import "fmt"
import "repath.io/config"
import "repath.io/data"

func main() {
  configuration := config.LoadConfiguration()
  fmt.Printf("initializing %s\n", configuration.ServiceName)
}