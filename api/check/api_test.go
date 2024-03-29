package check_test

import (
	"fmt"
	"os"
	"testing"

	tw "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

var (
	TC_secretId  = os.Getenv("tc_secretId")
	TC_secretKey = os.Getenv("tc_secretKey")
	Region       = "ap-guangzhou"
)

var client *tw.Client

// TestDescribeTagList 测试 DescribeTagList 方法是否正常工作


func TestDescribeTagList(t *testing.T) {
	// 这里使用真实的SecretId和SecretKey，但请注意保密性，最好使用测试账号的密钥

	// 调用 DescribeTagList 方法
	request := tw.NewDescribeTagListRequest()

	request.Limit = common.Uint64Ptr(100)
	request.Offset = common.Uint64Ptr(0)

	// 返回的resp是一个DescribeTagListResponse的实例，与请求对象对应
	response, err := client.DescribeTagList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	// fmt.Printf("%s", response.ToJsonString())
	// 遍历Data数组，打印每个标签的键和值
	for _, tag := range response.Response.Data {
		fmt.Printf("Tag: Key=%s,  UpdateTime=%s\n",
			*tag.TagKey, // 解引用获取字符串值
			*tag.UpdateTime) // 如果UpdateTime也是指针，同样需要解引用
	}
}


func init() {
	credential := common.NewCredential(TC_secretId, TC_secretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "billing.tencentcloudapi.com"

	// 创建 BillingClient 实例
	c, _ := tw.NewClient(credential, Region, cpf)
	client = c

}
