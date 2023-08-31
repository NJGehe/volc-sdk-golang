package vePfs

import (
	"encoding/json"
	"net/url"
)

func (p *VePfs) DescribeFileSystemOverview(req *DescribeFileSystemOverviewReq) (*DescribeFileSystemOverviewResp, int, error) {
	query := url.Values{}
	resp := new(DescribeFileSystemOverviewResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("DescribeFileSystemOverview", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) DescribeFileSystemStatistics(req *DescribeFileSystemStatisticsReq) (*DescribeFileSystemStatisticsResp, int, error) {
	query := url.Values{}
	resp := new(DescribeFileSystemStatisticsResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("DescribeFileSystemStatistics", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) DescribeRegions(req *DescribeRegionsReq) (*DescribeRegionsResp, int, error) {
	query := url.Values{}
	resp := new(DescribeRegionsResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("DescribeRegions", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) DescribeZones(req *DescribeZonesReq) (*DescribeZonesResp, int, error) {
	query := url.Values{}
	resp := new(DescribeZonesResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("DescribeZones", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) DescribeFileSystems(req *DescribeFileSystemsReq) (*DescribeFileSystemsResp, int, error) {
	query := url.Values{}
	resp := new(DescribeFileSystemsResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("DescribeFileSystems", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) CreateFileSystem(req *CreateFileSystemReq) (*CreateFileSystemResp, int, error) {
	query := url.Values{}
	resp := new(CreateFileSystemResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("CreateFileSystem", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) DeleteFileSystem(req *DeleteFileSystemReq) (*DeleteFileSystemResp, int, error) {
	query := url.Values{}
	resp := new(DeleteFileSystemResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("DeleteFileSystem", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) UpdateFileSystem(req *UpdateFileSystemReq) (*UpdateFileSystemResp, int, error) {
	query := url.Values{}
	resp := new(UpdateFileSystemResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("UpdateFileSystem", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) ExpandFileSystem(req *ExpandFileSystemReq) (*ExpandFileSystemResp, int, error) {
	query := url.Values{}
	resp := new(ExpandFileSystemResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("ExpandFileSystem", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) DescribeMountPoints(req *DescribeMountPointsReq) (*DescribeMountPointsResp, int, error) {
	query := url.Values{}
	resp := new(DescribeMountPointsResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("DescribeMountPoints", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) CreateMountPoint(req *CreateMountPointReq) (*CreateMountPointResp, int, error) {
	query := url.Values{}
	resp := new(CreateMountPointResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("CreateMountPoint", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) DeleteMountPoint(req *DeleteMountPointReq) (*DeleteMountPointResp, int, error) {
	query := url.Values{}
	resp := new(DeleteMountPointResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("DeleteMountPoint", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) UpdateMountPoint(req *UpdateMountPointReq) (*UpdateMountPointResp, int, error) {
	query := url.Values{}
	resp := new(UpdateMountPointResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("UpdateMountPoint", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *VePfs) commonHandlerJson(api string, query url.Values, reqBody []byte, resp interface{}) (statusCode int, err error) {
	var respBody []byte

	respBody, statusCode, err = p.Client.Json(api, query, string(reqBody))
	if err != nil {
		return
	}

	if err = json.Unmarshal(respBody, resp); err != nil {
		return
	}
	return
}
