package ticketmatic

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// API server to use.
var Server = "https://apps.ticketmatic.com"

// API version
var Version = "1"

// Library Version
const Build = "1.0.116"

// Rate limit error
type RateLimitError struct {
	Backoff int
}

func (r *RateLimitError) Error() string {
	return fmt.Sprintf("Rate Limit Exceeded, wait for %d seconds", r.Backoff)
}

// Request error
type RequestError struct {
	StatusCode      int         `json:"code,omitempty"`
	Body            []byte      `json:"-"`
	Message         string      `json:"message,omitempty"`
	ApplicationCode int         `json:"applicationcode,omitempty"`
	ApplicationData interface{} `json:"applicationdata,omitempty"`
}

func (r *RequestError) Error() string {
	if r.Message != "" {
		return fmt.Sprintf("Failed (%d): %s", r.StatusCode, r.Message)
	} else {
		return fmt.Sprintf("Failed (%d): %s", r.StatusCode, string(r.Body))
	}
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
	Language    string
}

// API Request
type Request struct {
	client            *Client
	method            string
	url               string
	resultContentType string

	params          map[string]interface{}
	query           map[string]interface{}
	body            interface{}
	bodyContentType string
}

func NewClient(accountcode, accesskey, secretkey string) *Client {
	client := &Client{
		AccountCode: accountcode,
		AccessKey:   accesskey,
		SecretKey:   secretkey,
	}

	return client
}

func (c *Client) NewRequest(method, url, resultContentType string) *Request {
	if resultContentType == "" {
		resultContentType = "json"
	}
	return &Request{
		client:            c,
		method:            method,
		url:               url,
		resultContentType: resultContentType,

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

func (r *Request) Body(body interface{}, bodyContentType string) {
	r.body = body
	r.bodyContentType = bodyContentType
}

func (r *Request) Run(obj interface{}) error {
	resp, err := r.prepareRequest()
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if obj != nil {
		if r.resultContentType == "json" {
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			err = json.Unmarshal(data, obj)
			if err != nil {
				return fmt.Errorf("Deserialization failed: %s in %s", err, string(data))
			}
		} else {
			buff, ok := obj.(*bytes.Buffer)
			if !ok {
				return errors.New("Given obj is not *bytes.Buffer")
			}
			_, err := buff.ReadFrom(resp.Body)
			return err
		}
	}
	return nil
}

func (r *Request) Stream() (*Stream, error) {
	resp, err := r.prepareRequest()
	if err != nil {
		return nil, err
	}

	return NewStream(resp), nil
}

func (r *Request) prepareRequest() (*http.Response, error) {
	var body io.Reader

	if r.body != nil {
		if r.bodyContentType == "json" {
			d, err := json.Marshal(r.body)
			if err != nil {
				return nil, err
			}
			body = bytes.NewReader(d)
		} else if r.bodyContentType == "svg" {
			sBody, ok := r.body.(string)
			if !ok {
				return nil, errors.New("Supplied body is not a string, which is needed for body content type svg")
			}
			body = strings.NewReader(sBody)
		}
	}

	u, err := r.prepareUrl()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(r.method, u, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", r.authHeader())
	if r.bodyContentType == "json" {
		req.Header.Add("Content-Type", "application/json")
	} else if r.bodyContentType == "svg" {
		req.Header.Add("Content-Type", "image/svg+xml")
	}
	req.Header.Add("User-Agent", fmt.Sprintf("ticketmatic/go (%s)", Build))
	if r.client.Language != "" {
		req.Header.Add("Accept-Language", r.client.Language)
	}
	req.Close = true

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
		return resp, nil
	case 429:
		resp.Body.Close()

		backoff := int(0)
		hdr := resp.Header.Get("Retry-After")
		if len(hdr) > 0 {
			b, err := strconv.ParseInt(hdr, 10, 64)
			if err != nil {
				return nil, err
			}
			backoff = int(b)
		}
		return nil, &RateLimitError{
			Backoff: backoff,
		}
	default:
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		// Try to unmarshal the error, pass it back
		r := &RequestError{}
		err := json.Unmarshal(body, r)
		if err == nil && r.StatusCode > 0 && r.Message != "" {
			return nil, r
		}

		return nil, &RequestError{
			StatusCode: resp.StatusCode,
			Body:       body,
		}
	}
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
