package main

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
)

const AK = "c1e0864ee3a84c72ab70c4ee5c0df9d7"
const SK = "377c19eea2c24132b57dba53ea50e630"
const ENDPOINT = "https://cdn.baidubce.com"

func main() {
	method := "PUT"
	path := "/v2/domain/nas.it927.com/config"

	headers := map[string]string{"Content-Type": "application/json"}
	params := map[string]string{"origin": ""}
	payload, _ := bce.NewBodyFromString("{\"origin\":[{\"peer\":\"[2409:8a50:819:c780:20c:29ff:fe7f:2e93]\",\"host\":\"nas.it927.com\",\"backup\":false,\"weight\":1}],\"follow302\":true}")

	req := &bce.BceRequest{}
	req.SetUri(path)
	req.SetMethod(method)
	req.SetParams(params)
	req.SetHeaders(headers)
	req.SetBody(payload)

	client, err := bce.NewBceClientWithAkSk(AK, SK, ENDPOINT)
	if err != nil {
		// client 初始化失败终止启动
		panic(err)
	}
	resp := &bce.BceResponse{}
	if err := client.SendRequest(req, resp); err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("http code: ", resp.StatusCode())
		respBody := make(map[string]any)
		if err := resp.ParseJsonBody(&respBody); err == nil {
			fmt.Println("resp:", respBody)
		} else {
			fmt.Println("data is null: ", err)
		}
	}
}
