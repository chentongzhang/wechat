// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package customservice

const (
	RecordPageSizeLimit = 1000 // 客户聊天记录每页最多拉取1000条
)

const (
	OnlineKFInfoStatusPC          = 1
	OnlineKFInfoStatusMobile      = 2
	OnlineKFInfoStatusPCAndMobile = 3
)
