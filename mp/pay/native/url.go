// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package native

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
)

// 生成 native 支付 URL.
//  @appId:      必须, 公众号 id
//  @nonceStr:   必须, 32个字符以内, 商户生成的随机字符串
//  @timestamp:  必须, unixtime, 商户生成
//  @productId:  必须, 32个字符以内, 商户需要定义并维护自己的商品id, 这个id与一张订单等价, 微信后台凭借该id通过POST商户后台获取交易必须信息;
//  @paySignKey: 必须, 公众号支付请求中用于加密的密钥 Key
func NativeURL(appId, nonceStr string, timestamp int64,
	productId, paySignKey string) string {

	TimeStampStr := strconv.FormatInt(timestamp, 10)

	const keysLen = len(`appid=&appkey=&noncestr=&productid=&timestamp=`)
	n := keysLen + len(appId) + len(paySignKey) + len(nonceStr) +
		len(productId) + len(TimeStampStr)

	string1 := make([]byte, 0, n)

	// 字典序
	// appid
	// appkey
	// noncestr
	// productid
	// timestamp
	string1 = append(string1, "appid="...)
	string1 = append(string1, appId...)
	string1 = append(string1, "&appkey="...)
	string1 = append(string1, paySignKey...)
	string1 = append(string1, "&noncestr="...)
	string1 = append(string1, nonceStr...)
	string1 = append(string1, "&productid="...)
	string1 = append(string1, productId...)
	string1 = append(string1, "&timestamp="...)
	string1 = append(string1, TimeStampStr...)

	hashsumArray := sha1.Sum(string1)
	signature := hex.EncodeToString(hashsumArray[:])

	return "weixin://wxpay/bizpayurl?appid=" + appId +
		"&noncestr=" + nonceStr +
		"&productid=" + productId +
		"&sign=" + signature +
		"&timestamp=" + TimeStampStr
}
