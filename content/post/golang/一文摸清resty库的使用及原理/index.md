---
title: "一文摸清resty库的使用及原理"
date: 2022-12-21T13:19:03+08:00
draft: false
image: "img/XXlYlcpWIio.jpg"
categories: 
  - golang
  - 源码学习
tag:
---

## 1. 序言

在项目中经常需要访问其他的服务，golang提供了`net/http`标准库可以帮助我们完成这样的需求，比如一个典型的GET请求的书写步骤是:

* 创建一个`http.Client`对象client
* 创建一个`http.Request`对象request，传入方法名、目标地址和请求体
* 以request对象为参数调用`client.Do`，获取响应
* 读取响应体，获取响应内容

代码如下:

```golang 
func main() {
	client := http.Client{}
	request, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Println(err)
		return
	}
	rsp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer rsp.Body.Close()
	content, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(content))
}
```

通过标准库可以很方便地发起一个请求，但是也存在一些使用效率上的问题，比如:

* 需要手动对content-type进行设置。不同的content-type对请求体有不同的设置，如json格式需要先序列化成字符串，x-www-form-urlencoded需要先进行encode，这就要求使用者明白理解content-type之间的区别。
* 不支持文件上传。本质上来说，文件上传也是通过conteng-type和请求体配合完成的，但是不明白原理，很难用好。
* 需要手动解析响应内容。正如上述案例代码，从body拿到的是一个字节流，需要使用者自己去决定怎么对字节流进行解析。
* ...

由于以上种种原因，导致net/http库虽然足够灵活，但是使用门槛较高。而resty库在`net/http`库的基础上，将使用上一些重要而又过于细节的部分封装起来，降低了使用难度，在项目中可以优先使用它来提高开发效率。

项目地址为: https://github.com/go-resty/resty

API文档地址: https://pkg.go.dev/github.com/go-resty/resty/v2#section-readme

> 申明: 本文会从该库的特点以及实现原理两个角度来进行阐述，不会涉及库的API使用

## 2. 特征

在深入学习该项目之前，首先看看README里对该项目功能上的描述，之后再从源码角度来分析是怎么实现的。

该库具备以下特征:

* 支持基本的HTTP方法，包括**GET、POST、PUT、DELETE、HEAD、PATCH、OPTIONS**等；
* 方法使用简单，支持链式调用；
* **支持多种类型的请求体，包括string、[]byte、struct、map、slice和io.Reader**；
  * 自动探测content-type；
  * 可以通过RawRequest访问内置的http.Request对象；
  * **可以通过RawRequest.GetBody()多次读取请求体；**
* 对于响应有更丰富的支持，比如:
  * 以字节流形式(`Body()`方法)，或以字符串形式(`String()`方法)访问响应内容
  * 通过`Time()`方法**获取响应处理时间**，通过`ReceivedAt()`方法获取收到响应的时间；
* 对于content-type为`json`和`xml`类型的请求和响应，可以**自动进行序列化和反序列化(可选，通过SetDoNotParseResponse选择是否启用)**；
  * 对于请求来说，如果将`struct/map`设置为请求的body，会默认将`content-type`设置为`json`;
  * 对于响应来说:
    * 如果成功，可以自动将响应内容反序列化，通过`SetResult()`和`Result()`设置或获取反序列化的目标对象；
    * 如果失败，可以自动将错误内容反序列化，通过`SetError()`和`Error()`设置或获取反序列化的目标对象；
    * 支持content-type为`application/problem+json` 和 `application/problem+xml`
    * 用户可以指定json/xml序列化和反序列化的具体实现方法；
  * 如果响应头中没有content-type，可以指定默认的Content-Type
* 支持content-type为`multipart/form-data`，可以方便地进行**多文件上传**；
* URL支持**路径参数**；
* 可以基于设定的重试函数进行**请求重试**；
* 支持自定义中间件；
* 请求头支持的Authorization类型包括`BasicAuth`和`Bearer`
* 支持自定义根证书和客户端证书；
* 支持将http响应结果保存到文件；
* 支持cookie；
* 支持一些客户端级别的配置，如`Timeout`、`RedirectPolicy`、`Proxy`、`TLSClientConfig`、`Transport`等；
* 允许为GET方法设置请求体；
* ...

## 3. 分析

在分析源代码之前，先来看看基于这个库如何发起一个post请求，下面是该库提供的一段测试代码:

