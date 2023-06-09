package params

type UploadReq struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    string       `json:"data"`
	Param   UploadParams `json:"param"`
}

type UploadParams struct {
	Env          string `json:"__env"`
	TraceID      string `json:"__traceId"`
	JobID        int    `json:"__jobId"`
	CrawlerType  int    `json:"__crawlerType"`
	TaskID       int    `json:"__taskId"`
	TaskKey      string `json:"__taskKey"`
	GroupID      int    `json:"__groupId"`
	SequenceID   int    `json:"__sequenceId"`
	BusinessID   int    `json:"__businessId"`
	FunctionID   int    `json:"__functionId"`
	FunctionName string `json:"__functionName"`
	CompanyID    string `json:"__companyId"`
	Source       string `json:"__source"`
	CompressTag  string `json:"_compress_tag"`
	ISLogin      bool   `json:"is_login"`
}

type UploadReqResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
