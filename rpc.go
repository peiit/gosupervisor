package gosupervisor

import "net/http"

type SupervisorRpc struct {
	Url string
	Transport http.RoundTripper
}

func New(url string, transport http.RoundTripper) *SupervisorRpc {
	return &SupervisorRpc{
		Url: url,
		Transport : transport,
	}
}

