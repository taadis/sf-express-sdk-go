package sf

// 订单结果查询请求
type ExpReceSearchOrderRespRequest struct {
	OrderId    string `json:"orderId"`    // 客户订单号
	SearchType string `json:"searchType"` // 查询类型
	Language   string `json:"language"`   // 语言(返回描述语)
}

// 订单结果查询响应
type ExpReceSearchOrderRespResponse struct {
	BaseResponse
	// ...
}

func (r *ExpReceSearchOrderRespRequest) ApiServiceCode() string {
	return "EXP_RECE_SEARCH_ORDER_RESP"
}
