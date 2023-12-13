package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/artyomtugaryov/vpnbot/pkg/utils/errors"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func (c *Client) Updates(offset, limit int) ([]Updates, error) {
	values := url.Values{}
	values.Add("offset", strconv.Itoa(offset))
	values.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, values)
	if err != nil {
		return nil, errors.Wrap("Cannot do the getUpdates request", err)
	}

	var result UpdatesResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.Wrap("Cannot parse results", err)
	}

	return result.Result, nil
}

func (c *Client) SendMessages(chatID int, text string) error {
	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chatID))
	values.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, values)
	if err != nil {
		return errors.Wrap("Cannot send message", err)
	}
	return nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	url_ := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}
	request, err := http.NewRequest(http.MethodGet, url_.String(), nil)
	if err != nil {
		return nil, errors.Wrap("Cannot create a request", err)
	}

	request.URL.RawQuery = query.Encode()

	response, err := c.client.Do(request)

	if err != nil {
		return nil, errors.Wrap("Cannot do a request", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap("Cannot parse a response", err)
	}
	return body, nil
}

func newBasePath(token string) string {
	return "bot" + token
}
