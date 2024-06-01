package main

import (
	"log"
	"strings"

	"github.com/eatmoreapple/openwechat"
)

func messageHandler(msg *openwechat.Message) {
	if msg.IsText() {
		procText(msg)
	} else if msg.IsFriendAdd() {
		msg.Agree()
		log.Printf("auto agree to user %s", msg.FromUserName)
	} else {
		log.Printf("unprocessed messages: %+v", msg)
	}
}

func procText(msg *openwechat.Message) {
	log.Printf("from:%s text: %s", msg.FromUserName, msg.Content)

	if strings.Contains(msg.Content, "入群") || strings.Contains(msg.Content, "1") {
		autoIntoGroup(msg)
	} else {
		msg.ReplyText("不识别的命令")
	}
}

func autoIntoGroup(msg *openwechat.Message) {
	self := msg.Owner()
	friends, err := self.Friends(true)
	if err != nil {
		log.Printf("get friends err %+v", err)
		return
	}

	groups, err := self.Groups(true)
	if err != nil {
		log.Printf("get groups err %+v", err)
		return
	}

	friend := friends.SearchByUserName(1, msg.FromUserName)
	if len(friend) != 1 {
		log.Printf("find friend err %+v", friend)
		return
	}

	group := groups.SearchByNickName(1, "CN DOTA")
	if len(group) != 1 {
		log.Printf("find group err %+v", group)
		return
	}

	err = friend[0].AddIntoGroup(group[0])
	if err != nil {
		log.Printf("add %s to group %s err %+v", friend[0], group[0], err)
		return
	}
	log.Printf("add %s to group %s ok", friend[0], group[0])
}
