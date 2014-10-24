// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	wechatjson "github.com/magicshui/wechat/json"
	"net/http"
)

// Client 封装了主动请求功能, 比如创建菜单, 回复客服消息等等
type Client struct {
	tokenService TokenService
	httpClient   *http.Client
}

// 创建一个新的 Client.
//  如果 httpClient == nil 则默认用 http.DefaultClient,
//  see github.com/magicshui/wechat/CommonHttpClient 和
//      github.com/magicshui/wechat/MediaHttpClient
func NewClient(tokenService TokenService, httpClient *http.Client) (clt *Client) {
	if tokenService == nil {
		panic("tokenService == nil")
	}

	clt = &Client{
		tokenService: tokenService,
	}

	if httpClient == nil {
		clt.httpClient = http.DefaultClient
	} else {
		clt.httpClient = httpClient
	}

	return
}

// 获取 access token
// 正常情况下 token != "" && err == nil, 否则 token == "" && err != nil
func (c *Client) Token() (token string, err error) {
	return c.tokenService.Token()
}

// 从微信服务器获取新的 access token.
// 正常情况下 token != "" && err == nil, 否则 token == "" && err != nil
//  NOTE:
//  1. 一般情况下无需调用该函数, 请使用 Token() 获取 access token.
//  2. 即使微信服务器返回了 access token 过期错误(错误代码 42001, 正常情况下不会出现),
//     也请谨慎调用 TokenRefresh, 建议直接返回错误! 因为很有可能高并发情况下造成雪崩效应!
//  3. 再次强调, 调用这个函数你应该知道发生了什么!!!
func (c *Client) TokenRefresh() (token string, err error) {
	return c.tokenService.TokenRefresh()
}

// Client 通用的 json post 请求
func (c *Client) postJSON(_url string, request interface{}, response interface{}) (err error) {
	buf := textBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
	buf.Reset()                                 // important
	defer textBufferPool.Put(buf)

	if err = wechatjson.NewEncoder(buf).Encode(request); err != nil {
		return
	}

	resp, err := c.httpClient.Post(_url, "application/json; charset=utf-8", buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return
	}

	return
}

// Client 通用的 json get 请求
func (c *Client) getJSON(_url string, response interface{}) (err error) {
	resp, err := c.httpClient.Get(_url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return
	}

	return
}
