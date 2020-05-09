package icapclient

import (
	"strconv"
)

// Client represents the icap client who makes the icap server calls
type Client struct {
	scktDriver *Driver
}

// Do makes the call
func (c *Client) Do(req *Request) (*Response, error) {

	port, err := strconv.Atoi(req.URL.Port())

	if err != nil {
		return nil, err
	}

	c.scktDriver = NewDriver(req.URL.Hostname(), port)

	if err := c.scktDriver.Connect(); err != nil {
		return nil, err
	}

	defer c.scktDriver.Close()

	d, err := DumpRequest(req)

	if err != nil {
		return nil, err
	}

	if err := c.scktDriver.Send(d); err != nil {
		return nil, err
	}

	resp, err := c.scktDriver.Receive()

	if err != nil {
		return nil, err
	}

	return resp, nil
}