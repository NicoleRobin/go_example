package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// var jsonBlob = []byte(`{"result":0,"dcdn_progress":0,"message":"","root_url":"http://up057.tw11a.filemail.xunlei.com","uri":"request_upload","query_str":"g=22596363b3de40b06f981fb85d82312e8c0ed511&s=12&t=1525689963&ver=1&tid=c28ebf311c2bbe6878c07f98edf253ea&ui=150007900&e=1526294763&ms=10485760&ak=0:0:0:0&pk=filemail&aid=30d30bbe013e1caecc375328fbe2e238","block_size":0}`)
	var jsonBlob = []byte(`{"result":5,"dcdn_progress":2}`)
	type UploadRes struct {
		Result        int
		Dcdn_Progress int
		// Message      string
		// RootUrl      string
		// Uri          string
		// QueryStr     string
		// BlockSize    int
	}
	var res UploadRes
	err := json.Unmarshal(jsonBlob, &res)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", res)
}
