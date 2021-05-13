package trace

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"sync"
)

const Header = "TRACE-ID"

var _ T = (*Trace)(nil)

type T interface {
	i()
	ID() string
	WithRequest(req *Request) *Trace
	WithResponse(resp *Response) *Trace
	AppendDialog(dialog *Dialog) *Trace
	AppendDebug(debug *Debug) *Trace
	AppendSQL(SQL *SQL) *Trace
	AppendRedis(redis *Redis) *Trace
	AppendGRPC(grpc *GRPC) *Trace
}

type Trace struct {
	mux                sync.Mutex
	Identifier         string    `json:"identifier"`
	Request            *Request  `json:"request"`
	Response           *Response `json:"response"`
	ThirdPartyRequests []*Dialog `json:"third_party_requests"`
	Debugs             []*Debug  `json:"debugs"`
	SQLs               []*SQL    `json:"sqls"`
	Redis              []*Redis  `json:"redis"`
	GRPCs              []*GRPC   `json:"grpcs"`
	Success            bool      `json:"success"`
	Cost               float64   `json:"cost"`
}

type Request struct {
	TTL        string      `json:"ttl"`
	Method     string      `json:"method"`
	DecodedURL string      `json:"decoded_url"`
	Header     interface{} `json:"header"`
	Body       interface{} `json:"body"`
}

type Response struct {
	Header      interface{} `json:"header"`       // 返回头信息
	Body        interface{} `json:"body"`         // 返回Body信息
	BusinessNo  int         `json:"business_no"`  // 业务编码
	BusinessMsg string      `json:"business_msg"` // 业务信息
	HttpCode    int         `json:"http_code"`    // HTTP状态码
	HttpMsg     string      `json:"http_msg"`     // HTTP状态信息
	Cost        float64     `json:"cost"`         // 耗时
}

func New(id string) *Trace {
	if id == "" {
		buf := make([]byte, 10)
		_, err := io.ReadFull(rand.Reader, buf)
		if err != nil {
			panic(err)
		}
		id = hex.EncodeToString(buf)
	}

	return &Trace{
		Identifier: id,
	}
}

func (t *Trace) i() {}

func (t *Trace) ID() string {
	return t.Identifier
}

func (t *Trace) WithRequest(req *Request) *Trace {
	t.Request = req
	return t
}

func (t *Trace) WithResponse(resp *Response) *Trace {
	t.Response = resp
	return t
}

func (t *Trace) AppendDialog(dialog *Dialog) *Trace {
	if dialog == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.ThirdPartyRequests = append(t.ThirdPartyRequests, dialog)
	return t
}

func (t *Trace) AppendDebug(debug *Debug) *Trace {
	if debug == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.Debugs = append(t.Debugs, debug)
	return t
}

func (t *Trace) AppendSQL(sql *SQL) *Trace {
	if sql == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.SQLs = append(t.SQLs, sql)
	return t
}

func (t *Trace) AppendRedis(redis *Redis) *Trace {
	if redis == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.Redis = append(t.Redis, redis)
	return t
}

func (t *Trace) AppendGRPC(grpc *GRPC) *Trace {
	if grpc == nil {
		return t
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.GRPCs = append(t.GRPCs, grpc)
	return t
}
