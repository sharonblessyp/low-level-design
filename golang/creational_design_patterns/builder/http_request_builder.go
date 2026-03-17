package builder

import "fmt"

type httpRequestBuilder struct {
	method      string
	url         string
	headers     map[string]string
	queryParams map[string]string
	body        []byte
	timeout     int // milliseconds
}

func GetNewHttpRequestBuilder(method, url string) *httpRequestBuilder {
	return &httpRequestBuilder{
		method:      method,
		url:         url,
		headers:     make(map[string]string),
		queryParams: make(map[string]string),
	}
}

func (builder *httpRequestBuilder) AddHeaders(key, val string) *httpRequestBuilder {
	builder.headers[key] = val
	return builder
}

func (builder *httpRequestBuilder) AddQueryParams(key, val string) *httpRequestBuilder {
	builder.queryParams[key] = val
	return builder
}

func (builder *httpRequestBuilder) AddBody(body []byte) *httpRequestBuilder {
	builder.body = body
	return builder
}

func (builder *httpRequestBuilder) SetTimeout(timeout int) *httpRequestBuilder {
	builder.timeout = timeout
	return builder
}

func (builder *httpRequestBuilder) Build() (*httpRequest, error) {
	if builder.method == "" || builder.url == "" {
		return nil, fmt.Errorf("empty fields, method: %s, url: %s", builder.method, builder.url)
	}

	return &httpRequest{
		method:      builder.method,
		url:         builder.url,
		headers:     builder.headers,
		queryParams: builder.queryParams,
		body:        builder.body,
		timeout:     builder.timeout,
	}, nil
}
