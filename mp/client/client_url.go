// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"net/url"
)

// https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
func messageCustomSendURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN
func messageTemplateSendURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
func messageMassSendByGroupURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
func messageMassSendByOpenIdURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
func messageMassDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com//cgi-bin/message/mass/delete?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
func menuCreateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN
func menuGetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
func menuDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=" +
		accesstoken
}

// http://file.api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
func mediaUploadURL(accesstoken string, mediatype string) string {
	return "http://file.api.weixin.qq.com/cgi-bin/media/upload?access_token=" +
		accesstoken +
		"&type=" +
		mediatype
}

// http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
func mediaDownloadURL(accesstoken string, mediaid string) string {
	return "http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=" +
		accesstoken +
		"&media_id=" +
		mediaid
}

// https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
func mediaCreateNewsURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=" +
		accesstoken
}

// https://file.api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=ACCESS_TOKEN
func mediaCreateVideoURL(accesstoken string) string {
	return "https://file.api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
func qrcodeCreateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" +
		accesstoken
}

// https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=TICKET
func qrcodeURL(ticket string) string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" +
		url.QueryEscape(ticket)
}

// https://api.weixin.qq.com/cgi-bin/groups/create?access_token=ACCESS_TOKEN
func userGroupCreateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/create?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/get?access_token=ACCESS_TOKEN
func userGroupGetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/update?access_token=ACCESS_TOKEN
func userGroupRenameURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=ACCESS_TOKEN
func userInWhichGroupURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/getid?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=ACCESS_TOKEN
func userMoveToGroupURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/groups/members/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
func userInfoURL(accesstoken, openid, lang string) string {
	return "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" +
		accesstoken +
		"&openid=" +
		openid +
		"&lang=" +
		lang
}

// https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID
func userGetURL(accesstoken, nextOpenId string) string {
	if nextOpenId == "" {
		return "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" +
			accesstoken
	}
	return "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" +
		accesstoken +
		"&next_openid=" +
		nextOpenId
}

// https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=ACCESS_TOKEN
func userUpdateRemarkURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/customservice/getrecord?access_token=ACCESS_TOKEN
func customServiceRecordGetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/customservice/getrecord?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token= ACCESS_TOKEN
func customServiceKFListURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token= ACCESS_TOKEN
func customServiceOnlineKFListURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/cgi-bin/shorturl?access_token=ACCESS_TOKEN
func shortURLURL(accesstoken string) string {
	return "https://api.weixin.qq.com/cgi-bin/shorturl?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/create?access_token=ACCESS_TOKEN
func cardCreateUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/create?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/getcolors?access_token=TOKEN
func cardGetColorsUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/getcolors?access_token=" +
		accesstoken

}

// https://api.weixin.qq.com/card/code/consume?access_token=TOKEN
func cardCodeConsumeUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/code/consume?access_token=" +
		accesstoken
}
func cardCodeDecryptUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/code/decrypt?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/delete?access_token=TOKEN
func cardDeleteUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/delete?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/code/get?access_token=TOKEN
func cardCodeGetUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/code/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/batchget?access_token=TOKEN
func cardBatchGetUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/batchget?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/get?access_token=TOKEN
func cardGetUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/code/update?access_token=TOKEN
func cardCodeUpdateUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/code/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/code/unavailable?access_token=TOKEN
func cardCodeUnavailableUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/code/unavailable?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/update?access_token=TOKEN
func cardUpdateUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/code/consume?access_token=TOKEN
func cardLocationBatchAddUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/location/batchadd?access_token=" +
		accesstoken
}

//
func cardLocationBatchGetUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/location/batchget?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/location/batchget?access_token=TOKEN
func cardTestWhiteListSetUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/location/batchget?access_token=" +
		accesstoken
}
