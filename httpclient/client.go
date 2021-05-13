package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	netHttp "net/url"
	"time"

	"github.com/SpaceX-io/rambo-core/trace"
	"github.com/pkg/errors"
)

const DefaultTTL = time.Minute

func Get(url string, form netHttp.Values, options ...Option) (body []byte, err error) {
	return withoutBody(http.MethodGet, url, form, options...)
}

func Delete(url string, form netHttp.Values, options ...Option) (body []byte, err error) {
	return withoutBody(http.MethodDelete, url, form, options...)
}

func PostForm(url string, form netHttp.Values, options ...Option) (body []byte, err error) {
	return withFormBody(http.MethodPost, url, form, options...)
}

func PostJSON(url string, raw json.RawMessage, options ...Option) (body []byte, err error) {
	return withJSONBody(http.MethodPost, url, raw, options...)
}

func PutForm(url string, form netHttp.Values, options ...Option) (body []byte, err error) {
	return withFormBody(http.MethodPut, url, form, options...)
}

func PutJSON(url string, raw json.RawMessage, options ...Option) (body []byte, err error) {
	return withJSONBody(http.MethodPut, url, raw, options...)
}

func PatchForm(url string, form netHttp.Values, options ...Option) (body []byte, err error) {
	return withFormBody(http.MethodPatch, url, form, options...)
}

func PathJSON(url string, raw json.RawMessage, options ...Option) (body []byte, err error) {
	return withJSONBody(http.MethodPatch, url, raw, options...)
}

func withoutBody(method, url string, form netHttp.Values, options ...Option) (body []byte, err error) {
	if url == "" {
		return nil, errors.New("url required")
	}
	if len(form) > 0 {
		if url, err = addFormValuesIntoURL(url, form); err != nil {
			return
		}
	}

	ts := time.Now()

	o := getOption()
	defer func() {
		if o.trace != nil {
			o.dialog.Success = err == nil
			o.dialog.Cost = time.Since(ts).Seconds()
			o.trace.AppendDialog(o.dialog)
		}

		releaseOption(o)
	}()

	for _, f := range options {
		f(o)
	}
	o.header["Content-Type"] = []string{"application/x-www-form-urlencoded;charset=utf-8"}
	if o.trace != nil {
		o.header[trace.Header] = []string{o.trace.ID()}
	}

	ttl := o.ttl
	if ttl <= 0 {
		ttl = DefaultTTL
	}

	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()

	if o.dialog != nil {
		decodedURL, _ := netHttp.QueryUnescape(url)
		o.dialog.Request = &trace.Request{
			TTL:        ttl.String(),
			Method:     method,
			DecodedURL: decodedURL,
			Header:     o.header,
		}
	}

	retryTimes := o.retryTimes
	if retryTimes <= 0 {
		retryTimes = DefaultRetryTimes
	}

	retryDelay := o.retryDelay
	if retryDelay <= 0 {
		retryDelay = DefaultRetryDelay
	}

	var httpCode int

	defer func() {
		if o.alarmObj == nil {
			return
		}

		if o.alarmVerify != nil && !o.alarmVerify(body) && err == nil {
			return
		}

		info := &struct {
			TraceID string `json:"trace_id"`
			Request struct {
				Method string `json:"method"`
				URL    string `json:"url"`
			} `json:"request"`
			Response struct {
				HttpCode int    `json:"http_code"`
				Body     string `json:"body"`
			} `json:"response"`
			Error string `json:"error"`
		}{}

		if o.trace != nil {
			info.TraceID = o.trace.ID()
		}
		info.Request.Method = method
		info.Request.URL = url
		info.Response.HttpCode = httpCode
		info.Response.Body = string(body)
		info.Error = ""
		if err != nil {
			info.Error = fmt.Sprintf("%+v", err)
		}

		raw, _ := json.MarshalIndent(info, "", " ")
		onFailedAlarm(o.alarmTitle, raw, o.logger, o.alarmObj)
	}()

	for k := 0; k < retryTimes; k++ {
		body, httpCode, err = doHTTP(ctx, method, url, nil, o)
		if shouldRetry(ctx, httpCode) || (o.retryVerify != nil && o.retryVerify(body)) {
			time.Sleep(retryDelay)
			continue
		}

		return
	}

	return
}

