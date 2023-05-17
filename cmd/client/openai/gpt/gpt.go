// Package gpt
// Create  2023-03-21 00:18:44
package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	GPT_API  = "https://api.openai.com/v1/chat/completions"
	ApiKey   = "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"
	gptModel = "gpt-3.5-turbo-0301"
)

var ErrGptReqEmpty = fmt.Errorf("GPT REQ EMPTY")

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string     `json:"model"`
	Messages []*Message `json:"messages"`
}

type response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptToken      int `json:"prompt_token"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []*struct {
		*Message     `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

// Do 执行请求.
func Do(content string) (string, error) {

	body := &Request{
		Model: gptModel,
		Messages: []*Message{{
			Role:    "user",
			Content: content,
		}},
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", GPT_API, bytes.NewReader(byteBody))
	if err != nil {
		return "", ErrGptReqEmpty
	}
	req.Header.Set("Authorization", "Bearer  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")
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{
		Timeout: 10 * time.Minute,
	}
	rsp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	if rsp.StatusCode != http.StatusOK {
		return "", ErrGptReqEmpty
	}

	allDataBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	rspData := &response{}
	if err = json.Unmarshal(allDataBody, rspData); err != nil {
		return "", err
	}
	if len(rspData.Choices) > 0 {
		return rspData.Choices[0].Message.Content, nil
	}
	return "", ErrGptReqEmpty
}
