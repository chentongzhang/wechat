// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"github.com/magicshui/wechat/corp/message/passive/request"
	"net/http"
)

var _ Agent = new(DefaultAgent)

type DefaultAgent struct {
	CorpId  string
	AgentId int64
	Token   string
	AESKey  [32]byte
}

func (this *DefaultAgent) Init(CorpId string, AgentId int64, Token string, AESKey []byte) {
	if len(AESKey) != 32 {
		panic("the length of AESKey must equal to 32")
	}
	this.CorpId = CorpId
	this.AgentId = AgentId
	this.Token = Token
	copy(this.AESKey[:], AESKey)
}

func (this *DefaultAgent) GetCorpId() string {
	return this.CorpId
}
func (this *DefaultAgent) GetAgentId() int64 {
	return this.AgentId
}
func (this *DefaultAgent) GetToken() string {
	return this.Token
}
func (this *DefaultAgent) GetAESKey() [32]byte {
	return this.AESKey
}

func (this *DefaultAgent) ServeUnknownMsg(w http.ResponseWriter, r *http.Request, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {
}

func (this *DefaultAgent) ServeTextMsg(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {
}
func (this *DefaultAgent) ServeImageMsg(w http.ResponseWriter, r *http.Request, msg *request.Image, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {
}
func (this *DefaultAgent) ServeVoiceMsg(w http.ResponseWriter, r *http.Request, msg *request.Voice, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {
}
func (this *DefaultAgent) ServeVideoMsg(w http.ResponseWriter, r *http.Request, msg *request.Video, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {
}
func (this *DefaultAgent) ServeLocationMsg(w http.ResponseWriter, r *http.Request, msg *request.Location, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {
}

func (this *DefaultAgent) ServeSubscribeEvent(w http.ResponseWriter, r *http.Request, event *request.SubscribeEvent, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {
}
func (this *DefaultAgent) ServeUnsubscribeEvent(w http.ResponseWriter, r *http.Request, event *request.UnsubscribeEvent, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {
}
func (this *DefaultAgent) ServeLocationEvent(w http.ResponseWriter, r *http.Request, event *request.LocationEvent, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {
}
func (this *DefaultAgent) ServeMenuClickEvent(w http.ResponseWriter, r *http.Request, event *request.MenuClickEvent, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {
}
func (this *DefaultAgent) ServeMenuViewEvent(w http.ResponseWriter, r *http.Request, event *request.MenuViewEvent, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {
}
