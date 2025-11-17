package api

import "encoding/json"

type Result[T any] struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type PageResult[T any] struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

func ResultSuccess() string {
	result := Result[string]{
		Code:    200,
		Message: "success",
		Data:    "",
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func ResultMsg(msg string) string {
	result := Result[string]{
		Code:    200,
		Message: msg,
		Data:    "",
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func ResultData[T any](data T) string {
	result := Result[T]{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func ResultMsgData[T any](msg string, data T) string {
	result := Result[T]{
		Code:    200,
		Message: msg,
		Data:    data,
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
