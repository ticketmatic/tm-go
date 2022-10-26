package ticketmatic

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"sort"
)

// Widget helper
type Widgets struct {
	AccountCode string
	AccessKey   string
	SecretKey   string
}

func NewWidgets(accountcode, accesskey, secretkey string) *Widgets {
	widgets := &Widgets{
		AccountCode: accountcode,
		AccessKey:   accesskey,
		SecretKey:   secretkey,
	}

	return widgets
}

type param struct {
	key   string
	value string
}

type paramSlice []param

func (a paramSlice) Len() int           { return len(a) }
func (a paramSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a paramSlice) Less(i, j int) bool { return a[i].key < a[j].key }

func (w *Widgets) GenerateUrl(widget string, params map[string]string) string {
	v := url.Values{}
	v.Set("accesskey", w.AccessKey)

	for k, val := range params {
		v.Set(k, val)
	}

	sig := w.calculateSignature(params)

	v.Set("signature", sig)

	url := fmt.Sprintf("%s/widgets/%s/%s?%s", Server, w.AccountCode, widget, v.Encode())
	return url
}

func (w *Widgets) VerifyReturnUrl(params map[string]string) error {
	key, ok := params["accesskey"]
	if !ok || key != w.AccessKey {
		return errors.New("Failed to verify return URL: Bad access key")
	}
	delete(params, "accesskey")

	sig, ok := params["signature"]
	if !ok {
		return errors.New("Failed to verify return URL: Signature missing")
	}
	delete(params, "signature")

	expected := w.calculateSignature(params)
	if expected != sig {
		return errors.New("Failed to verify return URL: Signature mismatch")
	}

	return nil
}

// This variant of VerifyReturnUrl allows you to pass in req.Form directly
func (w *Widgets) VerifyFormReturnUrl(values url.Values) error {
	params := make(map[string]string)
	for k, v := range values {
		for i, s := range v {
			if i == 0 {
				params[k] = s
			} else {
				params[k] += s
			}
		}
	}

	return w.VerifyReturnUrl(params)
}

func (w *Widgets) calculateSignature(params map[string]string) string {
	delete(params, "l")
	delete(params, "ordercode")

	paramObjs := make(paramSlice, 0)
	for k, v := range params {
		paramObjs = append(paramObjs, param{k, v})
	}

	sort.Sort(paramObjs)

	hash := ""
	for _, p := range paramObjs {
		hash += p.key + p.value
	}

	msg := fmt.Sprintf("%s%s%s", w.AccessKey, w.AccountCode, hash)

	mac := hmac.New(sha256.New, []byte(w.SecretKey))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}
