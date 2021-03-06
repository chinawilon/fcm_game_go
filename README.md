# FCM 网络游戏防沉迷实名认证系统包 - go

# Install
```shell script
go get github.com/chinawilon/fcm_game_go
```

# Example
*需要通过所有的测试案例，测试案例会有测试码，全部通过以后才可以使用正式接口地址*

```go

package fcm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestCheck(t *testing.T)  {
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
		}, time.Second * 10)

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

func TestQuery(t *testing.T)  {
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
		}, time.Second * 10)

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


func TestLoginOrOut(t *testing.T)  {
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
		}, time.Second * 10)

	// check
	response, err := fcm.TestLoginOrOut(
		&[]Behavior{{No: 1, Bt: 0, Ct: 2, Di: "fjfkfjfkfjfkfjfkfjfkfjjjjjjsjjss"}}, "HHatGD")

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

```

# Licence
MIT