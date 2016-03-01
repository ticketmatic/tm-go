package ticketmatic

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"
)

// API server to use.
var Server = "https://apps.ticketmatic.com"

// API version
var Version = "1"

// Library Version
const Build = "dc4cfcffeb976820efc293608e320568e014ca92"

// Rate limit error
type RateLimitError struct {
	Status *QueueStatus
}

func (r *RateLimitError) Error() string {
	return "Rate Limit Exceeded"
}

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

	u, err := r.prepareUrl()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(r.method, u, body)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", r.authHeader())
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", fmt.Sprintf("ticketmatic/go (%s)", Build))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		status := &QueueStatus{}
		err = json.NewDecoder(resp.Body).Decode(status)
		if err != nil {
			return err
		}
		return &RateLimitError{
			Status: status,
		}
	} else if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Failed (%d): %s", resp.StatusCode, string(body))
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
	hash := Sign(r.client.AccessKey, r.client.SecretKey, r.client.AccountCode, ts)

	return fmt.Sprintf("TM-HMAC-SHA256 key=%s ts=%s sign=%s", r.client.AccessKey, ts, hash)
}

func (r *Request) prepareUrl() (string, error) {
	u := r.url

	for k, v := range r.params {
		u = strings.Replace(u, fmt.Sprintf("{%s}", k), fmt.Sprintf("%v", v), 1)
	}
	u = strings.Replace(u, "{accountname}", r.client.AccountCode, 1)

	result := fmt.Sprintf("%s/api/%s%s", Server, Version, u)
	if len(r.query) > 0 {
		query := url.Values{}
		for k, v := range r.query {
			kind := reflect.ValueOf(v).Kind()
			if kind == reflect.Interface || kind == reflect.Map || kind == reflect.Ptr || kind == reflect.Slice || kind == reflect.Struct {
				d, err := json.Marshal(v)
				if err != nil {
					return "", err
				}
				query.Add(k, fmt.Sprintf("%s", string(d)))
			} else {
				query.Add(k, fmt.Sprintf("%v", v))
			}
		}
		result = fmt.Sprintf("%s?%s", result, query.Encode())
	}
	return result, nil
}

// Generates a signed authentication hash
func Sign(accesskey, secretkey, accountcode, ts string) string {
	mac := hmac.New(sha256.New, []byte(secretkey))
	mac.Write([]byte(fmt.Sprintf("%s%s%s", accesskey, accountcode, ts)))
	return fmt.Sprintf("%x", mac.Sum(nil))
}