```golang 
func Example_post() {
	// Create a resty client
	client := resty.New()

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"testuser", "password":"testpass"}`).
		SetResult(AuthSuccess{}). // or SetResult(&AuthSuccess{}).
		Post("https://myapp.com/login")

	printOutput(resp, err)
}
```

从上述代码中可以总结出基本的使用步骤为:

* 调用`resty.New()`创建`resty.Client`对象client
* 调用client的`SetXXX`方法进行一些客户端级别的配置
* 调用`client.NewRequest()`或`client.R()`创建`resty.Request`对象request 
* 调用request对象的`SetXXX`方法进行一些请求级别的配置，如`SetBody`、`SetHeader`等
* 调用request对象的`Post`方法发出请求，并接收响应

从上面可知，该库的基本请求功能是通过两个对象(resty.Client和resty.Request)完成的。

### 3.1 resty.Client对象

该对象定义位于`client.go`文件中，即: https://github.com/go-resty/resty/blob/master/client.go

#### 3.1.1 属性及其设置方法

```golang 
type Client struct {
	BaseURL               string
	HostURL               string // Deprecated: use BaseURL instead. To be removed in v3.0.0 release.
	QueryParam            url.Values
	FormData              url.Values
	PathParams            map[string]string
	Header                http.Header
	UserInfo              *User
	Token                 string
	AuthScheme            string
	Cookies               []*http.Cookie
	Error                 reflect.Type
	Debug                 bool
	DisableWarn           bool
	AllowGetMethodPayload bool
	RetryCount            int
	RetryWaitTime         time.Duration
	RetryMaxWaitTime      time.Duration
	RetryConditions       []RetryConditionFunc
	RetryHooks            []OnRetryFunc
	RetryAfter            RetryAfterFunc
	JSONMarshal           func(v interface{}) ([]byte, error)
	JSONUnmarshal         func(data []byte, v interface{}) error
	XMLMarshal            func(v interface{}) ([]byte, error)
	XMLUnmarshal          func(data []byte, v interface{}) error

	// HeaderAuthorizationKey is used to set/access Request Authorization header
	// value when `SetAuthToken` option is used.
	HeaderAuthorizationKey string

	jsonEscapeHTML     bool
	setContentLength   bool
	closeConnection    bool
	notParseResponse   bool
	trace              bool
	debugBodySizeLimit int64
	outputDirectory    string
	scheme             string
	log                Logger
	httpClient         *http.Client
	proxyURL           *url.URL
	beforeRequest      []RequestMiddleware
	udBeforeRequest    []RequestMiddleware
	preReqHook         PreRequestHook
	afterResponse      []ResponseMiddleware
	requestLog         RequestLogCallback
	responseLog        ResponseLogCallback
	errorHooks         []ErrorHook
}
```

从作用上可以将其属性分成以下几类:

* 请求所需的变量: 如HostUrl、QueryParam、FormData、PathParams、Header、Cookies等。这些变量其实是请求级别的，这里相当于是提供了一个默认的客户端级别配置。如果请求级别不对这些配置进行覆盖，则默认使用客户端级别的配置；
* 控制请求过程的变量: 如布尔类型的变量、RetryXXX(请求重试相关的)
* 拓展插件: 这些变量的类型为RequestMiddleware、XXXHook等，主要是定义一种抽象，方便用户对库的功能的进行拓展。

针对这些字段，该对象提供了用于设置的方法，如:

* SetHostURL用于设置请求的URL
* SetHeader用于设置请求头
* SetCookie用于设置Cookie
* SetQueryParam、SetFormData、SetAuthToken
* OnBeforeRequest用于设置请求的前置处理器
* ...

#### 3.1.2 对象构造方法

主要是Client对象和Request对象的构造方法。

##### 1. createClient方法

该方法用于创建Client对象，对Client对象的各个字段进行初始化。这里需要注意的是，该方法内为Client对象设置了几个前置钩子和后置钩子。

* 前置钩子用于在发起请求前对请求进行处理，包括`parseRequestURL、parseRequestHeader、parseRequestBody、createHTTPRequest和addCredentials`。这里先记着有这几个函数，后面再详细分析。
* 后置钩子用于对响应进行处理，包括`responseLogger、parseResponseBody、saveResponseIntoFile`。

```golang 
func createClient(hc *http.Client) *Client {
	if hc.Transport == nil {
		hc.Transport = createTransport(nil)
	}

	c := &Client{ // not setting lang default values
        ....
	}

	// Logger
	c.SetLogger(createLogger())

	// default before request middlewares
	// NOTE: 设置前置钩子
	c.beforeRequest = []RequestMiddleware{
		parseRequestURL,
		parseRequestHeader,
		parseRequestBody,
		createHTTPRequest,
		addCredentials,
	}

	// user defined request middlewares
	// NOTE: 设置用户自定义的前置钩子
	c.udBeforeRequest = []RequestMiddleware{}

	// default after response middlewares
	// NOTE: 设置后置钩子
	c.afterResponse = []ResponseMiddleware{
		responseLogger,
		parseResponseBody,
		saveResponseIntoFile,
	}
	return c
}
```

##### 2. R()方法

从这里就知道我们可以通过`R()`和`NewRequest()`方法来创建resty.Request对象。

```golang 
func (c *Client) R() *Request {
	r := &Request{
		QueryParam: url.Values{},
		FormData:   url.Values{},
		Header:     http.Header{},
		Cookies:    make([]*http.Cookie, 0),

		client:          c,
		multipartFiles:  []*File{},
		multipartFields: []*MultipartField{},
		PathParams:      map[string]string{},
		jsonEscapeHTML:  true,
	}
	return r
}

// NewRequest is an alias for method `R()`. Creates a new request instance, its used for
// Get, Post, Put, Delete, Patch, Head, Options, etc.
func (c *Client) NewRequest() *Request {
	return c.R()
}
```

#### 3.1.3 核心方法

`execute`方法是请求的核心方法，该方法接收`resty.Request`对象作为参数，返回`resty.Response`对象和error。

下面是执行过程:

1.执行用户自定义的请求前置处理器。遍历udBeforeRequest数组，将resty.Request对象作为参数，对该请求进行处理。

```golang
  for _, f := range c.udBeforeRequest {
    if err = f(c, req); err != nil {
      return nil, wrapNoRetryErr(err)
    }
  }
```

2.执行系统默认的请求前置处理器。遍历beforeRequest数组，将resty.Request对象作为参数，对该请求进行处理。这里默认的前置处理器就是3.1.2-1中提到的前置钩子。

```golang
for _, f := range c.beforeRequest {
  if err = f(c, req); err != nil {
    return nil, wrapNoRetryErr(err)
  }
}
```

3.设置底层`http.Request`对象的Host属性。该属性值来源于resty.Request的header。这里是否有疑问，RawRequest是哪里来的?

其实是在步骤2中的`createHTTPRequest`方法中创建的。

```golang
  if hostHeader := req.Header.Get("Host"); hostHeader != "" {
    req.RawRequest.Host = hostHeader
  }
