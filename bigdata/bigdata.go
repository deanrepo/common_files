package main

import (
	"fmt"

	"github.com/gpmgo/gopm/modules/log"
)

var BigData *BigDataManager

const CHAN_SIZE = 1000

func init() {
	BigData = &BigDataManager{
		IsOpen:      false,
		Url:         "http://receiver.ta.thinkingdata.cn:9080/sync_data",
		ContentType: "application/json;charset=utf-8",
		PostChan:    make(chan []byte, CHAN_SIZE),
	}
}

type BigDataManager struct {
	IsOpen      bool   // 是否启用大数据分析
	Url         string // 上传数据的url
	ContentType string
	PostChan    chan []byte // 上传数据的channel
}

type PostResust struct {
	Code int    `json:"code"` // 上传数据返回的结果状态码，
	Msg  string `josn:"msg"`  // 上传数据返回的消息
}

// 上传数据
func (this *BigDataManager) PostData() {
	if this.IsOpen {
		for {
			select {
			case data := <-this.PostChan:
				// reqBody := bytes.NewBuffer(data)
				// resp, _ := http.Post(this.Url, this.ContentType, reqBody)
				// body, _ := ioutil.ReadAll(resp.Body)

				// var ret PostResust
				// if err := json.Unmarshal(body, &ret); err != nil {
				// 	log.Info("反序列化PostResust err: %v\n", err)
				// 	return
				// }
				// if ret.Code != 0 {
				// 	log.Info("上传数据：%s失败, 失败msg: %s\n", string(data), ret.Msg)
				// }

				// debug
				fmt.Printf("receive data: %s\n", string(data))
			default:

			}
		}

	}

	log.Info("大数据采集模块未开启!")
	<-done
}
