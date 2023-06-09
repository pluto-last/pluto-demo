package constant

const (
	DisableStatus = "disable" // 禁用状态
	NormalStatus  = "normal"  // 正常状态
)

const (
	TaskStatusPlan      = "plan"          // 计划中
	TaskStatusRunning   = "running"       // 进行中
	TaskStatusTimeout   = "timeout"       // 超时
	TaskStatusSuccess   = "success"       //
	TaskTokenErr        = "token_err"     // 获取token失败
	TaskRespErr         = "resp_err"      // 接口返回失败
	TaskTokenTimeoutErr = "token_timeout" // deviceToken被风控
	TaskIPBlackErr      = "ip_black"      // ip被加入黑名单,限制访问了
	TaskProxyErr        = "proxy_err"     // ip被加入黑名单,限制访问了

)

const (
	ProxySwitch   = "proxy_switch"    // 代理选择
	ProxyTimeOutZ = "proxy_timeout:z" // 代理连续提取超时次数
	ProxyTimeOutK = "proxy_timeout:k" // 代理连续提取超时次数
	PyreqErrCount = "pyreq_err_count" // python程序异常
)

const (
	ProxyMaxUse = 8
)
