package request

// 点击菜单拉取消息时的事件推送
type CardPassCheckEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event  string `xml:"Event"    json:"Event"` // 事件类型，CLICK
	CardId string `xml:"CardId" json:"CardId"`  // 事件KEY值，与自定义菜单接口中KEY值对应
}

func (req *Request) CardPassCheckEvent() (event *CardPassCheckEvent) {
	event = &CardPassCheckEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		CardId:     req.CardId,
	}
	return
}

type UserGetCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event          string `xml:"Event"    json:"Event"` // 事件类型，CLICK
	CardId         string `xml:"CardId" json:"CardId"`  // 事件KEY值，与自定义菜单接口中KEY值对应
	IsGiveByFriend int    `xml:"IsGiveByFriend" json:"IsGiveByFriend"`
	UserCardCode   int64  `xml:"UserCardCode" json:"UserCardCode"`
}

func (req *Request) UserGetCardEvent() (event *UserGetCardEvent) {
	event = &UserGetCardEvent{
		CommonHead:     req.CommonHead,
		Event:          req.Event,
		CardId:         req.CardId,
		IsGiveByFriend: req.IsGiveByFriend,
		UserCardCode:   req.UserCardCode,
	}
	return
}

type UserDelCard struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event        string `xml:"Event"    json:"Event"` // 事件类型，CLICK
	CardId       string `xml:"CardId" json:"CardId"`  // 事件KEY值，与自定义菜单接口中KEY值对应
	UserCardCode int64  `xml:"UserCardCode" json:"UserCardCode"`
}

func (req *Request) UserDelCard() (event *UserDelCard) {
	event = &UserDelCard{
		CommonHead:   req.CommonHead,
		Event:        req.Event,
		CardId:       req.CardId,
		UserCardCode: req.UserCardCode,
	}
	return
}
