package check

import (
	"fmt"

	tw "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

// 定义一个subledger label结构体,用于接收返回的数据
type SubLedgerLabel struct {
	RecordNum  int8
	Data      []struct {
		TagKey     string
		Status     string
		UpdateTime string
	}
	RequestId string
}

// GetProjectLabel 获取项目标签
func GetRst(region string) (rst SubLedgerLabel, error) {
	// 定义分页参数变量
	var (
		limit  uint64 = 100
		offset uint64 = 0
	)

	cpf := profile.NewClientProfile() // 创建profile实例
	cpf.HttpProfile.Endpoint = "billing.tencentcloudapi.com" // 设置endpoint
	client, _ := tw.NewClient(GetSecret(), region, cpf) // 创建client实例
	request := tw.NewDescribeTagListRequest() // 创建请求实例

	request.Offset = &offset // 设置分页参数
	request.Limit = &limit // 设置分页参数

	// 调用DescribeTagList方法
	response, err := client.DescribeTagList(request)
	// 检查是否返回了腾讯云SDK的错误
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Println("调用API出错：", err)
		return
	}

	// 将返回的数据转换为结构体
	var rst SubLedgerLabel
	rst.RecordNum = *response.Response.RecordNum // 记录数
	rst.RequestId = *response.Response.RequestId // 请求ID
	for _, tag := range response.Response.Data { // 遍历Data数组
		rst.Data = append(rst.Data, struct {
			TagKey     string
			Status     string
			UpdateTime string
		}{
			TagKey:     *tag.TagKey,
			Status:     *tag.Status,
			UpdateTime: *tag.UpdateTime,
		})
	}
	return rst, nil
}
