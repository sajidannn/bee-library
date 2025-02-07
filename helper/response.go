package helper

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}