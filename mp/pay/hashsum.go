// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"crypto/sha1"
)

type hashSumFunc func([]byte) []byte

func sha1Sum(data []byte) (sum []byte) {
	arr := sha1.Sum(data)
	sum = arr[:]
	return
}
