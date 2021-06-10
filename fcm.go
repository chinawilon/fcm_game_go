package fcm

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// constructor
func NewFcm(appId, bizId, key string) (*fcm, error) {
	fcm := &fcm{
		AppID: appId,
		BizID: bizId,
		Key:   key,
		keys:  []string{"appId", "bizId", "timestamps"},
		client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
	// cipher
	b, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(b)
	if err != nil {
		return nil, err
	}
	// aead
	AEAD, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	fcm.aes = AEAD
	return fcm, nil
}

// check api
func (f *fcm) Check(c *Check) (*http.Response, error) {
	url := "https://api.wlc.nppa.gov.cn/idcard/authentication/check"
	header := f.getHeader()
	header["Content-Type"] = []string{"application/json; charset=utf-8"}
	return f.request("POST", url, c, header, nil)
}

// test check api
func (f *fcm) TestCheck(c *Check, testCode string) (*http.Response, error) {
	uri := "https://wlc.nppa.gov.cn/test/authentication/check/" + testCode
	header := f.getHeader()
	header["Content-Type"] = []string{"application/json; charset=utf-8"}
	return f.request("POST", uri, c, header, nil)
}

// query api
func (f *fcm) Query(q *Query) (*http.Response, error) {
	uri := "http://api2.wlc.nppa.gov.cn/idcard/authentication/query"
	header := f.getHeader()
	var data map[string]string
	if value, err := json.Marshal(q); err == nil {
		if err := json.Unmarshal(value, &data); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	uri = fmt.Sprintf("%s?%s=%s", uri, "ai", q.Ai)
	return f.request("GET", uri, nil, header, data)
}

// test query api
func (f *fcm) TestQuery(q *Query, testCode string) (*http.Response, error) {
	uri := "https://wlc.nppa.gov.cn/test/authentication/query/" + testCode
	header := f.getHeader()
	var data map[string]string
	if value, err := json.Marshal(q); err == nil {
		if err := json.Unmarshal(value, &data); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	uri = fmt.Sprintf("%s?%s=%s", uri, "ai", q.Ai)
	return f.request("GET", uri, nil, header, data)
}

// login or logout
func (f *fcm) LoginOrOut(c *Collections) (*http.Response, error) {
	url := "http://api2.wlc.nppa.gov.cn/behavior/collection/loginout"
	header := f.getHeader()
	header["Content-Type"] = []string{"application/json; charset=utf-8"}
	return f.request("POST", url, c, header, nil)
}

// test login or logout
func (f *fcm) TestLoginOrOut(c *Collections, testCode string) (*http.Response, error) {
	uri := "https://wlc.nppa.gov.cn/test/collection/loginout/" + testCode
	header := f.getHeader()
	header["Content-Type"] = []string{"application/json; charset=utf-8"}
	return f.request("POST", uri, c, header, nil)
}

// aes-128-gcm + base64
func (f *fcm) makeBody(body []byte) (string, error) {
	//random bytes
	nonce := make([]byte, f.aes.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	data := append(nonce, f.aes.Seal(nil, nonce, body, nil)...)
	return base64.StdEncoding.EncodeToString(data), nil
}

// sha256
func (f *fcm) makeSign(header http.Header, body string, query map[string]string) string {
	// except content-type
	header.Del("Content-Type")
	ks := f.keys
	for k, v := range query {
		ks = append(ks, k)
		header[k] = []string{v} //maybe lower case
	}
	sort.Strings(ks)
	raw := ""
	for _, k := range ks {
		raw += k + header[k][0]
	}
	hash := sha256.New()
	d := append(append([]byte(f.Key), raw...), body...)
	hash.Write(d)
	return hex.EncodeToString(hash.Sum(nil))
}

// get the header
func (f *fcm) getHeader() http.Header {
	return http.Header{
		"appId":      []string{f.AppID},
		"bizId":      []string{f.BizID},
		"timestamps": []string{strconv.FormatInt(time.Now().Unix()*1000, 10)},
	}
}

// set client
func (f *fcm) SetClient(transport http.RoundTripper, timeout time.Duration) {
	f.client = http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
}

// request
func (f *fcm) request(method, uri string, b interface{}, header http.Header, query map[string]string) (*http.Response, error) {
	var body []byte
	var err error
	if b != nil {
		body, err = json.Marshal(b)
		if err != nil {
			return nil, err
		}
	} else {
		body = []byte(`{}`)
	}

	data, err := f.makeBody(body)
	if err != nil {
		return nil, err
	}
	raw := `{"data":"` + strings.TrimRight(data, "=") + `"}`
	header["sign"] = []string{f.makeSign(header.Clone(), raw, query)}
	req, err := http.NewRequest(method, uri, bytes.NewReader([]byte(raw)))
	if err != nil {
		return nil, err
	}
	req.Header = header
	return f.client.Do(req)
}
