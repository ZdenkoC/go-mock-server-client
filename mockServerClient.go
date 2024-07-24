package mockServerClient

import (
	"bytes"
	"github.com/clarketm/json"
	"io"
	"net/http"
	"net/url"
)

type MockServerClient struct {
	url    *url.URL
	client *http.Client
}

func New(uri string) *MockServerClient {
	parsedUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	return &MockServerClient{
		client: http.DefaultClient,
		url:    parsedUrl,
	}
}

func (msc *MockServerClient) RegisterExpectation(expectation Expectation) {
	msc.sendData("/mockserver/expectation", expectation, http.MethodPut)
}

func (msc *MockServerClient) Verify(verify Verify) {
	msc.sendData("/mockserver/verify", verify, http.MethodPut)
}

func (msc *MockServerClient) VerifySequence(sequence VerifySequence) {
	msc.sendData("/mockserver/verifySequence", sequence, http.MethodPut)
}

func (msc *MockServerClient) ClearAll() {
	msc.sendData("/mockserver/clear?type=all", nil, http.MethodPut)
}

func (msc *MockServerClient) ClearById(expectationId string) {
	msc.sendData("/mockserver/clear?type=all", map[string]any{"id": expectationId}, http.MethodPut)
}

func (msc *MockServerClient) Clear(request Request) {
	msc.sendData("/mockserver/clear?type=all", request, http.MethodPut)
}

func (msc *MockServerClient) sendData(url string, data interface{}, method string) {
	var jsonData []byte
	var err error
	var req *http.Request
	if data != nil {
		jsonData, err = json.Marshal(data)
		if err != nil {
			panic(err)
		}
		req, err = http.NewRequest(method, msc.url.String()+url, bytes.NewBuffer(jsonData))
	} else {
		req, err = http.NewRequest(method, msc.url.String()+url, nil)
	}
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := msc.client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode/100 != 2 {
		d, _ := io.ReadAll(resp.Body)
		panic("bad response status: " + resp.Status + " with message: " + string(d))
	}
}
