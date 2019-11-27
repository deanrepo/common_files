package main

import (
	"fmt"
	"net/url"
)

func main() {
	// url encode
	v := url.Values{}
	v.Add("msg", "此订单不存在或已经提交")
	body := v.Encode()
	fmt.Println(v)
	fmt.Println(body)
	// url decode
	m, _ := url.ParseQuery(body)
	fmt.Println(m)

	u := url.Values{}
	u.Set("a", "1")
	u.Set("b", "2")
	u.Set("data", `{"sessionId":"0bts0W1DWKm70B4UZq3V1h3r2DpsbhDsc2WD","eventId":"8F2qNf0bts0W1DWKm70B4UZq3V1h3r2Dpsbh","androidId":"b22f3d41736f748c","userAgent":"Mozilla/5.0 (Linux; Android 5.1; MI PAD 2 Build/LMY47I; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 Safari/537.36","osVersion":"5.1.1","bundleId":"com.brianbaek.popstar","connectionType":"wifi","deviceMake":"Xiaomi","deviceModel":"MI PAD 2","language":"zh_CN","timeZone":"GMT+08:00","campaignId":3261,"mac":"38:a4:ed:fe:99:c8"}`)
	fmt.Println(u.Encode())
}
