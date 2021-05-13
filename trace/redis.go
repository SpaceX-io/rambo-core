package trace

type Redis struct {
	TimeStamp string  `json:"time_stamp"`
	Handle    string  `json:"handle"`
	Key       string  `json:"key"`
	Value     string  `json:"value"`
	TTL       float64 `json:"ttl"`
	Cost      float64 `json:"cost"`
}
