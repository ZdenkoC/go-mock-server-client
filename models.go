package mockServerClient

import (
	"github.com/clarketm/json"
)

type BodyType string

const (
	BINARY     BodyType = "BINARY"
	JSON       BodyType = "JSON"
	JSONSchema BodyType = "JSON_SCHEMA"
	JSONPath   BodyType = "JSON_PATH"
	Parameters BodyType = "PARAMETERS"
	Regex      BodyType = "REGEX"
	String     BodyType = "STRING"
	XML        BodyType = "XML"
	XMLSchema  BodyType = "XML_SCHEMA"
	XPath      BodyType = "XPATH"
	LogEvent   BodyType = "LOG_EVENT"
)

type Request struct {
	Secure                bool             `json:"secure,omitempty"`
	KeepAlive             bool             `json:"keepAlive,omitempty"`
	Method                string           `json:"method,omitempty"`
	Path                  string           `json:"path,omitempty"`
	PathParameters        []PathParameter  `json:"pathParameters,omitempty"`
	QueryStringParameters []QueryParameter `json:"queryStringParameters,omitempty"`
	Body                  Body             `json:"body,omitempty"`
	Headers               []HTTPHeader     `json:"headers,omitempty"`
	Cookies               []HTTPCookie     `json:"cookies,omitempty"`
	SocketAddress         SocketAddress    `json:"socketAddress,omitempty"`
	Protocol              string           `json:"protocol,omitempty"`
}

func (r *Request) WithPath(path string) *Request {
	r.Path = path
	return r
}

type PathParameter struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

type QueryParameter struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

func WithQueryParameter(name string, values ...string) []QueryParameter {
	return []QueryParameter{
		{Name: name, Values: values},
	}
}

type Body struct {
	Base64Bytes string   `json:"base64Bytes,omitempty"`
	Not         bool     `json:"not,omitempty"`
	Type        BodyType `json:"type,omitempty"`
	ContentType string   `json:"contentType,omitempty"`
}

type HTTPHeader struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

type HTTPCookie struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type SocketAddress struct {
	Host   string `json:"host,omitempty"`
	Port   int    `json:"port,omitempty"`
	Scheme string `json:"scheme,omitempty"`
}

type Response struct {
	Delay             TimeUnit          `json:"delay,omitempty"`
	Body              Body              `json:"body,omitempty"`
	Cookies           []HTTPCookie      `json:"cookies,omitempty"`
	ConnectionOptions ConnectionOptions `json:"connectionOptions,omitempty"`
	Headers           []HTTPHeader      `json:"headers,omitempty"`
	StatusCode        int               `json:"statusCode,omitempty"`
	ReasonPhrase      string            `json:"reasonPhrase,omitempty"`
}

func JsonResponse(val any) Response {
	marshal, _ := json.Marshal(val)
	return Response{
		Body: JsonBody(string(marshal)),
	}
}

func StringResponse(val string) Response {
	return Response{
		Body: StringBody(val),
	}
}

type TimeUnit struct {
	TimeUnit string `json:"timeUnit,omitempty"`
	Value    int    `json:"value,omitempty"`
}

func JsonBody(value string) Body {
	return Body{
		Type:        JSON,
		ContentType: "application/json",
		Base64Bytes: value,
	}
}

func StringBody(value string) Body {
	return Body{
		Type:        String,
		ContentType: "text/plain",
		Base64Bytes: value,
	}
}

type ConnectionOptions struct {
	SuppressContentLengthHeader bool     `json:"suppressContentLengthHeader,omitempty"`
	ContentLengthHeaderOverride int      `json:"contentLengthHeaderOverride,omitempty"`
	SuppressConnectionHeader    bool     `json:"suppressConnectionHeader,omitempty"`
	ChunkSize                   int      `json:"chunkSize,omitempty"`
	KeepAliveOverride           bool     `json:"keepAliveOverride,omitempty"`
	CloseSocket                 bool     `json:"closeSocket,omitempty"`
	CloseSocketDelay            TimeUnit `json:"closeSocketDelay,omitempty"`
}

type HTTPOverrideForwardedRequest struct {
	Delay            TimeUnit         `json:"delay,omitempty"`
	RequestOverride  Request          `json:"requestOverride,omitempty"`
	RequestModifier  RequestModifier  `json:"requestModifier,omitempty"`
	ResponseOverride Response         `json:"responseOverride,omitempty"`
	ResponseModifier ResponseModifier `json:"responseModifier,omitempty"`
}

type RequestModifier struct {
	Path                  PathModifier          `json:"path,omitempty"`
	QueryStringParameters QueryStringParameters `json:"queryStringParameters,omitempty"`
	Headers               HTTPHeaders           `json:"headers,omitempty"`
	Cookies               HTTPCookies           `json:"cookies,omitempty"`
}

type PathModifier struct {
	Regex        string `json:"regex,omitempty"`
	Substitution string `json:"substitution,omitempty"`
}

type QueryStringParameters struct {
	Add     []QueryParameter `json:"add,omitempty"`
	Replace []QueryParameter `json:"replace,omitempty"`
	Remove  []string         `json:"remove,omitempty"`
}

type HTTPHeaders struct {
	Add     []HTTPHeader `json:"add,omitempty"`
	Replace []HTTPHeader `json:"replace,omitempty"`
	Remove  []string     `json:"remove,omitempty"`
}

