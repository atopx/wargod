package core

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log/slog"
	"net/http"
	"time"
)

// Client 是自定义的HTTP客户端结构体
type Client struct {
	ctx        context.Context
	httpClient *http.Client
	dialer     *websocket.Dialer
	wsClient   *websocket.Conn
	baseURL    string
	auth       string
	handles    []*handle
}

func New(ctx context.Context, timeout time.Duration, ga *GameAuth) (*Client, error) {
	// 构建HTTP客户端
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		Timeout: timeout, // 设置超时时间
	}

	// 创建客户端实例
	client := &Client{
		ctx:        ctx,
		httpClient: httpClient,
		dialer:     &websocket.Dialer{TLSClientConfig: tlsConfig},
		auth:       ga.BasicHeader(),
		baseURL:    fmt.Sprintf("https://127.0.0.1:%s", ga.Port),
	}
	if err := client.refresh(ga); err != nil {
		return nil, fmt.Errorf("init client refresh error: %w", err)
	}

	return client, nil
}

// Request 创建HTTP请求并设置Basic Auth和默认Header
func (c *Client) Request(method, endpoint string, body []byte) ([]byte, error) {
	url := c.baseURL + endpoint
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", c.auth)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s request %s failed: %w", method, url, err)
	}
	defer func() { _ = resp.Body.Close() }()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to %s %s read response body: %w", method, url, err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("%s request %s failed with status: %d, resp: %s", method, url, resp.StatusCode, data)
	}
	return data, nil
}

// Get 发送GET请求
func (c *Client) Get(endpoint string) ([]byte, error) {
	return c.Request(http.MethodGet, endpoint, nil)
}

// Post 发送POST请求
func (c *Client) Post(endpoint string, body []byte) ([]byte, error) {
	return c.Request(http.MethodPost, endpoint, body)
}

// Put 发送PUT请求
func (c *Client) Put(endpoint string, body []byte) ([]byte, error) {
	return c.Request(http.MethodPut, endpoint, body)
}

// Delete 发送DELETE请求
func (c *Client) Delete(endpoint string) error {
	_, err := c.Request(http.MethodDelete, endpoint, nil)
	return err
}

func (c *Client) refresh(ga *GameAuth) error {
	c.auth = ga.BasicHeader()
	c.baseURL = fmt.Sprintf("https://127.0.0.1:%s", ga.Port)
	header := http.Header{"Authorization": {c.auth}}
	conn, _, err := c.dialer.Dial(fmt.Sprintf("wss://127.0.0.1:%s", ga.Port), header)
	if err != nil {
		return fmt.Errorf("websocket refresh token dial: %w", err)
	}
	c.wsClient = conn
	return nil
}

func (c *Client) Subscribe(events map[string]func(data []byte) error) error {
	c.handles = make([]*handle, 0, len(events))
	for k, f := range events {
		handler := newHandler(k, f)
		c.handles = append(c.handles, handler)
		slog.Debug("register event", handler.name)
		if err := c.wsClient.WriteMessage(websocket.TextMessage, handler.subscribe); err != nil {
			return fmt.Errorf("riot listerner subscribe failed: %w", err)
		}
		_, _, _ = c.wsClient.ReadMessage()
	}
	return nil
}

func (c *Client) Listen() {
	for {
		_, message, err := c.wsClient.ReadMessage()
		if err != nil {
			slog.Error("event read message failed", slog.String("error", err.Error()))
			break
		}

		var matched bool
		for _, h := range c.handles {
			if bytes.HasPrefix(message, h.prefix) {
				matched = true
				if err = h.handler(message[len(h.prefix)+1 : len(message)-1]); err != nil {
					slog.Error("event handle error", slog.String("name", h.name), slog.String("error", err.Error()))
				}
				break
			}
		}
		if !matched {
			fmt.Printf("unregister event: %s\n", message)
		}
	}
}
