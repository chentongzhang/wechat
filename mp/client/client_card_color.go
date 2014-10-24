package client

import (
	"github.com/magicshui/wechat/mp/card"
)

func (c *Client) CardGetColors() (_colors []card.Color, err error) {
	var result struct {
		Colors []card.Color `json:"colors"`
		Error
	}
	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := cardGetColorsUrl(token)
	if err = c.postJSON(_url, nil, &result); err != nil {
		return
	}
	switch result.ErrCode {
	case errCodeOK:
		_colors = result.Colors
		return
	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return

	}

}
