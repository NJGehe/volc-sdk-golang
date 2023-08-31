package vePfs

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

const (
	ServiceVersion20220101 = "2022-01-01"
	ServiceName            = "vepfs"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "open.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		// describe file system overview
		"DescribeFileSystemOverview": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeFileSystemOverview"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"DescribeFileSystemStatistics": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeFileSystemStatistics"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"DescribeRegions": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeRegions"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"DescribeZones": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeZones"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"DescribeFileSystems": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeFileSystems"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"CreateFileSystem": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateFileSystem"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"DeleteFileSystem": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteFileSystem"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"UpdateFileSystem": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateFileSystem"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"ExpandFileSystem": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ExpandFileSystem"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"DescribeMountPoints": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeMountPoints"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"CreateMountPoint": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateMountPoint"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"DeleteMountPoint": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteMountPoint"},
				"Version": []string{ServiceVersion20220101},
			},
		},
		"UpdateMountPoint": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateMountPoint"},
				"Version": []string{ServiceVersion20220101},
			},
		},
	}
)

// VePfs .
type VePfs struct {
	Client *base.Client
}

// NewInstance 创建一个实例，同时指定地域和服务
func NewInstance(regionId string, accessKey string, secretKey string) *VePfs {
	instance := &VePfs{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = regionId
	instance.Client.SetAccessKey(accessKey)
	instance.Client.SetSecretKey(secretKey)
	return instance
}

// GetServiceInfo interface
func (p *VePfs) GetServiceInfo() *base.ServiceInfo {
	return p.Client.ServiceInfo
}

// GetAPIInfo interface
func (p *VePfs) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetRegion .
func (p *VePfs) SetRegion(region string) {
	p.Client.ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *VePfs) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *VePfs) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}
