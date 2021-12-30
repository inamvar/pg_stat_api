package models


type SatementStat struct {
	Query string `json:"query"`
	Calls int64  `json:"calls"`
	TotalExecTime  float64 `json:"total_exec_time"`
}