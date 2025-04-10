package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Luna-CY/Golang-Project-Template/internal/configuration"
)

type FeishuWriter struct{}

func (cls *FeishuWriter) Write(p []byte) (n int, err error) {
	// if feishu is not enabled, return
	if !configuration.Configuration.Logger.CustomizeWriter.Feishu.Enabled {
		return len(p), nil
	}

	// build text message
	var body = map[string]any{
		"msg_type": "text",
		"content": map[string]any{
			"text": string(p),
		},
	}

	bs, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	request, err := http.NewRequest("POST", configuration.Configuration.Logger.CustomizeWriter.Feishu.Webhook, bytes.NewBuffer(bs))
	if err != nil {
		return 0, err
	}

	request.Header.Set("Content-Type", "application/json")

	// send request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()

	if 200 != response.StatusCode {
		return 0, fmt.Errorf("feishu webhook response status code: %d", response.StatusCode)
	}

	var responseBody struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	if err := json.NewDecoder(response.Body).Decode(&responseBody); nil != err {
		return 0, err
	}

	if 0 != responseBody.Code {
		return 0, fmt.Errorf("feishu webhook response body code: %d, msg: %s", responseBody.Code, responseBody.Msg)
	}

	return len(p), nil
}
