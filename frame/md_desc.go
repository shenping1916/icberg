package frame

// Medesc 方法描述
// 认证方式
// 流量统计
// 失败统计
type Medesc struct {
	MdName  string        `json:"md_name"`
	A       Authorization `json:"A"`
	FailCnt int64         `json:"fail_cnt"`
	Cnt     int64         `json:"cnt"`
}
