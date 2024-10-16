package goja

func (r *Runtime) axios_get(call FunctionCall) Value {
	//url := call.Argument(0).String()

	return _undefined
}

func (r *Runtime) axios_delete(call FunctionCall) Value {

	return _undefined
}

func (r *Runtime) axios_head(call FunctionCall) Value {

	return _undefined
}

func (r *Runtime) axios_options(call FunctionCall) Value {

	return _undefined
}

func (r *Runtime) axios_post(call FunctionCall) Value {

	return _undefined
}

func (r *Runtime) axios_put(call FunctionCall) Value {

	return _undefined
}

func (r *Runtime) axios_patch(call FunctionCall) Value {

	return _undefined
}

func (r *Runtime) createAxios(val *Object) objectImpl {
	axios := &baseObject{
		class:      classAxios,
		val:        val,
		extensible: true,
		prototype:  r.global.ObjectPrototype,
	}
	axios.init()

	axios._putProp("get", r.newNativeFunc(r.axios_get, nil, "get", nil, 2), true, false, true)
	axios._putProp("delete", r.newNativeFunc(r.axios_delete, nil, "delete", nil, 2), true, false, true)
	axios._putProp("head", r.newNativeFunc(r.axios_head, nil, "head", nil, 2), true, false, true)
	axios._putProp("options", r.newNativeFunc(r.axios_options, nil, "options", nil, 2), true, false, true)
	axios._putProp("post", r.newNativeFunc(r.axios_post, nil, "post", nil, 3), true, false, true)
	axios._putProp("put", r.newNativeFunc(r.axios_put, nil, "put", nil, 3), true, false, true)
	axios._putProp("patch", r.newNativeFunc(r.axios_patch, nil, "patch", nil, 3), true, false, true)

	return axios
}

func (r *Runtime) initAxios() {
	r.global.Axios = r.newLazyObject(r.createAxios)

	r.addToGlobal(classAxios, r.global.Axios)
}
