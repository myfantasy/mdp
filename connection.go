package mdp

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/valyala/fasthttp"
)

// Connection connection to myfdb
type Connection struct {
	Servers      []string
	Token        string
	QueryTimeout time.Duration
	QueryWait    time.Duration
	client       *fasthttp.Client
}

// CreateConnection Create connection to
func CreateConnection(servers []string, token string, maxConnsPerHost int, maxIdleConnDuration time.Duration,
	queryTimeout time.Duration, queryWait time.Duration) *Connection {

	c := &Connection{
		Servers:      servers,
		QueryTimeout: queryTimeout,
		Token:        token,
		QueryWait:    queryWait,
		client: &fasthttp.Client{
			MaxConnsPerHost:     maxConnsPerHost,
			MaxIdleConnDuration: maxIdleConnDuration,
		},
	}

	return c
}

// ConnectionFileStruct struct connection from file or json
type ConnectionFileStruct struct {
	Server  *string  `json:"server,omitempty"`
	Servers []string `json:"servers,omitempty"`
	Token   string   `json:"token"`

	QueryTimeout time.Duration `json:"query_timeout,omitempty"`
	QueryWait    time.Duration `json:"query_wait,omitempty"`

	MaxConnsPerHost     int           `json:"max_conns_per_host,omitempty"`
	MaxIdleConnDuration time.Duration `json:"max_idle_conn_duration,omitempty"`
}

// ConnectionGetFromConnectionFileStruct get ConnectionFileStruct from json and
func ConnectionGetFromConnectionFileStruct(cfs ConnectionFileStruct) *Connection {
	c := &Connection{
		Servers:      cfs.Servers,
		Token:        cfs.Token,
		QueryTimeout: 4 * time.Second,
		QueryWait:    5 * time.Second,
		client: &fasthttp.Client{
			MaxConnsPerHost:     5,
			MaxIdleConnDuration: 600 * time.Second,
		},
	}
	if cfs.Server != nil {
		c.Servers = append(c.Servers, *cfs.Server)
	}

	if cfs.MaxConnsPerHost > 0 {
		c.client.MaxConnsPerHost = cfs.MaxConnsPerHost
	}
	if cfs.MaxIdleConnDuration > 0 {
		c.client.MaxIdleConnDuration = cfs.MaxIdleConnDuration
	}

	if cfs.QueryTimeout > 0 {
		c.QueryTimeout = cfs.QueryWait
	}

	if cfs.QueryWait > 0 {
		c.QueryWait = cfs.QueryWait
	}

	return c
}

// ConnectionGetFromJSON get Connection from json
func ConnectionGetFromJSON(d []byte) (*Connection, error) {

	var cfs ConnectionFileStruct
	err := json.Unmarshal(d, &cfs)
	if err != nil {
		return nil, err
	}

	return ConnectionGetFromConnectionFileStruct(cfs), nil
}

// ConnectionGetFromFile get Connection from file with json
func ConnectionGetFromFile(path string) (*Connection, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return ConnectionGetFromJSON(data)
}

// DoQuery do some query
func (c *Connection) DoQuery(path string, query []byte) (body []byte, statusCode int, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(c.Servers[0] + path)
	req.SetBody(query)
	req.Header.SetMethod("POST")
	req.Header.Set("Token", c.Token)

	err = c.client.DoTimeout(req, resp, c.QueryWait)
	if err != nil {
		return body, 0, err
	}

	body = resp.Body()

	return body, resp.StatusCode(), err
}

// DoQueryObject do some query
func (c *Connection) DoQueryObject(path string, v interface{}) (body []byte, statusCode int, err error) {
	query, err := json.Marshal(v)

	if err != nil {
		return body, 0, err
	}

	return c.DoQuery(path, query)
}