```

4.执行前置钩子。这个钩子和2中的钩子不同之处在于，这里只能设置一个，处理顺序更靠后。

从这里也能看出，几种钩子的执行顺序为: udBeforeRequest > beforeRequest > preReqHook。

```golang 
if c.preReqHook != nil {
  if err = c.preReqHook(c, req.RawRequest); err != nil {
    return nil, wrapNoRetryErr(err)
  }
}
```

5.记录请求日志。

```golang 
// NOTE: 4.记录请求日志
if err = requestLogger(c, req); err != nil {
  return nil, wrapNoRetryErr(err)
}
```

6.重新设置底层http.requet的body，记录发请求的时间，发起请求。

```golang
// NOTE: 5.填充请求体，该请求体内置缓存结构
req.RawRequest.Body = newRequestBodyReleaser(req.RawRequest.Body, req.bodyBuf)
req.Time = time.Now()
// NOTE: 6.调用go标准库的Do方法发起请求
resp, err := c.httpClient.Do(req.RawRequest)
```

7.如果请求过程有错误(err != nil)或设置了不解析响应标记(notParseResponse)，则直接返回resty.Response对象，该对象内部封装了http.Response(RawResponse)。

```golang
// NOTE: 7.记录收到响应的时间
response := &Response{
  Request:     req,
  RawResponse: resp,
}
if err != nil || req.notParseResponse || c.notParseResponse {
  response.setReceivedAt()
  return response, err
}
```

下面是resty.Response对象的定义:

```golang
type Response struct {
	Request     *Request
	RawResponse *http.Response

	body       []byte
	size       int64
	receivedAt time.Time
}
```

8.如果设置了isSaveResponse，则执行该分支。进行以下处理:

* 如果设置了gzip压缩，先进行解压；
* 从body里面读取字节流到response.body字段中，该字段类型为字节数组；
* 设置响应体长度；

```golang 
if !req.isSaveResponse {
    defer closeq(resp.Body)
    body := resp.Body

    // GitHub #142 & #187
    // NOTE: 8.如果响应设置了压缩格式且响应体长度不为0，
    if strings.EqualFold(resp.Header.Get(hdrContentEncodingKey), "gzip") && resp.ContentLength != 0 {
        if _, ok := body.(*gzip.Reader); !ok {
            body, err = gzip.NewReader(body)
            if err != nil {
                response.setReceivedAt()
                return response, err
            }
            	defer closeq(body)
          }
    }

    // NOTE: 9.读取到response中
    if response.body, err = ioutil.ReadAll(body); err != nil {
        response.setReceivedAt()
        return response, err
    }

    // NOTE: 设置响应体长度
    response.size = int64(len(response.body))
}
```

9.记录收到响应的时间。

```golang
response.setReceivedAt() // after we read the body
```

10.后置钩子处理响应。即3.1.2-1中提到的后置钩子。

```golang
// NOTE: 10.执行默认的后置响应处理器
for _, f := range c.afterResponse {
    if err = f(c, response); err != nil {
    	  break
    }
}
```

以上就是一个`resty.Request`对象从处理到收到响应的核心流程，在以上流程中，暂时先略过了几个点，将在后续章节中一一捋清:

* 步骤6中使用`req.RawRequest.Body`和`req.bodyBuf`构造了一个`RequestBodyReleaser`对象，其中的bodybuf从哪里来的，有什么用处?
* 步骤2和步骤10中的处理钩子，具体进行了哪些处理?

### 3.2 resty.Request对象

该对象定义位于`request.go`文件中，即: https://github.com/go-resty/resty/blob/master/request.go

#### 3.2.1 属性及其设置方法

同Client对象类似，Request对象中也有很多请求有关的属性及设置方法。如URL、Method、Token、QueryParam等。

```golang 
type Request struct {
	URL        string
	Method     string
	Token      string
	AuthScheme string
	QueryParam url.Values
	FormData   url.Values
	PathParams map[string]string
	Header     http.Header
	Time       time.Time
	Body       interface{}
	Result     interface{}
	Error      interface{}
	RawRequest *http.Request
	SRV        *SRVRecord
	UserInfo   *User
	Cookies    []*http.Cookie

	// Attempt is to represent the request attempt made during a Resty
	// request execution flow, including retry count.
	//
	// Since v2.4.0
	Attempt int

	isMultiPart         bool
	isFormData          bool
	setContentLength    bool
	isSaveResponse      bool
	notParseResponse    bool
	jsonEscapeHTML      bool
	trace               bool
	outputFile          string
	fallbackContentType string
	forceContentType    string
	ctx                 context.Context
	values              map[string]interface{}
	client              *Client
	bodyBuf             *bytes.Buffer
	clientTrace         *clientTrace
	multipartFiles      []*File
	multipartFields     []*MultipartField
	retryConditions     []RetryConditionFunc
}
```

#### 3.2.2 请求方法

在章节3的开头，在一系列的链式调用最后，调用了`Post`方法，其实调用的就是这里的方法。当然除了Post方法外，http的所有基本方法都在这里进行了定义。如下:

```golang
// Get method does GET HTTP request. It's defined in section 4.3.1 of RFC7231.
func (r *Request) Get(url string) (*Response, error) {
	return r.Execute(MethodGet, url)
}

// Head method does HEAD HTTP request. It's defined in section 4.3.2 of RFC7231.
func (r *Request) Head(url string) (*Response, error) {
	return r.Execute(MethodHead, url)
}

// Post method does POST HTTP request. It's defined in section 4.3.3 of RFC7231.
func (r *Request) Post(url string) (*Response, error) {
	return r.Execute(MethodPost, url)
}

// Put method does PUT HTTP request. It's defined in section 4.3.4 of RFC7231.
func (r *Request) Put(url string) (*Response, error) {
	return r.Execute(MethodPut, url)
}

// Delete method does DELETE HTTP request. It's defined in section 4.3.5 of RFC7231.
func (r *Request) Delete(url string) (*Response, error) {
	return r.Execute(MethodDelete, url)
}

// Options method does OPTIONS HTTP request. It's defined in section 4.3.7 of RFC7231.
func (r *Request) Options(url string) (*Response, error) {
	return r.Execute(MethodOptions, url)
}

// Patch method does PATCH HTTP request. It's defined in section 2 of RFC5789.
func (r *Request) Patch(url string) (*Response, error) {
	return r.Execute(MethodPatch, url)
}

// Send method performs the HTTP request using the method and URL already defined
// for current `Request`.
//      req := client.R()
//      req.Method = resty.GET
//      req.URL = "http://httpbin.org/get"
// 		resp, err := client.R().Send()
func (r *Request) Send() (*Response, error) {
	return r.Execute(r.Method, r.URL)
}
```

这些函数都是调用的`Execute`方法。其处理流程是:

1.校验。`isMultiPart`表示是multipart请求，对应content-type为`multipart/formdata`，这种content-type下，请求方法只能为`Post/PUT/PATCH`。

```golang
  if r.isMultiPart && !(method == MethodPost || method == MethodPut || method == MethodPatch) {
      // No OnError hook here since this is a request validation error
      return nil, fmt.Errorf("multipart content is not allowed in HTTP verb [%v]", method)
  }
```

2.如果有设置SRV，则从srv记录里面获取该服务的目标主机地址，并选择一个地址作为请求的目标地址。这里调用了`selectAddr`方法，会将协议、服务的主机地址以及请求路径组合成一个完成的路径作为请求的URL。

```golang
  // NOTE: 查询服务的地址等信息
  if r.SRV != nil {
      _, addrs, err = net.LookupSRV(r.SRV.Service, "tcp", r.SRV.Domain)
      if err != nil {
          r.client.onErrorHooks(r, nil, err)
          return nil, err
      }
  }

  r.Method = method
  // NOTE: 获取完整的请求路径
  r.URL = r.selectAddr(addrs, url, 0)
