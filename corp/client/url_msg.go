// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

// https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=ACCESS_TOKEN
func _MsgSendURL(accesstoken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + accesstoken
}
