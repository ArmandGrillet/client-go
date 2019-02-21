package dcos

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"
)

const (
	// Set a 10 seconds timeout for the connection to be established.
	DefaultHTTPClientDialContextTimeout = 10 * time.Second
	// Set it to 10 seconds as well for the TLS handshake when using HTTPS.
	DefaultHTTPClientTLSHandshakeTimeout = 10 * time.Second
	// The client will be dealing with a single host (the one in baseURL),
	// set max idle connections to 30 regardless of the host.
	DefaultHTTPClientMaxIdleConns        = 30
	DefaultHTTPClientMaxIdleConnsPerHost = 30
)

// DefaultTransport is a http.RoundTripper that adds authentication based on Config
type DefaultTransport struct {
	Config *Config
	Base   http.RoundTripper
}

func (t *DefaultTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

// RoundTrip authorizes requests to DC/OS by adding dcos_acs_token to Authorization header
func (t *DefaultTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// meet the requirements of RoundTripper and only modify a copy
	req2 := cloneRequest(req)
	req2.Header.Set("Authorization", fmt.Sprintf("token=%s", t.Config.ACSToken()))

	return t.base().RoundTrip(req2)
}

func cloneRequest(req *http.Request) *http.Request {
	req2 := new(http.Request)
	*req2 = *req

	// until now we only clone headers as we only modify those.
	req2.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		req2.Header[k] = append([]string(nil), s...)
	}

	return req2
}

// NewHTTPClient provides a http.Client able to communicate to dcos in an authenticated way
func NewHTTPClient(config *Config) *http.Client {
	client := &http.Client{}
	client.Transport = &http.Transport{

		// Allow http_proxy, https_proxy, and no_proxy.
		Proxy: http.ProxyFromEnvironment,

		DialContext: (&net.Dialer{
			Timeout: DefaultHTTPClientDialContextTimeout,
		}).DialContext,

		TLSHandshakeTimeout: DefaultHTTPClientTLSHandshakeTimeout,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.TLS().Insecure,
			RootCAs:            config.TLS().RootCAs,
		},

		MaxIdleConns:        DefaultHTTPClientMaxIdleConns,
		MaxIdleConnsPerHost: DefaultHTTPClientMaxIdleConnsPerHost,
	}

	return AddTransportHTTPClient(client, config)
}

// AddTransportHTTPClient adds dcos.DefaultTransport to http.Client to add dcos authentication
func AddTransportHTTPClient(client *http.Client, config *Config) *http.Client {
	transport := DefaultTransport{
		Config: config,
		Base:   client.Transport,
	}

	client.Transport = &transport

	return client
}