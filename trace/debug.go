package trace

type Debug struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Cost  float64     `json:"cost"`
}
