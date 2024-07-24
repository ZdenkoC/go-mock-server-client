# MockServer Client for golang
Why I created this?  

I couldn't find one already that supported the latest version of MockServer. There are some implementations out there but do not have all the capabilities of the MockServer.

## Installation
```shell
go get github.com/zdenkoC/mockServerClient
```

## Usage
```go
import {
    m "github.com/zdenkoC/mockServerClient"
}

func main() {
    container = startMockServer() //starts mockServer as testcontainer. Not shown here
    mockServerUrl, _ := container.URL(context.Background())
    client := mockServerClient.New(mockServerUrl))

    client.RegisterExpectation(m.Expectation{
        Request:  m.Request{Path: "/", Headers: []m.HTTPHeader{{Name: "X-Session-Id", Values: []string{ctx.HttpSessionId}}}},
        Response: m.JsonResponse(map[string]any{"result": "alive"}),
        Times:    m.Once(),
        ID:       mockId(),
    })
    ...
    client.Verify(m.Verify{
        Request: m.Request{
			Path: "/", 
			Headers: []m.HTTPHeader{{Name: "X-Session-Id", Values: []string{ctx.HttpSessionId}}}, 
		},
        Times:   m.VerifyTimes{AtLeast: 1, AtMost: 1},
    })
    ...
    client.ClearById(id)
}
```

## License
Refer to the [Golang](https://github.com/golang/go/blob/master/LICENSE) license. See [LICENSE](LICENSE) for more information.