```

3.如果`RetryCount`为0，即失败时不进行重试，就调用`client.execute`方法进行请求处理(前面3.1.3已经分析过)，获取到resp和err对象。

除此以外，还要将返回的err对象传递给onErrorHooks函数，作为错误回调。

```golang
	// NOTE: 重试次数设置为0，不执行退避操作，执行完直接返回
	if r.client.RetryCount == 0 {
		r.Attempt = 1
		resp, err = r.client.execute(r)
		r.client.onErrorHooks(r, resp, unwrapNoRetryErr(err))
		return resp, unwrapNoRetryErr(err)
	}
```

4.如果`RetryCount`不为0，则需要先执行重试逻辑，即下面的`Backoff`函数部分。然后再执行错误回调。

有关重试逻辑将在章节4中进行详细的分析。

```golang
err = Backoff(
    // NOTE: backoff第一个函数型参数传入的是请求的操作，会多次执行
    func() (*Response, error) {
          r.Attempt++

          r.URL = r.selectAddr(addrs, url, r.Attempt)

          resp, err = r.client.execute(r)
          if err != nil {
             	r.client.log.Errorf("%v, Attempt %v", err, r.Attempt)
          }
           return resp, err
	  },
      // NOTE: 其他的选项参数
      Retries(r.client.RetryCount),
      WaitTime(r.client.RetryWaitTime),
      MaxWaitTime(r.client.RetryMaxWaitTime),
      RetryConditions(append(r.retryConditions, r.client.RetryConditions...)),
      RetryHooks(r.client.RetryHooks),
)

// NOTE: 重试请求完后执行错误钩子
r.client.onErrorHooks(r, resp, unwrapNoRetryErr(err))

return resp, unwrapNoRetryErr(err)
```

## 4.其他

前面从两个对象的字段及其方法的角度分析了resty库发起请求的一般流程及其原理，但是还留了一些比较细节的部分，这些虽然比较细节，却是resty库中不可或缺的。下面通过问题的形式来看看这部分是怎样的。

### 4.1 resty.Request对象的bodybuf字段是在哪里初始化的？有什么用?



### 4.2 钩子做了些啥?

在`client.go`中定义了这样一些类型，如ResponseMiddleware、RequestMiddleware等。

从名字上可知它们其实是为请求或响应提供某种中间件能力的。如RequestMiddleware接收`resty.Client`和`resty.Request`对象作为参数，用于对请求进行处理，而ResponseMiddleware接收`resty.Client`和`resty.Request`对象作为参数，用于对响应进行处理。

```golang
type (
	// RequestMiddleware type is for request middleware, called before a request is sent
	RequestMiddleware func(*Client, *Request) error

	// ResponseMiddleware type is for response middleware, called after a response has been received
	ResponseMiddleware func(*Client, *Response) error

	// PreRequestHook type is for the request hook, called right before the request is sent
	PreRequestHook func(*Client, *http.Request) error

	// RequestLogCallback type is for request logs, called before the request is logged
	RequestLogCallback func(*RequestLog) error

	// ResponseLogCallback type is for response logs, called before the response is logged
	ResponseLogCallback func(*ResponseLog) error

	// ErrorHook type is for reacting to request errors, called after all retries were attempted
	ErrorHook func(*Request, error)
)
```

不难猜出，肯定有一个地方会统一保存这些类型的结构，然后在一个地方进行注册并在某个地方进行使用。

**在哪里保存?**

resty.Client中申明了一些RequestMiddleware和ResponseMiddleware类型(当然还有其他类型)的数组，作为客户端级别的配置。

```golang
type Client struct {
	RetryHooks            []OnRetryFunc
	beforeRequest      []RequestMiddleware
	udBeforeRequest    []RequestMiddleware
	preReqHook         PreRequestHook
	afterResponse      []ResponseMiddleware
	requestLog         RequestLogCallback
	responseLog        ResponseLogCallback
	errorHooks         []ErrorHook
}
```

**在哪里注册?**

在调用`resty.New()`方法创建Client对象时，实际调用`resty.createClient()`方法。该方法会注册一些默认的middleware。

通过代码可知，默认的前置钩子包括:

* parseRequestURL
* parseRequestHeader
* parseRequestBody
* createHTTPRequest
* addCredentials

默认的后置钩子包括:

* responseLogger
* parseResponseBody
* saveResponseIntoFile

```golang
c.beforeRequest = []RequestMiddleware{
    parseRequestURL,
    parseRequestHeader,
    parseRequestBody,
    createHTTPRequest,
    addCredentials,
}
c.afterResponse = []ResponseMiddleware{
    responseLogger,
    parseResponseBody,
    saveResponseIntoFile,
}
```

**在哪里使用?**

在章节3.1.3分析`execute`方法时讲过，这些前置和后置钩子，会按照某种优先级依次执行。实现细节将在4.5节中解析。

### 4.3 重试机制是如何实现的?

在3.2.2结尾讲过，如果`RetryCount`不为0，则执行`Backoff`函数。其位于`retry.go`文件中。

https://github.com/go-resty/resty/blob/master/retry.go

下面是其执行逻辑:

1.使用options模式进行初始化，有关这部分可以参考: https://www.zhihu.com/search?type=content&q=option%20golang。结合调用时传入的参数可知，这里先设置了重试次数、重试等待时间、最大重试等待时间、重试条件、重试钩子函数。

```golang
  opts := Options{
      maxRetries:      defaultMaxRetries,
      waitTime:        defaultWaitTime,
      maxWaitTime:     defaultMaxWaitTime,
      retryConditions: []RetryConditionFunc{},
  }

  for _, o := range options {
      o(&opts)
  }
```

2.进入控制重试的循环，用attempt代表已重试次数，当达到最大重试次数maxRetries时退出循环。

3.执行通过参数传进来的`operation`函数。结合调用传入的参数可知，执行的其实是下面这个函数，该函数执行发请求及相应处理的逻辑。

```golang
func() (*Response, error) {
      r.Attempt++

      r.URL = r.selectAddr(addrs, url, r.Attempt)

      resp, err = r.client.execute(r)
      if err != nil {
       		 r.client.log.Errorf("%v, Attempt %v", err, r.Attempt)
      }

      return resp, err
}
```

4.判断是否需要重试。重试要求为: 响应有err 或注册的重试条件有一个满足。如果无需重试，直接返回err；否则继续执行。

```GOLANG
needsRetry := err != nil && err == err1 // retry on a few operation errors by default

