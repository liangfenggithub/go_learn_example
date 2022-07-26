package main

import (
	"encoding/json"
	"fmt"
)

type Metadata_t struct {
	Type_code string `json:"type_code"`
	Src_addr  string `json:"src_addr"`
	Rssi      int    `json:"rssi"`
	Rcv_time  int    `json:"rcv_time"`
}
type UpMsg_t struct {
	Id              int         `json:"id"`
	Msg_type        string      `json:"msg_type"`
	Source_dev_type string      `json:"source_dev_type"`
	Data            interface{} `json:"data"`
	Metadata        Metadata_t  `json:"metadata,omitempty"`
	Is_ack          bool        `json:"is_ack,omitempty` //当结构体有omitempty的时候，如果name为空，那么最终生成的json中没有name字段。
}

func main() {

	var rcvMsg UpMsg_t
	str2 := `{
		"id": 313,
		"msg_type": "confirmed_up",
		"source_dev_type": "node",
		"data": "01010000",
		"metadata": {
			"type_code": "1080",
			"src_addr": "112234",
			"rssi": -21,
			"rcv_time": 1657878680800
		}
	}`

	err2 := json.Unmarshal([]byte(str2), &rcvMsg)
	if err2 != nil {
		fmt.Println("JSON -> struct ERR:", err2)
	}
	fmt.Println("JSON -> struct:", rcvMsg)

	b, err := json.Marshal(rcvMsg)
	if err != nil {
		fmt.Println("struct -> JSON ERR:", err)
	}
	fmt.Println("struct -> JSON:", string(b))
}
