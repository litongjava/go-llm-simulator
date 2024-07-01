package controller

import (
  "context"
  "encoding/json"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/common/utils"
  "github.com/hertz-contrib/sse"
  "github.com/sashabaranov/go-openai"
)

func ChatCompletions(ctx context.Context, requestCtx *app.RequestContext) {
  authorization := requestCtx.GetHeader("Authorization")
  if authorization == nil {
    requestCtx.JSON(401, utils.H{
      "error": utils.H{
        "message": "You didn't provide an API key. You need to provide your API key in an Authorization header using Bearer auth (i.e. Authorization: Bearer YOUR_KEY), or as the password field (with blank username) if you're accessing the API from your browser and are prompted for a username and password. You can obtain an API key from https://platform.openai.com/account/api-keys.",
        "type":    "invalid_request_error",
        "param":   nil,
        "code":    nil,
      },
    })
    return
  }
  var reqVo openai.ChatCompletionRequest
  err := requestCtx.BindJSON(&reqVo)
  if err != nil {
    requestCtx.JSON(400, map[string]string{"error": "Invalid JSON"})
    return
  }
  if reqVo.Stream {
    var emitter = sse.NewStream(requestCtx)

    var streamRespVo = openai.ChatCompletionStreamResponse{
      ID:                "chatcmpl-9gBByG3g99HJzzwrtuxgHBQhuR87c",
      Object:            "chat.completion.chunk",
      Created:           1719839826,
      Model:             "gpt-4o-2024-05-13",
      SystemFingerprint: "fp_d576307f90",
    }

    var delta = openai.ChatCompletionStreamChoiceDelta{
      Role:    "assistant",
      Content: "",
    }
    var choices = []openai.ChatCompletionStreamChoice{{
      Index:                0,
      Delta:                delta,
      ContentFilterResults: nil,
    }}

    streamRespVo.Choices = choices
    push(requestCtx, emitter, streamRespVo)

    delta = openai.ChatCompletionStreamChoiceDelta{
      Content: "This",
    }
    choices = []openai.ChatCompletionStreamChoice{{
      Index: 0,
      Delta: delta,
    }}

    streamRespVo.Choices = choices
    push(requestCtx, emitter, streamRespVo)
    streamRespVo.Choices[0].Delta.Content = " is"
    push(requestCtx, emitter, streamRespVo)

    streamRespVo.Choices[0].Delta.Content = " ChatGPT"
    push(requestCtx, emitter, streamRespVo)

    streamRespVo.Choices[0].Delta.Content = "."
    push(requestCtx, emitter, streamRespVo)

    delta = openai.ChatCompletionStreamChoiceDelta{}
    choices = []openai.ChatCompletionStreamChoice{{
      Index:        0,
      Delta:        delta,
      FinishReason: "stop",
    }}
    streamRespVo.Choices = choices
    push(requestCtx, emitter, streamRespVo)

  } else {
    var choices = []openai.ChatCompletionChoice{{
      Index: 0,
      Message: openai.ChatCompletionMessage{
        Role:    "assistant",
        Content: "this is ChatGPT",
      },
      FinishReason: "stop",
      LogProbs:     nil,
    }}
    var usage = openai.Usage{
      PromptTokens:     10,
      CompletionTokens: 26,
      TotalTokens:      36,
    }

    var respVo = openai.ChatCompletionResponse{
      ID:                "chatcmpl-9gAs6fnGjUuW5phZ8fzqkoLo1nFtA",
      Object:            "chat.completion",
      Created:           1719838594,
      Model:             "gpt-4o-2024-05-13",
      Choices:           choices,
      Usage:             usage,
      SystemFingerprint: "fp_d576307f90",
    }
    requestCtx.JSON(200, respVo)
  }

}

func push(requestCtx *app.RequestContext, emitter *sse.Stream, streamRespVo openai.ChatCompletionStreamResponse) {
  //data, err := sonic.Marshal(streamRespVo)
  data, err := json.Marshal(streamRespVo)
  if err != nil {
    requestCtx.JSON(401, utils.H{"error": err.Error()})
  }
  event := &sse.Event{
    Data: data,
  }
  emitter.Publish(event)
}
