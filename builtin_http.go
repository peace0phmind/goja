package goja

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

//go:generate ag

/*
HC

	@Enum {
		url
		method
		baseURL
		headers
		params
		data
		timeout
		responseType
		responseEncoding
	}
*/
type HC string

type httpResponse struct {
	Status     int
	StatusText string
	Headers    Value
	Data       Value
}

type HttpConfig struct {
	Url              string
	Method           string
	BaseURL          string
	Headers          map[string]any
	Params           map[string]any
	Data             []byte
	Timeout          int
	ResponseType     string
	ResponseEncoding string
}

func (hc *HttpConfig) CheckError() error {
	if len(hc.Url) == 0 {
		return errors.New("missing url")
	}

	hc.Method = strings.ToUpper(hc.Method)
	switch hc.Method {
	case http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
	case http.MethodConnect, http.MethodOptions, http.MethodTrace:
	default:
		return errors.New("invalid method")
	}

	hc.ResponseType = strings.ToLower(hc.ResponseType)
	switch hc.ResponseType {
	case "json", "text":
	default:
		return errors.New("invalid response type, only: json, text")
	}

	return nil
}

func (hc *HttpConfig) getUrl() (string, error) {
	if len(hc.BaseURL) == 0 {
		parsedPath, err := url.Parse(hc.Url)
		if err != nil {
			fmt.Println("Error parsing path:", err)
			return "", err
		}

		return parsedPath.String(), nil
	} else {
		parsedBase, err := url.Parse(hc.BaseURL)
		if err != nil {
			fmt.Println("Error parsing base URL:", err)
			return "", err
		}

		// 使用 ResolveReference 来处理路径
		parsedPath, err := url.Parse(hc.Url)
		if err != nil {
			fmt.Println("Error parsing path:", err)
			return "", err
		}

		// 解析后的 URL
		finalURL := parsedBase.ResolveReference(parsedPath)
		return finalURL.String(), nil
	}
}

var defaultHttpConfig = HttpConfig{
	Method:           http.MethodGet,
	Timeout:          0,
	ResponseType:     "json",
	ResponseEncoding: "utf8",
}

func (r *Runtime) http_getHttpConfig(argValue Value) HttpConfig {
	conf := defaultHttpConfig

	if confArg := argValue; confArg != _undefined {
		confObj := confArg.ToObject(r)

		if urlStr := confObj.Get(HCUrl.Val()); urlStr != nil {
			conf.Url = urlStr.String()
		}

		if method := confObj.Get(HCMethod.Val()); method != nil {
			conf.Method = method.String()
		}

		if baseURL := confObj.Get(HCBaseUrl.Val()); baseURL != nil {
			conf.BaseURL = baseURL.String()
		}

		if headers := confObj.Get(HCHeaders.Val()); headers != nil {
			conf.Headers = headers.Export().(map[string]any)
		}

		if params := confObj.Get(HCParams.Val()); params != nil {
			conf.Params = params.Export().(map[string]any)
		}

		if data := confObj.Get(HCData.Val()); data != nil {
			switch dataType := data.(type) {
			case *Object:

			default:
				fmt.Printf("%v", dataType.String())
			}
			conf.Data = nil
		}

		if timeout := confObj.Get(HCTimeout.Val()); timeout != nil {
			conf.Timeout = int(timeout.ToInteger())
		}

		if responseType := confObj.Get(HCResponseType.Val()); responseType != nil {
			conf.ResponseType = responseType.String()
		}

		if responseEncoding := confObj.Get(HCResponseEncoding.Val()); responseEncoding != nil {
			conf.ResponseEncoding = responseEncoding.String()
		}
	}

	return conf
}

func (r *Runtime) http_responseToHttpResponse(response *http.Response, conf HttpConfig) *httpResponse {
	resp := &httpResponse{}
	resp.Status = response.StatusCode
	resp.StatusText = response.Status

	headerMap := make(map[string]string)
	for k, v := range response.Header {
		headerMap[k] = strings.Join(v, ",")
	}
	resp.Headers = r.ToValue(headerMap)

	buf, err := io.ReadAll(response.Body)
	if err != nil {
		// log out err
		return nil
	}

	switch strings.ToLower(conf.ResponseType) {
	case "json":
		d := json.NewDecoder(strings.NewReader(string(buf)))
		resp.Data, err = r.builtinJSON_decodeValue(d)
		if err != nil {
			return nil
		}
	case "text":
		resp.Data = asciiString(buf)
	}

	return resp
}

