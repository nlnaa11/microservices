package loms

type Client struct {
	url     string
	urlGoal string
}

func New(url string, pathTo string) *Client {
	return &Client{
		url:     url,
		urlGoal: url + pathTo,
	}
}
