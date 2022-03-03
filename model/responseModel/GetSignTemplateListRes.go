package responseModel

type SignTemplateListInfo struct {
	SignTemplateId     string `json:"signTemplateId,omitempty"`
	SignTemplateName   string `json:"signTemplateName,omitempty"`
	SignTemplateStatus string `json:"signTemplateStatus,omitempty"`
	CreateTime         string `json:"createTime,omitempty"`
	UpdateTime         string `json:"updateTime,omitempty"`
}

type SignTemListResData struct {
	SignTemplates []SignTemplateListInfo `json:"signTemplates,omitempty"`
}

type GetSignTemplateListRes struct {
	RequestId string             `json:"requestId,omitempty"`
	Code      string             `json:"code,omitempty"`
	Msg       string             `json:"msg,omitempty"`
	Data      SignTemListResData `json:"data,omitempty"`
}
