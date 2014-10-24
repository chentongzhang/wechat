// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"github.com/magicshui/wechat/mp/merchant/shelf"
)

// 增加货架
//  NOTE: 无需指定 Id 字段
func (c *Client) MerchantShelfAdd(_shelf *shelf.Shelf) (shelfId int64, err error) {
	if _shelf == nil {
		err = errors.New("_shelf == nil")
		return
	}

	var shelfx struct {
		Id     int64  `json:"shelf_id,omitempty"`
		Name   string `json:"shelf_name"`
		Banner string `json:"shelf_banner,omitempty"`
		Info   struct {
			ModuleInfos []shelf.Module `json:"module_infos,omitempty"`
		} `json:"shelf_data"`
	}

	// shelfx.Id = _shelf.Id
	shelfx.Name = _shelf.Name
	shelfx.Banner = _shelf.Banner
	shelfx.Info.ModuleInfos = _shelf.Info.ModuleInfos

	var result struct {
		Error
		ShelfId int64 `json:"shelf_id"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantShelfAddURL(token)
	if err = c.postJSON(_url, &shelfx, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		shelfId = result.ShelfId
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

// 删除货架
func (c *Client) MerchantShelfDelete(shelfId int64) (err error) {
	var request = struct {
		ShelfId int64 `json:"shelf_id"`
	}{
		ShelfId: shelfId,
	}

	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantShelfDeleteURL(token)
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result
		return
	}
}

// 修改货架
func (c *Client) MerchantShelfModify(_shelf *shelf.Shelf) (err error) {
	if _shelf == nil {
		return errors.New("_shelf == nil")
	}

	var shelfx struct {
		Id     int64  `json:"shelf_id,omitempty"`
		Name   string `json:"shelf_name"`
		Banner string `json:"shelf_banner,omitempty"`
		Info   struct {
			ModuleInfos []shelf.Module `json:"module_infos,omitempty"`
		} `json:"shelf_data"`
	}

	shelfx.Id = _shelf.Id
	shelfx.Name = _shelf.Name
	shelfx.Banner = _shelf.Banner
	shelfx.Info.ModuleInfos = _shelf.Info.ModuleInfos

	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantShelfModifyURL(token)
	if err = c.postJSON(_url, &shelfx, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result
		return
	}
}

// 获取所有货架
func (c *Client) MerchantShelfGetAll() (shelves []shelf.Shelf, err error) {
	var result = struct {
		Error
		Shelves []shelf.Shelf `json:"shelves"`
	}{
		Shelves: make([]shelf.Shelf, 0, 16),
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantShelfGetAllURL(token)
	if err = c.getJSON(_url, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		shelves = result.Shelves
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

// 根据货架ID获取货架信息
func (c *Client) MerchantShelfGetById(shelfId int64) (_shelf *shelf.Shelf, err error) {
	var request = struct {
		ShelfId int64 `json:"shelf_id"`
	}{
		ShelfId: shelfId,
	}

	var result struct {
		Error
		shelf.Shelf
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantShelfGetByIdURL(token)
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		_shelf = &result.Shelf
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
