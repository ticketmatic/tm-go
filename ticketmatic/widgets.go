package ticketmatic

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

	if l, ok := params["l"]; ok {
		v.Set("l", l)
		delete(params, "l")
	}

	paramObjs := make(paramSlice, 0)
	for k, v := range params {
		paramObjs = append(paramObjs, param{k, v})
	}

	sort.Sort(paramObjs)

	hash := ""
	for _, p := range paramObjs {
		hash += p.key + p.value
		v.Set(p.key, p.value)
	}

	msg := fmt.Sprintf("%s%s%s", w.AccessKey, w.AccountCode, hash)

	mac := hmac.New(sha256.New, []byte(w.SecretKey))
	mac.Write([]byte(msg))
	sig := hex.EncodeToString(mac.Sum(nil))

	v.Set("signature", sig)

	url := fmt.Sprintf("%s/widgets/%s/%s?%s", Server, w.AccountCode, widget, v.Encode())
	return url
}
