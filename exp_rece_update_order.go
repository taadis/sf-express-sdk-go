package sf

// 订单确认或取消请求
type ExpReceUpdateOrderRequest struct {
	// ...
}

// 订单确认或取消响应
type ExpReceUpdateOrderResponse struct {
	BaseResponse
	// ...
}

func (r *ExpReceUpdateOrderRequest) ApiServiceCode() string {
	return "EXP_RECE_UPDATE_ORDER"
}
