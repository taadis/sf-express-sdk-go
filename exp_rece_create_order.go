package sf

// 下订单请求
type ExpReceCreateOrderRequest struct {
	// ...
}

// 下订单响应
type ExpReceCreateOrderResponse struct {
	BaseResponse
	// ...
}

func (r *ExpReceCreateOrderRequest) ApiServiceCode() string {
	return "EXP_RECE_CREATE_ORDER"
}