func (r *Runtime) _http_request(conf HttpConfig) Value {
	p, resolve, reject := r.NewPromise()

	if err := conf.CheckError(); err != nil {
		reject(err)
		return r.ToValue(p)
	}

	baseURL, err := conf.getUrl()
	if err != nil {
		reject(err)
		return r.ToValue(p)
	}

	conf.BaseURL = baseURL

	u, err := url.Parse(baseURL)
	if err != nil {
		reject(err)
		return r.ToValue(p)
	}

	values := u.Query()
	for k, v := range conf.Params {
		values.Set(k, fmt.Sprintf("%v", v))
	}

	u.RawQuery = values.Encode()

	req, err := http.NewRequest(conf.Method, u.String(), bytes.NewReader(conf.Data))
	if err != nil {
		reject(err)
		return r.ToValue(p)
	}

	req.Header.Set("Content-Type", "application/"+conf.ResponseType)
	for k, v := range conf.Headers {
		req.Header.Set(k, fmt.Sprintf("%v", v))
	}

	client := &http.Client{
		Timeout: time.Duration(conf.Timeout) * time.Millisecond,
	}

	response, err := client.Do(req)
	if err != nil {
		reject(err)
		return r.ToValue(p)
	}

	resp := r.http_responseToHttpResponse(response, conf)
	if resp == nil {
		reject(err)
		return r.ToValue(p)
	}

	resolve(resp)
	return r.ToValue(p)
}

func (r *Runtime) http_request(call FunctionCall) Value {
	conf := r.http_getHttpConfig(call.Argument(0))
	return r._http_request(conf)
}

func (r *Runtime) http_get(call FunctionCall) Value {
	urlStr := call.Argument(0).String()
	conf := r.http_getHttpConfig(call.Argument(1))
	conf.Url = urlStr
	conf.Method = http.MethodGet

	return r._http_request(conf)
}

func (r *Runtime) http_delete(call FunctionCall) Value {
	urlStr := call.Argument(0).String()
	conf := r.http_getHttpConfig(call.Argument(1))
	conf.Url = urlStr
	conf.Method = http.MethodDelete

	return r._http_request(conf)
}

func (r *Runtime) http_head(call FunctionCall) Value {
	urlStr := call.Argument(0).String()
	conf := r.http_getHttpConfig(call.Argument(1))
	conf.Url = urlStr
	conf.Method = http.MethodHead

	return r._http_request(conf)
}

func (r *Runtime) http_options(call FunctionCall) Value {
	urlStr := call.Argument(0).String()
	conf := r.http_getHttpConfig(call.Argument(1))
	conf.Url = urlStr
	conf.Method = http.MethodOptions

	return r._http_request(conf)
}

func (r *Runtime) http_post(call FunctionCall) Value {
	urlStr := call.Argument(0).String()
	//data := call.Argument(1).ToObject(r)
	conf := r.http_getHttpConfig(call.Argument(2))
	conf.Url = urlStr
	conf.Method = http.MethodPost
	//conf.Data = data

	return r._http_request(conf)
}

func (r *Runtime) http_put(call FunctionCall) Value {
	urlStr := call.Argument(0).String()
	//data := call.Argument(1).ToObject(r)
	conf := r.http_getHttpConfig(call.Argument(2))
	conf.Url = urlStr
	conf.Method = http.MethodPut
	//conf.Data = data

	return r._http_request(conf)
}

func (r *Runtime) http_patch(call FunctionCall) Value {
	urlStr := call.Argument(0).String()
	//data := call.Argument(1).ToObject(r)
	conf := r.http_getHttpConfig(call.Argument(2))
	conf.Url = urlStr
	conf.Method = http.MethodPatch
	//conf.Data = data

	return r._http_request(conf)
}

func createHttpTemplate() *objectTemplate {
	t := newObjectTemplate()
	t.protoFactory = func(r *Runtime) *Object {
		return r.global.ObjectPrototype
	}

	t.putStr("request", func(r *Runtime) Value { return r.methodProp(r.http_request, "request", 1) })
	t.putStr("get", func(r *Runtime) Value { return r.methodProp(r.http_get, "get", 2) })
	t.putStr("delete", func(r *Runtime) Value { return r.methodProp(r.http_delete, "delete", 2) })
	t.putStr("head", func(r *Runtime) Value { return r.methodProp(r.http_head, "head", 2) })
	t.putStr("options", func(r *Runtime) Value { return r.methodProp(r.http_options, "options", 2) })
	t.putStr("post", func(r *Runtime) Value { return r.methodProp(r.http_post, "post", 3) })
	t.putStr("put", func(r *Runtime) Value { return r.methodProp(r.http_put, "put", 3) })
	t.putStr("patch", func(r *Runtime) Value { return r.methodProp(r.http_patch, "patch", 3) })

	t.putSym(SymToStringTag, func(r *Runtime) Value { return valueProp(asciiString(classHttp), false, false, true) })

	return t
}

var httpTemplate *objectTemplate
var httpTemplateOnce sync.Once

func getHttpTemplate() *objectTemplate {
	httpTemplateOnce.Do(func() {
		httpTemplate = createHttpTemplate()
	})
	return httpTemplate
}

func (r *Runtime) getHttp() *Object {
	ret := r.global.Http
	if ret == nil {
		ret = &Object{runtime: r}
		r.global.Http = ret
		r.newTemplatedObject(getHttpTemplate(), ret)
	}
	return ret
}
