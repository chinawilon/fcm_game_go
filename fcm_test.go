package fcm

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	// Preset parameters
	fcm, err := NewFcm("e44158030c7341819aedf04a147f3e8a", "1101999999", "d59bbdefd68b71f906c4d67e52841700")
	if err != nil {
		t.Errorf("new fcm err : %v", err)
	}

	// proxy for capture package
	proxy, _ := url.Parse("http://127.0.0.1:9091")
	fcm.SetClient(
		&http.Transport{
			Proxy: http.ProxyURL(proxy),
		}, time.Second*10)

	// test check
	response, err := fcm.TestCheck(
		&Check{Ai: "100000000000000001", Name: "某一一", IdNum: "110000190101010001"}, "yA2RxS")

	if err != nil {
		t.Errorf("fcm check err : %v", err)
	}

	// status
	if response.StatusCode != http.StatusOK {
		t.Errorf("fcm check status code err : %v", response.StatusCode)
	}

	// body
	if response.Body != nil {
		defer response.Body.Close()
	}

	// response body info
	body, _ := ioutil.ReadAll(response.Body)
	info := Status{}

	err = json.Unmarshal(body, &info)
	if err != nil {
		t.Errorf("unmarshal err : %v", err)
	}

	if info.ErrCode != 0 {
		t.Errorf("response err : %v", info)
	}
}

func TestQuery(t *testing.T) {
	// Preset parameters
	fcm, err := NewFcm("e44158030c7341819aedf04a147f3e8a", "1101999999", "d59bbdefd68b71f906c4d67e52841700")
	if err != nil {
		t.Errorf("new fcm err : %v", err)
	}

	// proxy for capture package
	proxy, _ := url.Parse("http://127.0.0.1:9091")
	fcm.SetClient(
		&http.Transport{
			Proxy: http.ProxyURL(proxy),
		}, time.Second*10)

	// check
	response, err := fcm.TestQuery(
		&Query{Ai: "100000000000000001"}, "HHatGD")

	if err != nil {
		t.Errorf("fcm check err : %v", err)
	}

	// status
	if response.StatusCode != http.StatusOK {
		t.Errorf("fcm check status code err : %v", response.StatusCode)
	}

	// body
	if response.Body != nil {
		defer response.Body.Close()
	}

	// response body info
	body, _ := ioutil.ReadAll(response.Body)
	info := Status{}

	err = json.Unmarshal(body, &info)
	if err != nil {
		t.Errorf("unmarshal err : %v", err)
	}

	if info.ErrCode != 0 {
		t.Errorf("response err : %v", info)
	}
}

func TestLoginOrOut(t *testing.T) {
	// Preset parameters
	fcm, err := NewFcm("e44158030c7341819aedf04a147f3e8a", "1101999999", "d59bbdefd68b71f906c4d67e52841700")
	if err != nil {
		t.Errorf("new fcm err : %v", err)
	}

	// proxy for capture package
	proxy, _ := url.Parse("http://127.0.0.1:9091")
	fcm.SetClient(
		&http.Transport{
			Proxy: http.ProxyURL(proxy),
		}, time.Second*10)

	// upload collection
	// 游客用户
	di := makeMD5("device")
	behavior := Behavior{
		No: 1,
		Si: di,
		Bt: 1,
		Ot: time.Now().Unix(),
		Ct: 2,
		Di: di,
		Pi: "",
	}
	//// 已认证通过用户
	//pi := "1fffbjzos82bs9cnyj1dna7d6d29zg4esnh99u"
	//behavior := Behavior{
	//	No: 2,
	//	Si: pi,
	//	Bt: 1,
	//	Ot: time.Now().Unix(),
	//	Ct: 0,
	//	Pi: pi,
	//}
	collections := &Collections{Collections: &[]Behavior{behavior}}
	response, err := fcm.TestLoginOrOut(collections, "HHatGD")

	if err != nil {
		t.Errorf("fcm check err : %v", err)
	}

	// status
	if response.StatusCode != http.StatusOK {
		t.Errorf("fcm check status code err : %v", response.StatusCode)
	}

	// body
	if response.Body != nil {
		defer response.Body.Close()
	}

	// response body info
	body, _ := ioutil.ReadAll(response.Body)
	info := Status{}

	err = json.Unmarshal(body, &info)
	if err != nil {
		t.Errorf("unmarshal err : %v", err)
	}

	if info.ErrCode != 0 {
		t.Errorf("response err : %v", info)
	}
}

func makeMD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	s := ctx.Sum(nil)
	return hex.EncodeToString(s)
}
