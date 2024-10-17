package goja

import (
	"io"
	"net/http"
	"strings"
	"sync"
)

type httpResponse struct {
	Status     int
	StatusText string
	Headers    Value
	Data       string
}

func (r *Runtime) http_get(call FunctionCall) Value {
	url := call.Argument(0).String()

	response, err := http.Get(url)
	if err != nil {
		// log out err
		return _undefined
	}

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
		return _undefined
	}

	resp.Data = string(buf)

	return r.ToValue(resp)
}

func (r *Runtime) http_delete(call FunctionCall) Value {

	return _undefined
}

func (r *Runtime) http_head(call FunctionCall) Value {

	return _undefined
}

func (r *Runtime) http_options(call FunctionCall) Value {

	return _undefined
}

func (r *Runtime) http_post(call FunctionCall) Value {

	return _undefined
}

func (r *Runtime) http_put(call FunctionCall) Value {

	return _undefined
}

func (r *Runtime) http_patch(call FunctionCall) Value {

	return _undefined
}

func createHttpTemplate() *objectTemplate {
	t := newObjectTemplate()
	t.protoFactory = func(r *Runtime) *Object {
		return r.global.ObjectPrototype
	}

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
