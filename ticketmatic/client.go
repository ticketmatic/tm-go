package ticketmatic

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"
)

// API server to use.
var Server = "https://apps.ticketmatic.com/api"

func init() {
	s := os.Getenv("TM_TEST_SERVER")
	if s != "" {
		Server = s
	}
}

// API Client
type Client struct {
	AccountCode string
	AccessKey   string
	SecretKey   string
}

// API Request
type Request struct {
	client *Client
	method string
	url    string

	params map[string]interface{}
	query  map[string]interface{}
	body   interface{}
}

func NewClient(accountcode, accesskey, secretkey string) *Client {
	client := &Client{
		AccountCode: accountcode,
		AccessKey:   accesskey,
		SecretKey:   secretkey,
	}

	return client
}

func (c *Client) NewRequest(method, url string) *Request {
	return &Request{
		client: c,
		method: method,
		url:    url,

		query: make(map[string]interface{}),
	}
}

func (r *Request) AddParameter(key string, val interface{}) {
	// Try to omit empty parameters by not sending them when they're set to
	// their default values.
	v := reflect.ValueOf(val)
	if v.Interface() != reflect.Zero(v.Type()).Interface() {
		r.query[key] = val
	}
}

func (r *Request) UrlParameters(params map[string]interface{}) {
	r.params = params
}

func (r *Request) Body(body interface{}) {
	r.body = body
}

func (r *Request) Run(obj interface{}) error {
	var body io.Reader

	if r.body != nil {
		d, err := json.Marshal(r.body)
		if err != nil {
			return err
		}
		body = bytes.NewReader(d)
	}

	u := r.prepareUrl()

	req, err := http.NewRequest(r.method, u, body)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", r.authHeader())
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "ticketmatic/go")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed: %d", resp.StatusCode)
	}

	if obj != nil {
		err = json.NewDecoder(resp.Body).Decode(obj)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Request) authHeader() string {
	ts := time.Now().UTC().Format("2006-01-02T15:04:05")

	mac := hmac.New(sha256.New, []byte(r.client.SecretKey))
	mac.Write([]byte(fmt.Sprintf("%s%s%s", r.client.AccessKey, r.client.AccountCode, ts)))
	hash := fmt.Sprintf("%x", mac.Sum(nil))

	return fmt.Sprintf("TM-HMAC-SHA256 key=%s ts=%s sign=%s", r.client.AccessKey, ts, hash)
}

func (r *Request) prepareUrl() string {
	u := r.url

	for k, v := range r.params {
		u = strings.Replace(u, fmt.Sprintf("{%s}", k), fmt.Sprintf("%v", v), 1)
	}
	u = strings.Replace(u, "{accountname}", r.client.AccountCode, 1)

	result := fmt.Sprintf("%s%s", Server, u)
	if len(r.query) > 0 {
		query := url.Values{}
		for k, v := range r.query {
			query.Add(k, fmt.Sprintf("%v", v))
		}
		result = fmt.Sprintf("%s?%s", result, query.Encode())
	}
	return result
}