func withFormBody(method, url string, form netHttp.Values, options ...Option) (body []byte, err error) {
	if url == "" {
		return nil, errors.New("url required")
	}
	if len(form) == 0 {
		return nil, errors.New("form required")
	}

	ts := time.Now()

	opt := getOption()
	defer func() {
		if opt.trace != nil {
			opt.dialog.Success = err == nil
			opt.dialog.Cost = time.Since(ts).Seconds()
			opt.trace.AppendDialog(opt.dialog)
		}

		releaseOption(opt)
	}()

	for _, f := range options {
		f(opt)
	}

	opt.header["Context-Type"] = []string{"application/x-www-form-urlencoded; charset=utf-8"}
	if opt.trace != nil {
		opt.header[trace.Header] = []string{opt.trace.ID()}
	}

	ttl := opt.ttl
	if ttl <= 0 {
		ttl = DefaultTTL
	}

	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()

	formValue := form.Encode()
	if opt.dialog != nil {
		decodedURL, _ := netHttp.QueryUnescape(url)
		opt.dialog.Request = &trace.Request{
			TTL:        ttl.String(),
			Method:     method,
			DecodedURL: decodedURL,
			Header:     opt.header,
			Body:       formValue,
		}
	}

	retryTimes := opt.retryTimes
	if retryTimes <= 0 {
		retryTimes = DefaultRetryTimes
	}

	retryDelay := opt.retryDelay
	if retryDelay <= 0 {
		retryDelay = DefaultRetryDelay
	}

	var httpCode int

	defer func() {
		if opt.alarmObj == nil {
			return
		}

		if opt.alarmVerify != nil && !opt.alarmVerify(body) && err == nil {
			return
		}

		info := &struct {
			TraceID string `json:"trace_id"`
			Request struct {
				Method string `json:"method"`
				URL    string `json:"url"`
			} `json:"request"`
			Response struct {
				HttpCode int    `json:"http_code"`
				Body     string `json:"body"`
			} `json:"response"`
			Error string `json:"error"`
		}{}

		if opt.trace != nil {
			info.TraceID = opt.trace.ID()
		}
		info.Request.Method = method
		info.Request.URL = url
		info.Response.HttpCode = httpCode
		info.Response.Body = string(body)
		info.Error = ""
		if err != nil {
			info.Error = fmt.Sprintf("%+v", err)
		}

		raw, _ := json.MarshalIndent(info, "", " ")
		onFailedAlarm(opt.alarmTitle, raw, opt.logger, opt.alarmObj)
	}()

	for k := 0; k < retryTimes; k++ {
		body, httpCode, err = doHTTP(ctx, method, url, []byte(formValue), opt)
		if shouldRetry(ctx, httpCode) || (opt.retryVerify != nil && opt.retryVerify(body)) {
			time.Sleep(retryDelay)
			continue
		}
		return
	}
	return
}

func withJSONBody(method, url string, raw json.RawMessage, options ...Option) (body []byte, err error) {
	if url == "" {
		return nil, errors.New("url required")
	}
	if len(raw) == 0 {
		return nil, errors.New("raw required")
	}

	ts := time.Now()

	opt := getOption()
	defer func() {
		if opt.trace != nil {
			opt.dialog.Success = err == nil
			opt.dialog.Cost = time.Since(ts).Seconds()
			opt.trace.AppendDialog(opt.dialog)
		}

		releaseOption(opt)
	}()

	for _, f := range options {
		f(opt)
	}

	opt.header["Content-Type"] = []string{"application/json; charset=utf-8"}
	if opt.trace != nil {
		opt.header[trace.Header] = []string{opt.trace.ID()}
	}

	ttl := opt.ttl
	if ttl <= 0 {
		ttl = DefaultTTL
	}

	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()

	if opt.dialog != nil {
		decodedURL, _ := netHttp.QueryUnescape(url)
		opt.dialog.Request = &trace.Request{
			TTL:        ttl.String(),
			Method:     method,
			DecodedURL: decodedURL,
			Header:     opt.header,
			Body:       string(raw),
		}
	}

	retryTimes := opt.retryTimes
	if retryTimes <= 0 {
		retryTimes = DefaultRetryTimes
	}

	retryDelay := opt.retryDelay
	if retryDelay <= 0 {
		retryDelay = DefaultRetryDelay
	}

	var httpCode int

	defer func() {
		if opt.alarmObj == nil {
			return
		}

		if opt.alarmVerify != nil && !opt.alarmVerify(body) && err == nil {
			return
		}

		info := &struct {
			TraceID string `json:"trace_id"`
			Request struct {
				Method string `json:"method"`
				URL    string `json:"url"`
			} `json:"request"`
			Response struct {
				HTTPCode int    `json:"http_code"`
				Body     string `json:"body"`
			} `json:"response"`
			Error string `json:"error"`
		}{}

		if opt.trace != nil {
			info.TraceID = opt.trace.ID()
		}
		info.Request.Method = method
		info.Request.URL = url
		info.Response.HTTPCode = httpCode
		info.Response.Body = string(body)
		info.Error = ""
		if err != nil {
			info.Error = fmt.Sprintf("%+v", err)
		}

		raw, _ := json.MarshalIndent(info, "", " ")
		onFailedAlarm(opt.alarmTitle, raw, opt.logger, opt.alarmObj)
	}()

	for k := 0; k < retryTimes; k++ {
		body, httpCode, err = doHTTP(ctx, method, url, raw, opt)
		if shouldRetry(ctx, httpCode) || (opt.retryVerify != nil && opt.retryVerify(body)) {
			time.Sleep(retryDelay)
			continue
		}
		return
	}

	return
}
