package trace

import "sync"

var _ D = (*Dialog)(nil)

type D interface {
	i()
	AppendResp(resp *Response)
}

type Dialog struct {
	mux       sync.Mutex
	Request   *Request    `json:"request"`
	Responses []*Response `json:"response"`
	Success   bool        `json:"success"`
	Cost      float64     `json:"cost"`
}

func (d *Dialog) i() {}

func (d *Dialog) AppendResp(resp *Response) {
	if resp == nil {
		return
	}

	d.mux.Lock()

	d.Responses = append(d.Responses, resp)
	d.mux.Unlock()
}
