package zhaiker

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strings"
)

var debug = false

const (
	API_PREFIX          = "https://www.zhaiker.cn"
	API_GET_KEY         = "/open/OpenAPI!apiKey.m"
	API_GET_GYMS        = "/open/OpenAPI!getGyms.m"
	API_GET_EXAMS       = "/open/OpenAPI!getExams.m"
	API_GET_EXAMS_COUNT = "/open/OpenAPI!getExamsCount.m"
	API_GET_EXAM        = "/open/OpenAPI!getExamById.m"
	API_START_EXAM      = "/open/OpenAPI!startExam.m"
	API_CANCEL_EXAM     = "/open/OpenAPI!cancelExam.m"
	API_GET_DEVICE      = "/open/OpenAPI!getExamState.m"
	API_GET_FILE        = "/file/FileAction!loadFile.msg"
)

type Client struct {
	managerId string
	apiKey    string
}

// NewClient creates a client given manager id and api key.
func NewClient(managerId, apiKey string) *Client {
	return &Client{
		managerId: managerId,
		apiKey:    apiKey,
	}
}

// MustRequest makes an HTTP request to path with body params containing api
// key, unmarshals json response to pointer-key pair destinations, panics if
// operation fails.
func (client Client) MustRequest(ctx context.Context, path string, body url.Values, dest ...interface{}) {
	if err := client.Request(ctx, path, body, dest...); err != nil {
		panic(err)
	}
}

// Request makes an HTTP request to path with body params containing api key,
// unmarshals json response to pointer-key pair destinations.
func (client Client) Request(ctx context.Context, path string, body url.Values, dest ...interface{}) error {
	body.Set("apiKey", client.apiKey)
	if path == API_GET_GYMS {
		body.Set("managerId", client.managerId)
	}
	return Request(ctx, path, body, dest...)
}

// MustGetKey gets api key with account name and password, panics if operation
// fails.
func MustGetKey(ctx context.Context, accountName, accountPass string) (managerId, apiKey string) {
	managerId, apiKey, err := GetKey(ctx, accountName, accountPass)
	if err != nil {
		panic(err)
	}
	return managerId, apiKey
}

// GetKey gets api key with account name and password.
func GetKey(ctx context.Context, accountName, accountPass string) (managerId, apiKey string, err error) {
	v := url.Values{}
	v.Set("account", accountName)
	v.Set("password", fmt.Sprintf("%x", md5.Sum([]byte(accountPass))))
	err = Request(ctx, API_GET_KEY, v, &managerId, "managerId", &apiKey, "apiKey")
	return
}

// MustGetFileURL gets the URL of image path key, panics if operation fails.
func MustGetFileURL(ctx context.Context, path string) *url.URL {
	url, err := GetFileURL(ctx, path)
	if err != nil {
		panic(err)
	}
	return url
}

// GetFileURL gets the URL of image path key.
func GetFileURL(ctx context.Context, path string) (*url.URL, error) {
	body := Params("key", path)
	req, err := http.NewRequestWithContext(ctx, "POST", API_PREFIX+API_GET_FILE, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Location()
}

// Generate url.Values from list of keys and values.
func Params(keysAndValues ...interface{}) url.Values {
	params := url.Values{}
	for n := 0; n < len(keysAndValues)/2; n++ {
		params.Set(fmt.Sprint(keysAndValues[2*n]), fmt.Sprint(keysAndValues[2*n+1]))
	}
	return params
}

// MustRequest makes an HTTP request to path with body params, unmarshals json
// response to pointer-key pair destinations, panics if operation fails.
func MustRequest(ctx context.Context, path string, body url.Values, dest ...interface{}) {
	if err := Request(ctx, path, body, dest...); err != nil {
		panic(err)
	}
}

// Request makes an HTTP request to path with body params, unmarshals json
// response to pointer-key pair destinations.
func Request(ctx context.Context, path string, body url.Values, dest ...interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "POST", API_PREFIX+path, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}
	if debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return err
		}
		log.Println(string(dump))
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if debug {
		dumpBody := strings.Contains(resp.Header.Get("Content-Type"), "json")
		dump, err := httputil.DumpResponse(resp, dumpBody)
		if err != nil {
			return err
		}
		log.Println(string(dump))
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var respErr ResponseError
	json.Unmarshal(b, &respErr)
	if respErr.OK == false {
		return respErr
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("returned status %d instead of 200", resp.StatusCode)
	}
	if len(dest) == 0 {
		return nil
	}
	if len(dest) > 1 {
		for n := 0; n < len(dest)/2; n++ {
			arrange(b, dest[2*n], dest[2*n+1].(string))
		}
		return nil
	}
	if x, ok := dest[0].(*[]byte); ok {
		*x = b
		return nil
	}
	return json.Unmarshal(b, dest[0])
}

type ResponseError struct {
	Code    int    `json:"CODE"`
	Message string `json:"INFO"`
	OK      bool   `json:"STATUS"`
}

func (e ResponseError) Error() string {
	return fmt.Sprintf("zhaiker api error #%d: %s", e.Code, e.Message)
}

func arrange(data []byte, target interface{}, key string) {
	keys := strings.Split(key, ".")
	baseType := reflect.TypeOf(target).Elem()
	if baseType.Kind() == reflect.Slice {
		baseType = baseType.Elem()
	}
	typ := baseType
	for i := len(keys) - 1; i > -1; i-- {
		key := keys[i]
		if key == "*" {
			typ = reflect.SliceOf(typ)
		} else if key != "" {
			typ = reflect.MapOf(reflect.TypeOf(key), typ)
		}
	}
	d := reflect.New(typ)
	json.Unmarshal(data, d.Interface())
	items := collect(d.Elem(), keys)
	v := reflect.Indirect(reflect.ValueOf(target))
	if !v.IsValid() {
		return
	}
	for n := range items {
		item := items[n]
		if !item.IsValid() {
			item = reflect.New(baseType).Elem()
		}
		if v.Kind() == reflect.Slice {
			v.Set(reflect.Append(v, item))
		} else {
			v.Set(item)
		}
	}
}

func collect(x reflect.Value, keys []string) (out []reflect.Value) {
	for i, key := range keys {
		if key == "*" {
			k := keys[i+1:]
			for i := 0; i < x.Len(); i++ {
				out = append(out, collect(x.Index(i), k)...)
			}
			return
		} else if key != "" {
			x = x.MapIndex(reflect.ValueOf(key))
		}
	}
	out = append(out, x)
	return
}
