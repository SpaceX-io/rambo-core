package trace

import "google.golang.org/grpc/metadata"

type GRPC struct {
	TimeStamp string                 `json:"time_stamp"`
	Addr      string                 `json:"addr"`
	Method    string                 `json:"method"`
	Meta      metadata.MD            `json:"meta"`
	Request   map[string]interface{} `json:"request"`
	Response  map[string]interface{} `json:"response"`
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Cost      float64                `json:"cost"`
}
