package responseModel

type UserEnableRes struct {
	RequestId string `json:"requestId"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
}
