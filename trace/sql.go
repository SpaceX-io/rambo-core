package trace

type SQL struct {
	TimeStamp string  `json:"time_stamp"`
	Stack     string  `json:"stack"`
	SQL       string  `json:"sql"`
	Rows      int64   `json:"rows"`
	Cost      float64 `json:"cost"`
}
