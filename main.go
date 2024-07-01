package main

import (
  "github.com/cloudwego/hertz/pkg/app/server"
  "go-llm-simulator/controller"
)

func main() {
  h := server.Default()
  controller.RegisterHadlder(h)
  h.Spin()
}