// NOTE: 满足一个重试条件即跳出
for _, condition := range opts.retryConditions {
    needsRetry = condition(resp, err1)
    if needsRetry {
    	  break
    }
}

// NOTE: 不重试直接返回错误
if !needsRetry {
  return err
}
```

5.执行重试钩子。

```golang 
for _, hook := range opts.retryHooks {
  	hook(resp, err)
}
```

6.判断是否已达到最大重试次数，如果已达到则返回响应的err。

```golang
if attempt == opts.maxRetries {
  	return err
}
```

7.计算等待时间。这里用到了退避算法，将在后面详细说明。

```golang
waitTime, err2 := sleepDuration(resp, opts.waitTime, opts.maxWaitTime, attempt)
if err2 != nil {
    if err == nil {
     	 err = err2
    }
    return err
}
```

8.利用select-case实现等待。select会阻塞当前协程，直到调用了context.Cancel方法或者计时器时间到了。

```golang
// NOTE: 等待一定时间后再重试
select {
  	case <-time.After(waitTime):
	  case <-ctx.Done():
  			return ctx.Err()
}
```

### 4.4 重试间隔是怎么计算的?

上面4.3中我们提到了会计算一个等待时间。这部分由`sleepDuration`函数来完成，且看看。

```golang
func sleepDuration(resp *Response, min, max time.Duration, attempt int) (time.Duration, error) {
	const maxInt = 1<<31 - 1 // max int for arch 386
	if max < 0 {
		max = maxInt
	}
	if resp == nil {
		return jitterBackoff(min, max, attempt), nil
	}

	retryAfterFunc := resp.Request.client.RetryAfter

	// Check for custom callback
	if retryAfterFunc == nil {
		return jitterBackoff(min, max, attempt), nil
	}

	result, err := retryAfterFunc(resp.Request.client, resp)
	if err != nil {
		return 0, err // i.e. 'API quota exceeded'
	}
	if result == 0 {
		return jitterBackoff(min, max, attempt), nil
	}
	if result < 0 || max < result {
		result = max
	}
	if result < min {
		result = min
	}
	return result, nil
}
```

最外层是sleepDuration函数，该函数主要是做一些检查工作。将waitTime赋值给参数min，将maxWaitTime赋值给max，表示左右边界。

1.判断最大间隔时间是否合法，不合法则将max设置为最大32位整数值；

2.判断是否定义了`RetryAfter`重试钩子，如果没定义则基于min、max和attempt调用`jitterBackoff`进行计算并作为等待时间；

3.如果定义了`RetryAfter`钩子，先调用该钩子计算等待时间，如果计算出来的等待时间为0，再调用`jitterBackoff`进行计算；

4.如果3中钩子计算出来的等待时间不为0，还需保证该时间在[min,max]之间。

`jitterBackoff`函数是算法`Exponential Backoff And Jitter`的具体实现。这里其实就是实现了这个公式，其中cap是最大值，base是最小值。

```golang
temp = min(cap, base * 2 ** attempt)
sleep = temp / 2 + rand_between(0, temp / 2)
```

可参考: https://aws.amazon.com/cn/blogs/architecture/exponential-backoff-and-jitter/

```golang
func jitterBackoff(min, max time.Duration, attempt int) time.Duration {
	base := float64(min)
	capLevel := float64(max)

	// temp = min(cap, base * 2 ** attempt)
	temp := math.Min(capLevel, base*math.Exp2(float64(attempt)))
	// sleep = temp / 2 + rand_between(0, temp / 2)
	ri := time.Duration(temp / 2)
	result := randDuration(ri)

	if result < min {
		result = min
	}

	return result
}

var rnd = newRnd()
var rndMu sync.Mutex

// NOTE: 在center基础上添加一个随机的漂移时间
func randDuration(center time.Duration) time.Duration {
	rndMu.Lock()
	defer rndMu.Unlock()

	var ri = int64(center)
	var jitter = rnd.Int63n(ri)
	return time.Duration(math.Abs(float64(ri + jitter)))
}
```

### 4.5 钩子详解

这里选取几个比较重要的钩子进行解析，包括parseRequestURL、parseRequestHeader、parseRequestBody、createHTTPRequest、addCredentials和parseResponseBody。

#### 4.5.1 parseRequestURL

这个钩子的作用是对请求的URL进行解析。

1.对路径参数进行解析。比如我们设置的URL为`getInfo/{id}`，则会将`{id}`替换为id变量实际的值。

```golang
//  NOTE: 替换掉请求路径中的{XXX}(请求级别的)
if len(r.PathParams) > 0 {
    for p, v := range r.PathParams {
     	 r.URL = strings.Replace(r.URL, "{"+p+"}", url.PathEscape(v), -1)
    }
}
//  NOTE: 替换掉请求路径中的{XXX}(客户端级别的)
if len(c.PathParams) > 0 {
    for p, v := range c.PathParams {
      	r.URL = strings.Replace(r.URL, "{"+p+"}", url.PathEscape(v), -1)
    }
}
```

2.对相对路径的URL进行处理。如果r.URL不是绝对路径，首先会在头部加一个/，然后将HostUrl和r.URL组合成完整的URL。

```golang
// NOTE: 解析URL
reqURL, err := url.Parse(r.URL)
if err != nil {
  	return err
}

