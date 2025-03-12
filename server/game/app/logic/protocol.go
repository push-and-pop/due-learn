package logic

type GreetReq struct {
	Message string `json:"message"`
}

type GreetRes struct {
	Message string `json:"message"`
}
