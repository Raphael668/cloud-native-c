package controllers

type ResCode int

const (
	RES_OK                 ResCode = 100
	RES_BALANCE_NOT_ENOUGH ResCode = 200
	RES_ERROR_UNKNOWN      ResCode = 300
	RES_ERROR_BAD_REQUEST  ResCode = 301
	RES_USER_INVALID_TOKEN ResCode = 302
)

type ResError struct {
	Title string `json:"title"`
	Desc  string `json:"description"`
}

type ResBody struct {
	ResCode ResCode `json:"resCode"`
	// Error   *ResError   `json:"error,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
