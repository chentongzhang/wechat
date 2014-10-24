// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"github.com/magicshui/wechat/mp/message/passive/request"
	"net/http"
)

// 公众号对外暴露的接口
type Agent interface {
	GetId() string    // 获取公众号的原始ID, 等于后台中的 公众号设置-->帐号详情-->原始ID
	GetToken() string // 获取公众号的 Token, 和后台中的设置相等

	// fuck, AppId 貌似需要认证才会有的???
	// 如果不知道自己的 AppId 是多少, 可以先随便填入一个字符串,
	// 这样正常情况下会出现 AppId mismatch 错误, 错误中 have 后面的就是正确的 AppId
	GetAppId() string
	GetLastAESKey() [32]byte    // 获取最后一个有效的 AES 加密 Key
	GetCurrentAESKey() [32]byte // 获取当前有效的 AES 加密 Key

	// 明文模式 情况下需要实现的方法
	// 未知类型的消息处理方法
	//  rawXMLMsg 是 xml 消息体
	//  timestamp 是请求 URL 中的时间戳
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 user-agent
	ServeUnknownMsg(w http.ResponseWriter, r *http.Request, rawXMLMsg []byte, timestamp int64)

	// 明文模式 情况下需要实现的方法
	// 消息处理函数
	//  rawXMLMsg 是 xml 消息体
	//  timestamp 是请求 URL 中的时间戳
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 user-agent
	ServeTextMsg(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64)
	ServeImageMsg(w http.ResponseWriter, r *http.Request, msg *request.Image, rawXMLMsg []byte, timestamp int64)
	ServeVoiceMsg(w http.ResponseWriter, r *http.Request, msg *request.Voice, rawXMLMsg []byte, timestamp int64)
	ServeVideoMsg(w http.ResponseWriter, r *http.Request, msg *request.Video, rawXMLMsg []byte, timestamp int64)
	ServeLocationMsg(w http.ResponseWriter, r *http.Request, msg *request.Location, rawXMLMsg []byte, timestamp int64)
	ServeLinkMsg(w http.ResponseWriter, r *http.Request, msg *request.Link, rawXMLMsg []byte, timestamp int64)

	// 明文模式 情况下需要实现的方法
	// 事件处理函数
	//  rawXMLMsg 是 xml 消息体
	//  timestamp 是请求 URL 中的时间戳
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 user-agent
	ServeSubscribeEvent(w http.ResponseWriter, r *http.Request, msg *request.SubscribeEvent, rawXMLMsg []byte, timestamp int64)
	ServeUnsubscribeEvent(w http.ResponseWriter, r *http.Request, msg *request.UnsubscribeEvent, rawXMLMsg []byte, timestamp int64)
	ServeSubscribeByScanEvent(w http.ResponseWriter, r *http.Request, msg *request.SubscribeByScanEvent, rawXMLMsg []byte, timestamp int64)
	ServeScanEvent(w http.ResponseWriter, r *http.Request, msg *request.ScanEvent, rawXMLMsg []byte, timestamp int64)
	ServeLocationEvent(w http.ResponseWriter, r *http.Request, msg *request.LocationEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuClickEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuClickEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuViewEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuViewEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuScanCodePushEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodePushEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuScanCodeWaitMsgEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodeWaitMsgEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuPicSysPhotoEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicSysPhotoEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuPicPhotoOrAlbumEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicPhotoOrAlbumEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuPicWeixinEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicWeixinEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuLocationSelectEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuLocationSelectEvent, rawXMLMsg []byte, timestamp int64)
	ServeMassSendJobFinishEvent(w http.ResponseWriter, r *http.Request, msg *request.MassSendJobFinishEvent, rawXMLMsg []byte, timestamp int64)
	ServeTemplateSendJobFinishEvent(w http.ResponseWriter, r *http.Request, msg *request.TemplateSendJobFinishEvent, rawXMLMsg []byte, timestamp int64)
	ServeMerchantOrderEvent(w http.ResponseWriter, r *http.Request, msg *request.MerchantOrderEvent, rawXMLMsg []byte, timestamp int64)

	// 兼容模式, 安全模式 情况下需要实现的方法
	// 未知类型的消息处理方法
	//  rawXMLMsg   是解密后的"明文" xml 消息体
	//  timestamp   是请求 URL 中的时间戳
	//  nonce       是请求 URL 中的随机数
	//  AESKey      是微信"当前"消息加密所用的 AES key
	//  random      是请求 http body 中的密文消息加密时所用的 random
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 user-agent
	ServeAESUnknownMsg(w http.ResponseWriter, r *http.Request, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)

	// 兼容模式, 安全模式 情况下需要实现的方法
	// 消息处理函数
	//  rawXMLMsg   是解密后的"明文" xml 消息体
	//  timestamp   是请求 URL 中的时间戳
	//  nonce       是请求 URL 中的随机数
	//  AESKey      是微信"当前"消息加密所用的 AES key
	//  random      是请求 http body 中的密文消息加密时所用的 random
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 user-agent
	ServeAESTextMsg(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESImageMsg(w http.ResponseWriter, r *http.Request, msg *request.Image, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESVoiceMsg(w http.ResponseWriter, r *http.Request, msg *request.Voice, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESVideoMsg(w http.ResponseWriter, r *http.Request, msg *request.Video, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESLocationMsg(w http.ResponseWriter, r *http.Request, msg *request.Location, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESLinkMsg(w http.ResponseWriter, r *http.Request, msg *request.Link, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)

	// 兼容模式, 安全模式 情况下需要实现的方法
	// 事件处理函数
	//  rawXMLMsg   是解密后的"明文" xml 消息体
	//  timestamp   是请求 URL 中的时间戳
	//  nonce       是请求 URL 中的随机数
	//  AESKey      是微信"当前"消息加密所用的 AES key
	//  random      是请求 http body 中的密文消息加密时所用的 random
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 user-agent
	ServeAESSubscribeEvent(w http.ResponseWriter, r *http.Request, msg *request.SubscribeEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESUnsubscribeEvent(w http.ResponseWriter, r *http.Request, msg *request.UnsubscribeEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESSubscribeByScanEvent(w http.ResponseWriter, r *http.Request, msg *request.SubscribeByScanEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESScanEvent(w http.ResponseWriter, r *http.Request, msg *request.ScanEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESLocationEvent(w http.ResponseWriter, r *http.Request, msg *request.LocationEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESMenuClickEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuClickEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESMenuViewEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuViewEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESMenuScanCodePushEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodePushEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESMenuScanCodeWaitMsgEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodeWaitMsgEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESMenuPicSysPhotoEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicSysPhotoEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESMenuPicPhotoOrAlbumEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicPhotoOrAlbumEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESMenuPicWeixinEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicWeixinEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESMenuLocationSelectEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuLocationSelectEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESMassSendJobFinishEvent(w http.ResponseWriter, r *http.Request, msg *request.MassSendJobFinishEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESTemplateSendJobFinishEvent(w http.ResponseWriter, r *http.Request, msg *request.TemplateSendJobFinishEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
	ServeAESMerchantOrderEvent(w http.ResponseWriter, r *http.Request, msg *request.MerchantOrderEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte)
}
