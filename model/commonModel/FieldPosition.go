package commonModel

type FieldPosition struct {
	PositionMode    string `json:"positionMode,omitempty"`
	PositionPageNo  int    `json:"positionPageNo,omitempty"`
	PositionX       string `json:"positionX"`
	PositionY       string `json:"positionY"`
	PositionKeyword string `json:"positionKeyword"`
}
