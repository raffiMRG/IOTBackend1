package Model

type BaseResponseModel struct {
	CodeResponse  int         `json:"Code"`
	HeaderMessage string      `json:"HeaderMessage"`
	Message       string      `json:"Message"`
	Data          interface{} `json:"Data"`
}
