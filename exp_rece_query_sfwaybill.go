package sf

// 清单运费查询请求
type ExpReceQuerySfwaybillRequest struct {
	TrackingType    string `json:"trackingType"`    // 1:表示按订单查询
	TrackingNum     string `json:"trackingNum"`     // 订单号
	BizTemplateCode string `json:"bizTemplateCode"` // 业务配置代码
}

// 清单运费查询响应
type ExpReceQuerySfwaybillResponse struct {
	BaseResponse
	// ...
}

func (r *ExpReceQuerySfwaybillRequest) ApiServiceCode() string {
	return "EXP_RECE_QUERY_SFWAYBILL"
}
