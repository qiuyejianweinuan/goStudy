package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

const ( //全局变量
	BASEURL     = "https://api.openai.com/v1/completions"
	maxTokens   = 2000 //openai api的参数配置
	temperature = 0.7
	engine      = "text-davinci-003"
)

// ChatGPTResponseBody 响应体 接受chatGpt的响应数据
type ChatGPTResponseBody struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChoiceItem           `json:"choices"`
	Usage   map[string]interface{} `json:"usage"`
}

type ChoiceItem struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	Logprobs     int    `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

// ChatGPTRequestBody 请求体 用来包装数据向chatGpt api发送请求
type ChatGPTRequestBody struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
}

func Completions(msg string) (string, error) {
	//封装请求数据,发放参数为发送的消息数据,封装到chatGpt的请求数据中
	requestBody := ChatGPTRequestBody{
		Model:            engine,
		Prompt:           msg,
		MaxTokens:        maxTokens,
		Temperature:      temperature,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	log.Printf("request gtp json string : %v", string(requestData))
	//配置请求
	req, err := http.NewRequest("POST", BASEURL, bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	//获取配置文件中的openai_key
	apiKey := viper.GetString("OPENAI_KEY")
	//配置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}

	//向chatGpt发送请求数据
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("gtp api status code not equals 200,code is %d", response.StatusCode))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}

	var reply string
	if len(gptResponseBody.Choices) > 0 {
		//获取chatGpt返回响应的数据
		reply = gptResponseBody.Choices[0].Text
	}
	//输出api返回的消息
	log.Printf("gpt response text: %s \n", reply)
	return reply, nil
}

func FormatQuestion(question string) string {
	return "Answer:" + question
}
