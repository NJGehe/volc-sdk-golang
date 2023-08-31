package vePfs

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"go/types"
)

type NullResultResp struct {
	ResponseMetadata *base.ResponseMetadata
}

// models for DescribeFileSystemOverview API start

type DescribeFileSystemOverviewReq struct {
}

type DescribeFileSystemOverviewResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *FileSystemOverviewResult `json:",omitempty"`
}

type FileSystemOverviewResult struct {
	OverView *FileSystemOverView `json:",omitempty"`
}

type FileSystemOverView struct {
	AccountId    string
	ErrorCount   int32
	OtherCount   int32
	RegionId     string
	RunningCount int32
	TotalCount   int32
}

// models for DescribeFileSystemOverview API end

// models for DescribeFileSystemStatistics API start

type DescribeFileSystemStatisticsReq struct {
}

type DescribeFileSystemStatisticsResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *FileSystemStatisticsResult `json:",omitempty"`
}

type FileSystemStatisticsResult struct {
	Statistics []*FileSystemStatistic
}

type FileSystemStatistic struct {
	AccountId      string
	CapacityInfo   *CapacityInfo `json:",omitempty"`
	FileSystemType string
	RegionId       string
	StoreType      string
	TotalCount     int32
}

type CapacityInfo struct {
	TotalTiB int32
	UsedGiB  int32
}

// models for DescribeFileSystemStatistics API end

// models for DescribeRegions API start

type DescribeRegionsReq struct {
}

type DescribeRegionsResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *RegionsResult `json:",omitempty"`
}

type RegionsResult struct {
	Regions    []*Region
	TotalCount int32
}

type Region struct {
	RegionId   string
	RegionName string
}

// models for DescribeRegions API end

// models for DescribeZones API start

type DescribeZonesReq struct {
	RegionId string
}

type DescribeZonesResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ZonesResult `json:",omitempty"`
}

type ZonesResult struct {
	TotalCount int32
	Zones      []*ZoneInfo
}

type ZoneInfo struct {
	RegionId  string
	ZoneId    string
	ZoneName  string
	SaleInfos []*SaleInfos
}

type SaleInfos struct {
	FileSystemType string
	ProtocolType   string
	RegionId       string
	Status         string
	StoreType      string
	ZoneId         string
}

// models for DescribeZones API end

// models for DescribeFileSystems API start

type DescribeFileSystemsReq struct {
	PageSize      int32
	PageNumber    int32
	OrderBy       string
	Filters       []*FileSystemFilter
	Project       string
	FileSystemIds string
}
type FileSystemFilter struct {
	Key   string
	Value string
}

type DescribeFileSystemsResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *FileSystemsResult `json:",omitempty"`
}

type FileSystemsResult struct {
	FileSystems []*FileSystem
}

type FileSystem struct {
	AccountId       string
	AutoRenew       bool
	Bandwidth       int32
	CapacityInfo    *CapacityInfo `json:",omitempty"`
	ChargeStatus    string
	ChargeType      string
	CreateTime      string
	Description     string
	ExpireTime      string
	FileSystemId    string
	FileSystemName  string
	FileSystemType  string
	FreeTime        string
	LastModifyTime  string
	Month           int32
	Project         string
	ProtocolType    string
	RegionId        string
	ReplicasNum     int32
	Status          string
	StopServiceTime string
	Storage         *Storage `json:",omitempty"`
	StoreType       string
	Tags            []*Tag
	Version         string
	ZoneId          string
	ZoneName        string
}

type Tag struct {
	Key   string
	Value string
	Type  string
}

type Storage struct {
}

// models for DescribeFileSystems API end

// models for CreateFileSystem API start

type CreateFileSystemReq struct {
	FileSystemName string
	ZoneId         string
	ChargeType     string
	FileSystemType string
	StoreType      string
	ProtocolType   string
	Description    string
	Project        string
	Capacity       int32
	Tags           []*Tag
}

type CreateFileSystemResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *FileSystemsCreationResult `json:",omitempty"`
}

type FileSystemsCreationResult struct {
	FileSystemId string
	OrderNO      string
	Tags         []*Tag
}

type DeleteFileSystemReq struct {
	FileSystemId string
}

type DeleteFileSystemResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           types.Nil
}

// models for CreateFileSystem API end

// models for UpdateFileSystem API start

type UpdateFileSystemReq struct {
	FileSystemId   string
	FileSystemName string
	Description    string
	Tags           []*Tag
}

type UpdateFileSystemResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           types.Nil
}

// models for UpdateFileSystem API end

// models for ExpandFileSystem API start

type ExpandFileSystemReq struct {
	FileSystemId   string
	Capacity       int32
	EnableRestripe bool
}

type ExpandFileSystemResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *ExpendFileSystemResult `json:",omitempty"`
}
type ExpendFileSystemResult struct {
	OrderNO string
}

// models for ExpandFileSystem API end

// models for DescribeMountPoints API start

type DescribeMountPointsReq struct {
	FileSystemId   string `json:"file_system_id,omitempty"`
	MountPointName string `json:"mount_point_name,omitempty"`
	MountPointId   string `json:"mount_point_id,omitempty"`
}

type DescribeMountPointsResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *MountPointsResult `json:",omitempty"`
}

type MountPointsResult struct {
	MountPoints []*MountPoint
	TotalCount  int32
}

type MountPoint struct {
	CreateTime     string
	FileSystemId   string
	MountPointId   string
	MountPointName string
	Nodes          []*Node
	Status         string
	SubnetId       string
	SubnetName     string
	VpcId          string
	VpcName        string
}

type Node struct {
	DefaultPassword string
	EcsIP           string
	EcsId           string
}

// models for DescribeMountPoints API end

// models for CreateMountPoint API start

type CreateMountPointReq struct {
	FileSystemId   string
	MountPointName string
	VpcId          string
	SubnetId       string
}

type CreateMountPointResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *CreateMountPointsResult `json:",omitempty"`
}

type CreateMountPointsResult struct {
	MountPointId string
}

// models for CreateMountPoint API end

// models for DeleteMountPoint API start

type DeleteMountPointReq struct {
	FileSystemId string
	MountPointId string
}

type DeleteMountPointResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *DeleteMountPointsResult `json:",omitempty"`
}

type DeleteMountPointsResult struct {
}

// models for UpdateMountPoint API start

type UpdateMountPointReq struct {
	FileSystemId   string
	MountPointId   string
	MountPointName string
}

type UpdateMountPointResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *UpdateMountPointsResult `json:",omitempty"`
}

type UpdateMountPointsResult struct {
}

// models for UpdateMountPoint API end
