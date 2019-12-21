package sdk

import (
	"flag"
	"strconv"
	"strings"
	"testing"
	"time"
)

var (
	clientCode = flag.String("clientCode", "", "顾客编码")
	checkWord  = flag.String("checkWord", "", "校检码")
)

// 解析命令行参数
func parseArgs(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}
	t.Logf("-args -checkWord %s -checkWord %s", *clientCode, *checkWord)
}

// 测试新客户端
func TestNewClient(t *testing.T) {
	parseArgs(t)
	xmlRequest := `
<Request service="OrderService" lang="zh-CN">
  <Head>clientCode</Head>
  <Body>
    <Order orderid="orderId" express_type="1" j_province="广东省" j_city="深圳市" j_company="顺丰镖局" j_contact="艾丽丝" j_tel="15012345678" j_address="福田区新洲十一街万基商务大厦26楼" d_province="广东省" d_city="广州市" d_county="" d_company="神罗科技" d_contact="风一样的旭哥" d_tel="33992159" d_address="海珠区宝芝林大厦701室" parcel_quantity="1" pay_method="3" custid="7551234567" customs_batchs="" cargo="iphone7plus">
      <AddedService name="COD" value="1.01" value1="7551234567"/>
    </Order>
  </Body>
</Request>
`
	client, err := NewClient()
	if err != nil {
		t.Error(err)
	}
	orderId := strconv.FormatInt(time.Now().Unix(), 10)
	xmlRequest = strings.ReplaceAll(xmlRequest, "clientCode", *clientCode)
	xmlRequest = strings.ReplaceAll(xmlRequest, "orderId", orderId)
	res, err := client.ExpressService(xmlRequest, *checkWord)
	if err != nil {
		t.Fail()
	}
	t.Log(res)
	if !strings.Contains(res, "OK") {
		t.Fail()
	}
}

// 测试生成认证码函数
func TestComputeVerifyCode(t *testing.T) {
	parseArgs(t)
	verifyCode := computeVerifyCode("123456")
	if verifyCode != "4QrcOUm6Wau+VuBX8g+IPg==" {
		t.Fail()
	}
}
