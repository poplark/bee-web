package common

type BaseReq struct {
	Action string
	Data   interface{}
}

type BaseResp struct {
	Action  string
	RetCode int
	Data    interface{}
}

type ListReq struct {
	BaseReq
	Offset int
	Limit  int
}

type ListResp struct {
	BaseResp
	Offset     int
	Limit      int
	TotalCount int64
}
