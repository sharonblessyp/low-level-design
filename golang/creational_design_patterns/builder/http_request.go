package builder

type httpRequest struct {
	method      string
	url         string
	headers     map[string]string
	queryParams map[string]string
	body        []byte
	timeout     int
}

func (r *httpRequest) GetMethod() string {
	return r.method
}

func (r *httpRequest) GetUrl() string {
	return r.url
}

func (r *httpRequest) GetTimeout() int {
	return r.timeout
}

func (r *httpRequest) GetHeaders() map[string]string {
	return r.headers
}

func (r *httpRequest) GetQueryParams() map[string]string {
	return r.queryParams
}