// If Request.URL is relative path then added c.HostURL into
// the request URL otherwise Request.URL will be used as-is
// NOTE: 如果请求地址不是绝对地址，即不包含协议，需要进行处理
if !reqURL.IsAbs() {
    // NOTE: 在请求路径前加/
    r.URL = reqURL.String()
    if len(r.URL) > 0 && r.URL[0] != '/' {
      	r.URL = "/" + r.URL
    }

  // NOTE: 用client的host部分补充请求的路径，使之成为完整的路径
  reqURL, err = url.Parse(c.HostURL + r.URL)
  if err != nil {
    	return err
  }
}
```

3.添加协议。如果2中组合之后的URL没有协议部分(且Client设置了协议)，则将客户端协议拼接到URL之前。

```golang
if reqURL.Scheme == "" && len(c.scheme) > 0 {
  	reqURL.Scheme = c.scheme
}
```

4.拼接查询参数。这里涉及到两部分，先通过一个map收纳客户端级别的参数，再收纳请求级别的参数，如果遇到同名的，则优先用请求级别的。最后将各查询参数用&连接成字符串。

```golang
query := make(url.Values)
for k, v := range c.QueryParam {
    for _, iv := range v {
      	query.Add(k, iv)
    }
}
// NOTE: 如果请求中对查询参数进行了设置，则覆盖掉client中对该查询参数的设置
for k, v := range r.QueryParam {
    // remove query param from client level by key
    // since overrides happens for that key in the request
    query.Del(k)

    for _, iv := range v {
     	 query.Add(k, iv)
    }
}
// NOTE: 拼接查询字符串
if len(query) > 0 {
    if IsStringEmpty(reqURL.RawQuery) {
     	 reqURL.RawQuery = query.Encode()
    } else {
     	 reqURL.RawQuery = reqURL.RawQuery + "&" + query.Encode()
    }
}
```

5.url.URL结构转换成string，至此req.URL就是一个完整的请求路径了。包括协议、主机名、请求路径、查询参数。

```golang
r.URL = reqURL.String()
```

#### 4.5.2 parseRequestHeader

这个函数比较短，主要作用就是为请求添加header。

由于header在client和request级别都可以进行配置，所以也需要指定一个优先级，这里要求请求级别的header优先级大于client级别。

先把所有的header收纳到hdr这个结构中，然后为请求添加默认的header(如果不存在的话)，包括`Agent`和`Accept`。处理后的header回写回到resty.Request中。

这里可以看到，当前header还是停留在resty.Request中，并没有写入http.Request中。

```golang
func parseRequestHeader(c *Client, r *Request) error {
	// NOTE: 先将client中的header添加进来
	hdr := make(http.Header)
	for k := range c.Header {
		hdr[k] = append(hdr[k], c.Header[k]...)
	}

	// NOTE: 再将request设置的header添加进来，如果遇到和client设置同名的，覆盖掉
	for k := range r.Header {
		hdr.Del(k)
		hdr[k] = append(hdr[k], r.Header[k]...)
	}

	// NOTE: 添加默认的User-Agent头
	if IsStringEmpty(hdr.Get(hdrUserAgentKey)) {
		hdr.Set(hdrUserAgentKey, hdrUserAgentValue)
	}

	ct := hdr.Get(hdrContentTypeKey)
	// NOTE: 如果设置了Accept且ContentType为json或xml，则设置Accept为ContentType的值
	if IsStringEmpty(hdr.Get(hdrAcceptKey)) && !IsStringEmpty(ct) &&
		(IsJSONType(ct) || IsXMLType(ct)) {
		hdr.Set(hdrAcceptKey, hdr.Get(hdrContentTypeKey))
	}
  
	// NOTE: 添加到请求中
	r.Header = hdr
	return nil
}
```

#### 4.5.3 parseRequestBody

这一部分对请求体进行处理，同样还是对resty.Request对象的处理，还没有进入http.Request中。

既然是处理请求体，那首先得要求当前请求支持请求体，比如get方法一般就不支持请求体，那么首先要做的就是把这些不支持的请求过滤掉，不进行处理。

这里判断就是通过`isPayloadSupported`函数完成的。`Head/Options`方法不支持请求体，`Get`方法可以通过`AllowGetMethodPayload`进行设置。

```golang
func isPayloadSupported(m string, allowMethodGet bool) bool {
	return !(m == MethodHead || m == MethodOptions || (m == MethodGet && !allowMethodGet))
}
```

如果请求支持设置请求体，接下来进入不同类型数据的处理。首先是multipart类型的处理，真实的处理逻辑位于`handleMultipart`函数中，这里先知道有这么回事儿，将在[文件是怎样上传的](#文件是怎样上传的)这一节进行详细分析。

```golang
if r.isMultiPart && !(r.Method == MethodPatch) {
      if err = handleMultipart(c, r); err != nil {
          return
      }
      goto CL
}
```

所有的类型处理完后，都会跳转到CL这里，按照配置要求决定是否设置请求体长度。

```golang
CL:
	// NOTE: 设置ContentLength
	// by default resty won't set content length, you can if you want to :)
	if (c.setContentLength || r.setContentLength) && r.bodyBuf != nil {
		r.Header.Set(hdrContentLengthKey, fmt.Sprintf("%d", r.bodyBuf.Len()))
	}
	return
```

接下来是FormData类型的处理。

```golang
if len(c.FormData) > 0 || len(r.FormData) > 0 {
  	handleFormData(c, r)
	  goto CL
}
```

真实的逻辑处理是在`handleFormData`函数中，且看下面的源码。

其实原理差不多，可以总结为两步:

1.把client级别和request级别添加的FormData数据添加到url.Values{}结构中；

2.对FormData进行编码后将数据写入到bodyBuf中，并设置请求的`Content-Type`为`application/x-www-form-urlencoded`;

```golang
func handleFormData(c *Client, r *Request) {
	// NOTE: 将client和Request里的formdata保存下来
	formData := url.Values{}

	for k, v := range c.FormData {
		for _, iv := range v {
			formData.Add(k, iv)
		}
	}

	for k, v := range r.FormData {
		// remove form data field from client level by key
		// since overrides happens for that key in the request
		formData.Del(k)

		for _, iv := range v {
			formData.Add(k, iv)
		}
	}

	// NOTE: 写入formdata并设置content-type
	r.bodyBuf = bytes.NewBuffer([]byte(formData.Encode()))
	r.Header.Set(hdrContentTypeKey, formContentType)
	r.isFormData = true
}
```

除了以上两种类型，还支持普通的请求体，比如通过`SetBody`设置的内容，它的过程和上面类似。首先设置请求头中的Content-Type，然后对请求体进行处理后写到bodyBuf中，细节将在[body数据是怎样发送的](#body数据是怎样发送的)中进行分析。

```golang
if r.Body != nil {
      handleContentType(c, r)

      if err = handleRequestBody(c, r); err != nil {
       		 return
      }
}
```

#### 4.5.4 createHTTPRequest

经过上面的钩子处理后，发起一个请求所需的所有数据基本都已经准备好了。

resty库复用了golang标准库提供的http包的能力，现在只需要用准备好的数据(resty.Request)构造一个http.Request对象即可。这就是createHTTPRequest中间件的作用。

首先调用`http.NewRequest`方法创建一个http.Request对象(即RawRequest)。这里进行了一下判断，如果bodyBuf为nil，则根据`SetBody`设置的对象创建，可能是reader，也可能是空结构体。如果bodyBuf不为nil，直接将bodyBuf丢过去。

```golang
if r.bodyBuf == nil {
    if reader, ok := r.Body.(io.Reader); ok {
     	 r.RawRequest, err = http.NewRequest(r.Method, r.URL, reader)
    } else if c.setContentLength || r.setContentLength {
     	 r.RawRequest, err = http.NewRequest(r.Method, r.URL, http.NoBody)
    } else {
     	 r.RawRequest, err = http.NewRequest(r.Method, r.URL, nil)
    }
} else {
	  r.RawRequest, err = http.NewRequest(r.Method, r.URL, r.bodyBuf)
}
```

接下来把Header和cookie复制给RawRequest，基本这一步就结束了。

```golang
// Assign close connection option
r.RawRequest.Close = c.closeConnection

