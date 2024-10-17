package goja

func (r *Runtime) http_get(call FunctionCall) Value {
	//url := call.Argument(0).String()

	return _undefined
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

func (r *Runtime) createHttpTemplate() *objectTemplate {
	t := newObjectTemplate()
	t.protoFactory = func(r *Runtime) *Object {
		return r.global.ObjectPrototype
	}

	t.putStr("get", func(r *Runtime) Value { return r.methodProp(r.http_get, "get", 0) })
	t.putStr("delete", func(r *Runtime) Value { return r.methodProp(r.http_delete, "delete", 0) })
	t.putStr("head", func(r *Runtime) Value { return r.methodProp(r.http_head, "head", 0) })
	t.putStr("options", func(r *Runtime) Value { return r.methodProp(r.http_options, "options", 0) })
	t.putStr("post", func(r *Runtime) Value { return r.methodProp(r.http_post, "post", 0) })
	t.putStr("put", func(r *Runtime) Value { return r.methodProp(r.http_put, "put", 0) })
	t.putStr("patch", func(r *Runtime) Value { return r.methodProp(r.http_patch, "patch", 0) })

	return t
}
