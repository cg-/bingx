package toolbox

import (
	"net/http"
)

const RECV_WINDOW int64 = 0

type HTTPClient struct {
	url       string
	apikey    string
	secretkey string
	client    *http.Client
}

func New(url, apikey, secretkey string) *HTTPClient {
	c := &HTTPClient{
		url:       url,
		apikey:    apikey,
		secretkey: secretkey,
		client:    http.DefaultClient,
	}
	return c
}

// Handle API requests. Takes in the method and endpoint as strings and a map of querys to send
func (c *HTTPClient) Request(method, endpoint string, payload map[string]string) (*http.Response, error) {
	var res *http.Response
	req, err := http.NewRequest(method, c.url+endpoint, nil)
	req.Header.Set("X-BX-APIKEY", c.apikey)

	if err != nil {
		return nil, err
	}

	query := GenerateQuery(c.secretkey, RECV_WINDOW, payload)

	switch method {
	case "GET":
		res, err = c.get(req, query)
	case "POST":
		res, err = c.post(req, query)
	case "PUT":
		res, err = c.put(req, query)
	case "DELETE":
		res, err = c.delete(req, query)
	default:
		return nil, &NotImplementedError{}
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Not sure we need to break these out like this but we'll do it just in case. Can refactor this out later
func (c *HTTPClient) get(req *http.Request, query string) (*http.Response, error) {
	return c.client.Do(req)
}
func (c *HTTPClient) post(req *http.Request, query string) (*http.Response, error) {
	return c.client.Do(req)
}
func (c *HTTPClient) put(req *http.Request, query string) (*http.Response, error) {
	return c.client.Do(req)
}
func (c *HTTPClient) delete(req *http.Request, query string) (*http.Response, error) {
	return c.client.Do(req)
}