// Add headers into http request
r.RawRequest.Header = r.Header

// Add cookies from client instance into http request
// NOTE: 设置cookie
for _, cookie := range c.Cookies {
 	  r.RawRequest.AddCookie(cookie)
}

// Add cookies from request instance into http request
for _, cookie := range r.Cookies {
  	r.RawRequest.AddCookie(cookie)
}
```

#### 4.5.5 addCredentials

这个钩子主要和身份认证相关。

比如比较简单的`BasicAuth`，如果我们通过`SetBasicAuth`设置了用户名和密码，就会在这里进行处理。

```golang
if r.UserInfo != nil { // takes precedence
    r.RawRequest.SetBasicAuth(r.UserInfo.Username, r.UserInfo.Password)
    isBasicAuth = true
} else if c.UserInfo != nil {
    r.RawRequest.SetBasicAuth(c.UserInfo.Username, c.UserInfo.Password)
    isBasicAuth = true
}
```

除了这种方式，还支持添加`Authorization`头方式的认证。默认的schema为`Bearer`，当然可以自己指定其他的schema。

执行时都是将schema + token设置为Authorization头的内容。

```golang
var authScheme string
if !IsStringEmpty(r.AuthScheme) {
  	authScheme = r.AuthScheme
} else if !IsStringEmpty(c.AuthScheme) {
 	 authScheme = c.AuthScheme
} else {
  	authScheme = "Bearer"
}

// Build the Token Auth header
// NOTE: 设置Authorization头
if !IsStringEmpty(r.Token) { // takes precedence
  	r.RawRequest.Header.Set(c.HeaderAuthorizationKey, authScheme+" "+r.Token)
} else if !IsStringEmpty(c.Token) {
  	r.RawRequest.Header.Set(c.HeaderAuthorizationKey, authScheme+" "+c.Token)
}
```

#### 4.5.6 parseResponseBody

在发请求前可以通过`SetResult`和`SetError`设置正常响应和异常响应的返回对象，执行到这里时会自动进行反序列化，这样我们在使用时就不用自己读取响应体并反序列化了。

当然，有响应体时才需要反序列化，如果是响应码为204(NoContent)，那就不需要执行了。

```golang
// NOTE: 无内容无需解析
if res.StatusCode() == http.StatusNoContent {
  	return
}
```

接下来判断用什么方式反序列化，当前只支持JSON和XML，也就是说只有响应的Content-Type为json或xml时才会自动反序列化到我们设置的结构体中，其他类型我们只能通过body自己反序列化。

对于响应码为(199, 300)之间的响应，会将结果反序列化到Result中；

对于响应码大于399的响应，视为失败，将返回内容反序列化到Error中。

```golang
	// NOTE: 确定解析响应体的content-type，优先级为forceContentType > 响应头的content-type > 请求设置的fallbackContentType
	ct := firstNonEmpty(res.Request.forceContentType, res.Header().Get(hdrContentTypeKey), res.Request.fallbackContentType)
	// NOTE: 只支持json和xml类型的响应自动反序列化
	if IsJSONType(ct) || IsXMLType(ct) {
		// HTTP status code > 199 and < 300, considered as Result
		// NOTE: http状态码在(199,300)之间视为成功的请求
		if res.IsSuccess() {
			res.Request.Error = nil
			if res.Request.Result != nil {
				// NOTE: 利用预先设置的反序列化器将body反序列化到result中
				err = Unmarshalc(c, ct, res.body, res.Request.Result)
				return
			}
		}

		// HTTP status code > 399, considered as Error
		// NOTE: http状态码大于399时视为有错误发生
		if res.IsError() {
			// global error interface
			if res.Request.Error == nil && c.Error != nil {
				res.Request.Error = reflect.New(c.Error).Interface()
			}

			if res.Request.Error != nil {
				err = Unmarshalc(c, ct, res.body, res.Request.Error)
			}
		}
	}
