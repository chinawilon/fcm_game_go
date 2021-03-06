package fcm

import (
	"crypto/cipher"
	"net/http"
)

// fcm
type fcm struct {
	AppID string
	BizID string
	Key string
	aes cipher.AEAD
	keys []string
	client http.Client
}

// check data
type Check struct {
	Ai string `json:"ai"`
	Name string `json:"name"`
	IdNum string `json:"idNum"`
}

// query data
type Query struct {
	Ai string `json:"ai"`
}

// login or logout
type Behavior struct {
	No int `json:"no"`
	Si string `json:"si"`
	Bt int `json:"bt"`
	Ot int64 `json:"ot"`
	Ct int `json:"ct"`
	Di string `json:"di"`
	Pi string `json:"pi"`
}

// response status
type Status struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}
