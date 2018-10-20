package response

type JsonResult struct {
	Error int32       `json:"error"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}