```

### 4.6 文件是怎样上传的?

有关原理可阅读: https://www.cnblogs.com/zyzzz/p/16636555.html

我们可以通过resty.Request对象的`SetFile`方法指定要上传的文件。这里进行了两个设置:

1.将`isMultiPart`设置为true，表示是一个multipart类型的请求；

2.将param和@符号拼接，作为formdata的key，这样做是为了与其他的表单项区分开；

```golang
func (r *Request) SetFile(param, filePath string) *Request {
    // NOTE: 标识是MultiPart请求
    r.isMultiPart = true
    // NOTE: file表单参数名前有添加@与普通表单区分
    r.FormData.Set("@"+param, filePath)
    return r
}
```

在4.5.3中讲到有这么一段代码，这里就连接上了。

```golang
if r.isMultiPart && !(r.Method == MethodPatch) {
      if err = handleMultipart(c, r); err != nil {
          return
      }
      goto CL
}
```

调用`SetFile`方法设置了isMultiPart标记，在createRequestBody中间件中就会进入handleMultipart函数进行处理。主要步骤如下:

1.给bodyBuf初始化，并创建一个指向此结构的Writer。

该Writer中包含一个boundary字段，这个字段对于multipart请求至关重要。

```golang
r.bodyBuf = acquireBuffer()
w := multipart.NewWriter(r.bodyBuf)
// 
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w:        w,
		boundary: randomBoundary(),
	}
}
```

2.将resty.Client中设置的formdata项写入Writer，实际就是写到了bodyBuf中。通过分析WriteField可知，每一个表单项都会设置一个Content-Disposition，相邻的表单项之间会用w.boundary隔开。

```golang
// NOTE: 添加客户端设置的普通表单域，每个表单域有单独的Content-Disposition，且相邻表单域之间用Content-Type头后的随机字符串隔开
for k, v := range c.FormData {
    for _, iv := range v {
        if err = w.WriteField(k, iv); err != nil {
          return err
        }
    }
}
```

3.处理resty.Request级别的formdata。分为两部分:

一部分是普通的表单域，形式上和2差不多，处理方式一致;

另一部分是文件表单域，这类我们在设置时通过@符号标记了，很容易就区分开了，但是处理方式不一样，需要将文件内容根据其路径读出来，然后写入到Writer里面。

```golang
for k, v := range r.FormData {
    for _, iv := range v {
        // TODO: @是哪里添加的?
        if strings.HasPrefix(k, "@") { // file
            // NOTE: 添加文件表单域
            err = addFile(w, k[1:], iv)
            if err != nil {
              	return
            }
        } else { // form value
            // NOTE: 添加普通表单域
            if err = w.WriteField(k, iv); err != nil {
              	return err
            }
        }
    }
}
```

4.处理通过其他函数创建的multipart类型。

```golang
// #21 - adding io.Reader support
// NOTE: multipartFiles保存每个上传的文件的信息，这里为每个文件域添加一个section并写入文件内容
if len(r.multipartFiles) > 0 {
    for _, f := range r.multipartFiles {
        err = addFileReader(w, f)
        if err != nil {
          	return
        }
    }
}

// GitHub #130 adding multipart field support with content type
if len(r.multipartFields) > 0 {
    for _, mf := range r.multipartFields {
        if err = addMultipartFormField(w, mf); err != nil {
         	 return
        }
    }
}
```

5.最后是很重要的一步，将content-type设置为`multipart/form-data; boundary=xxx`，这里的boundary就是前面Writer里的boundary。

```golang
r.Header.Set(hdrContentTypeKey, w.FormDataContentType())

//
func (w *Writer) FormDataContentType() string {
	b := w.boundary
	// We must quote the boundary if it contains any of the
	// tspecials characters defined by RFC 2045, or space.
	if strings.ContainsAny(b, `()<>@,;:\"/[]?= `) {
		b = `"` + b + `"`
	}
	return "multipart/form-data; boundary=" + b
}
```

至此，上传的文件就成功变成了http报文的请求体的一部分了。

### 4.7 body数据是怎样发送的?

4.5.4最后提了一下对于存在body时的处理，如下:

```golang
if r.Body != nil {
    handleContentType(c, r)

    if err = handleRequestBody(c, r); err != nil {
      	return
    }
}
```

首先调用handleContentType设置content-type。判断逻辑是:

1.优先使用header里的content-type；

2.如果header里面没有设置content-type，则根据body数据类型进行判断。如果body是struct/map/Slice类型，则content-type设置为`application/json`；如果是string类型，则设置为`text/plain; charset=utf-8"`；如果都不是，则调用`http.DetectContentType`进行判断。

```golang
func handleContentType(c *Client, r *Request) {
	contentType := r.Header.Get(hdrContentTypeKey)
	if IsStringEmpty(contentType) {
		contentType = DetectContentType(r.Body)
		r.Header.Set(hdrContentTypeKey, contentType)
	}
}

func DetectContentType(body interface{}) string {
	contentType := plainTextType
	kind := kindOf(body)
	switch kind {
	case reflect.Struct, reflect.Map:
		contentType = jsonContentType
	case reflect.String:
		contentType = plainTextType
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = jsonContentType
		}
	}
	return contentType
}
```

确定了content-type之后，剩下没做的就是写入body数据了。由于body有不同类型，写入请求时也就涉及不同类型的处理。这一部分逻辑在`handleRequestBody`中。总结如下:

1.如果`SetBody`时传入的是`io.Reader`类型，直接从该reader读到bodyBuf中；

2.如果是`[]byte`或`string`类型，直接将字节序列写入到bodyBuf中；

3.如果是`json`或`xml`类型，先调用对应的序列化方法转换成字节序列，再写入bodyBuf中；

```golang
func handleRequestBody(c *Client, r *Request) (err error) {
	var bodyBytes []byte
	contentType := r.Header.Get(hdrContentTypeKey)
	kind := kindOf(r.Body)
	r.bodyBuf = nil

	if reader, ok := r.Body.(io.Reader); ok {
		if c.setContentLength || r.setContentLength { // keep backward compatibility
			r.bodyBuf = acquireBuffer()
			_, err = r.bodyBuf.ReadFrom(reader)
			// TODO: why
			r.Body = nil
		} else {
			// Otherwise buffer less processing for `io.Reader`, sounds good.
			return
		}
	} else if b, ok := r.Body.([]byte); ok {
		bodyBytes = b
	} else if s, ok := r.Body.(string); ok {
		bodyBytes = []byte(s)
	} else if IsJSONType(contentType) &&
		(kind == reflect.Struct || kind == reflect.Map || kind == reflect.Slice) {
		r.bodyBuf, err = jsonMarshal(c, r, r.Body)
		if err != nil {
			return
		}
	} else if IsXMLType(contentType) && (kind == reflect.Struct) {
		bodyBytes, err = c.XMLMarshal(r.Body)
		if err != nil {
			return
		}
	}

	if bodyBytes == nil && r.bodyBuf == nil {
		err = errors.New("unsupported 'Body' type/value")
	}

	// if any errors during body bytes handling, return it
	if err != nil {
		return
	}

	// []byte into Buffer
	if bodyBytes != nil && r.bodyBuf == nil {
		r.bodyBuf = acquireBuffer()
		_, _ = r.bodyBuf.Write(bodyBytes)
	}

	return
}
```

## 5.总结

resty提供了一个比较完备的http客户端库，能够很方便地帮助我们简化一些重复性的操作。其在标准库的能力基础上进行拓展，采用了中间件机制，提供了一定的扩展性。通过学习这个库，不仅能够学到一些设计方法和编程范式，也能够了解到有关http请求响应的一些基本原理。当然，由于是在标准库的基础上建立起来的，所以屏蔽了很多细节，后面如果有时间的话，将会专门针对net包出一个系列进行专门的学习。
























