// Package main 调用openai接口,国内通过硅谷云服务调用.
// Create  2023-03-25 00:39:20
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/sashabaranov/go-openai"

	"github.com/gin-gonic/gin"
)

// TODO: 调用openai 不在本地调用,调用硅谷服务器.

/*
	1. https://api.openai.com/v1/chat/completions

curl https://api.openai.com/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $KEY" \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [{"role": "user", "content": "golang 五子棋"}]
  }'

*/

const (
	CodeSuccess = iota
	ErrCodeParamErr
	ErrCodeBuildHttpReq
	ErrCodeHttpResponseErr
	ErrCodeHttpStatusNotSuccess
	ErrCodeReadHttpResponseBody
	ErrCodeUnmarshalHttpResponse
)

const (
	OpenAIUrl   = "https://api.openai.com/v1/chat/completions"
	OpenAIModel = "gpt-3.5-turbo"
)

var AuthHeaderValue string
var OpenAIKey string

func init() {
	OpenAIKey = os.Getenv("OPENAIKEY")
	fmt.Println("openai key:" + OpenAIKey)
	AuthHeaderValue = fmt.Sprintf("Bearer %s", OpenAIKey)
	fmt.Printf("AuthHeader:%+v\n", AuthHeaderValue)
}

// OpenAIResponse OpenAI响应.
type OpenAIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}

// Choice 聊天回复内容.
type Choice struct {
	Message      *Message `json:"message"`
	FinishReason string   `json:"finish_reason"`
	Index        int64    `json:"index"`
}

// Message 具体响应内容.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Usage token使用情况.
type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

// OpenAIRequest 请求OpenAI.
type OpenAIRequest struct {
	Model   string     `json:"model"`
	Message []*Message `json:"message"`
}

// Response ServerResponse.
type Response struct {
	Code    int8   `json:"code"`
	Content string `json:"content"`
}

// NewResponse 构造响应体.
func NewResponse(code int8, msg string) *Response {
	return &Response{
		Code:    code,
		Content: msg,
	}
}

// GetMessageContent 获取MessageContent.
func (o *OpenAIResponse) GetMessageContent() (content string) {

	for i := 0; i < len(o.Choices); i++ {
		if o.Choices[i].Message != nil {
			content = o.Choices[i].Message.Content
		}
	}
	return
}

// QueryOpenAI 查询OpenAI接口.
func QueryOpenAI(c *gin.Context) {

	defer func() {
	}()
	// 不输出日志.
	content := c.Query("content")
	if content == "" {
		c.JSON(http.StatusOK, NewResponse(ErrCodeParamErr, "content is empty"))
		return
	}
	fmt.Printf("query ip:[%+v] content:[%+v]\n", c.ClientIP(), content)
	rsp, err := UserGoGPT(content)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, rsp)
	//openReq := &OpenAIRequest{
	//	Model: OpenAIModel,
	//	Message: []*Message{
	//		{
	//			Role:    "user",
	//			Content: content,
	//		},
	//	},
	//}
	//
	//data, _ := json.Marshal(openReq)
	//httpReq, err := http.NewRequest("POST", OpenAIUrl, bytes.NewReader(data))
	//if err != nil {
	//	c.JSON(http.StatusOK, NewResponse(ErrCodeBuildHttpReq, err.Error()))
	//	return
	//}
	//httpReq.Header.Add("Content-Type", "application/json")
	//httpReq.Header.Add("Authorization", AuthHeaderValue)
	//
	//reqBod, _ := json.Marshal(httpReq)
	//fmt.Printf("http req body:[%+v] hearder:[%+v]\n", string(reqBod), httpReq.Header)
	//
	//client := &http.Client{
	//	Timeout: 5 * time.Minute,
	//}
	//fmt.Println("start query ai")
	//start := time.Now().Unix()
	//httpRsp, err := client.Do(httpReq)
	//if err != nil {
	//	fmt.Println(err)
	//	c.JSON(http.StatusOK, NewResponse(ErrCodeHttpResponseErr, err.Error()))
	//	return
	//}
	//end := time.Now().Unix()
	//fmt.Printf("cost time:%+v status:[%+v] body:[%+v]", end-start, httpRsp.StatusCode, httpRsp)
	//if httpRsp.StatusCode != http.StatusOK {
	//	c.JSON(http.StatusOK, NewResponse(ErrCodeHttpStatusNotSuccess, "http not ok"))
	//	return
	//}
	//
	//// 正常解析.
	//openAIResponse := &OpenAIResponse{}
	//responseBody, err := ioutil.ReadAll(httpRsp.Body)
	//if err != nil {
	//	c.JSON(http.StatusOK, NewResponse(ErrCodeReadHttpResponseBody, err.Error()))
	//	return
	//}
	//if err := json.Unmarshal(responseBody, openAIResponse); err != nil {
	//	c.JSON(http.StatusOK, NewResponse(ErrCodeUnmarshalHttpResponse, err.Error()))
	//	return
	//}
	//ddd, _ := json.Marshal(openAIResponse)
	//fmt.Printf("http response body:[%+v]\n", string(ddd))
	//c.JSON(http.StatusOK, NewResponse(CodeSuccess, openAIResponse.GetMessageContent()))
}

func UserGoGPT(content string) (string, error) {
	client := openai.NewClient(OpenAIKey)
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: OpenAIModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
	})
	if err != nil {
		return "", err
	}

	if len(resp.Choices) < 0 {
		return "", fmt.Errorf("choices < 0 ")
	}
	return resp.Choices[0].Message.Content, nil
}
