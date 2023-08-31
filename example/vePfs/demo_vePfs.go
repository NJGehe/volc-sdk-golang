package main

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vePfs"
	"net/http"
)

const (
	testAk   = "testAk"
	testSk   = "testSk"
	regionId = "cn-beijing"
)

func main() {
	vePfsClient := vePfs.NewInstance(regionId, testAk, testSk)
	vePfsClient.Client.SetRetrySettings(&base.RetrySettings{AutoRetry: true})

	var tags []*vePfs.Tag
	var tagOwner = vePfs.Tag{"owner", "sre", "Custom"}
	tags = append(tags, &tagOwner)
	req := &vePfs.CreateFileSystemReq{
		FileSystemName: "test-sre",
		ZoneId:         "cn-beijing-b",
		ChargeType:     "PayAsYouGo",
		FileSystemType: "VePFS",
		StoreType:      "Advance_100",
		ProtocolType:   "VePFS",
		Tags:           tags,
		Capacity:       10,
		Description:    "just_for_test",
	}
	resp, statusCode, err := vePfsClient.CreateFileSystem(
		req)
	if err != nil {
		fmt.Println("statusCode:", statusCode, "error:", err)
		return
	}
	if statusCode != http.StatusOK {
		fmt.Println("statusCode:", statusCode)
		return
	}
	r, _ := json.Marshal(resp)
	fmt.Println(string(r))
}
