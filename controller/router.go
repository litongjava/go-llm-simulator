package controller

import "github.com/cloudwego/hertz/pkg/app/server"

func RegisterHadlder(h *server.Hertz) {
  h.GET("/PingHandler", PingHandler)
  h.POST("/v1/chat/completions", ChatCompletions)
  h.POST("/v1/embeddings", EmbeddingIndex)
}
