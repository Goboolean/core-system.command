package entity

type Client struct {
	response chan<- *TaskAllocResponse
	config   *Config
}

func NewClient(response chan<- *TaskAllocResponse, config *Config) *Client {
	return &Client{
		response: response,
		config:   config,
	}
}

func (c *Client) Send(response *TaskAllocResponse, buy bool) {
	if c.config.Channel == true {
		// 채널이 닫혀있는지 확인하는 과정이 필요함
		c.response <- response
	}
	if c.config.Notification == true && buy == true {
		// send notification
	}
	if c.config.AutoTrading == true && buy == true {
		// send auto trading
	}
}

func (c *Client) Close() {
	close(c.response)
}