type HTTPCookies struct {
	Add     []HTTPCookie `json:"add,omitempty"`
	Replace []HTTPCookie `json:"replace,omitempty"`
	Remove  []string     `json:"remove,omitempty"`
}

type ResponseModifier struct {
	Headers HTTPHeaders `json:"headers,omitempty"`
	Cookies HTTPCookies `json:"cookies,omitempty"`
}

type HTTPResponseTemplate struct {
	Delay        TimeUnit `json:"delay,omitempty"`
	TemplateType string   `json:"templateType,omitempty"`
	Template     string   `json:"template,omitempty"`
}

type HTTPResponseClassCallback struct {
	Delay         TimeUnit `json:"delay,omitempty"`
	CallbackClass string   `json:"callbackClass,omitempty"`
}

type HTTPResponseObjectCallback struct {
	Delay            TimeUnit `json:"delay,omitempty"`
	ClientID         string   `json:"clientId,omitempty"`
	ResponseCallback bool     `json:"responseCallback,omitempty"`
}

type HTTPForward struct {
	Delay  TimeUnit `json:"delay,omitempty"`
	Host   string   `json:"host,omitempty"`
	Port   int      `json:"port,omitempty"`
	Scheme string   `json:"scheme,omitempty"`
}

type HTTPForwardTemplate struct {
	Delay        TimeUnit `json:"delay,omitempty"`
	TemplateType string   `json:"templateType,omitempty"`
	Template     string   `json:"template,omitempty"`
}

type HTTPForwardClassCallback struct {
	Delay         TimeUnit `json:"delay,omitempty"`
	CallbackClass string   `json:"callbackClass,omitempty"`
}

type HTTPForwardObjectCallback struct {
	Delay            TimeUnit `json:"delay,omitempty"`
	ClientID         string   `json:"clientId,omitempty"`
	ResponseCallback bool     `json:"responseCallback,omitempty"`
}

type HTTPError struct {
	Delay          TimeUnit `json:"delay,omitempty"`
	DropConnection bool     `json:"dropConnection,omitempty"`
	ResponseBytes  string   `json:"responseBytes,omitempty"`
}

type Times struct {
	RemainingTimes int  `json:"remainingTimes,omitempty"`
	Unlimited      bool `json:"unlimited,omitempty"`
}

func Once() Times {
	return Times{
		RemainingTimes: 1,
		Unlimited:      false,
	}
}

type TimeToLive struct {
	TimeUnit   string `json:"timeUnit,omitempty"`
	TimeToLive int    `json:"timeToLive,omitempty"`
	Unlimited  bool   `json:"unlimited,omitempty"`
}

type Expectation struct {
	ID                           string                       `json:"id,omitempty"`
	Priority                     int                          `json:"priority,omitempty"`
	Request                      Request                      `json:"httpRequest,omitempty"`
	Response                     Response                     `json:"httpResponse,omitempty"`
	HTTPResponseTemplate         HTTPResponseTemplate         `json:"httpResponseTemplate,omitempty"`
	HTTPResponseClassCallback    HTTPResponseClassCallback    `json:"httpResponseClassCallback,omitempty"`
	HTTPResponseObjectCallback   HTTPResponseObjectCallback   `json:"httpResponseObjectCallback,omitempty"`
	HTTPForward                  HTTPForward                  `json:"httpForward,omitempty"`
	HTTPForwardTemplate          HTTPForwardTemplate          `json:"httpForwardTemplate,omitempty"`
	HTTPForwardClassCallback     HTTPForwardClassCallback     `json:"httpForwardClassCallback,omitempty"`
	HTTPForwardObjectCallback    HTTPForwardObjectCallback    `json:"httpForwardObjectCallback,omitempty"`
	HTTPOverrideForwardedRequest HTTPOverrideForwardedRequest `json:"httpOverrideForwardedRequest,omitempty"`
	HTTPError                    HTTPError                    `json:"httpError,omitempty"`
	Times                        Times                        `json:"times,omitempty"`
	TimeToLive                   TimeToLive                   `json:"timeToLive,omitempty"`
}

type ExpectationId struct {
	ID string `json:"id,omitempty"`
}

type VerifySequence struct {
	ExpectationIds                                      []ExpectationId `json:"expectationIds,omitempty"`
	Requests                                            []Request       `json:"httpRequests,omitempty"`
	MaximumNumberOfRequestToReturnInVerificationFailure int             `json:"maximumNumberOfRequestToReturnInVerificationFailure,omitempty"`
}

func Sequence() *VerifySequence {
	return &VerifySequence{}
}

func (vs *VerifySequence) WithIds(ids []string) *VerifySequence {
	var expectationIds []ExpectationId
	for _, id := range ids {
		expectationIds = append(expectationIds, ExpectationId{ID: id})
	}
	vs.ExpectationIds = expectationIds
	return vs
}

type VerifyTimes struct {
	AtLeast int `json:"atLeast"`
	AtMost  int `json:"atMost"`
}

func NeverCalled() VerifyTimes {
	return VerifyTimes{
		AtLeast: 0,
		AtMost:  0,
	}
}

type Verify struct {
	ExpectationID                                       ExpectationId `json:"expectationId,omitempty"`
	Request                                             Request       `json:"httpRequest,omitempty"`
	Times                                               VerifyTimes   `json:"times"`
	MaximumNumberOfRequestToReturnInVerificationFailure int           `json:"maximumNumberOfRequestToReturnInVerificationFailure,omitempty"`
}
