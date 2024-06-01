package main

import (
	"log"
	"os"

	"github.com/eatmoreapple/openwechat"
)

func main() {
	fd, err := initLog()
	if err != nil {
		log.Fatalf("init log failed:%+v", err)
	}
	defer fd.Close()

	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.MessageHandler = messageHandler
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()

	if err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		log.Fatal(err)
	}

	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Fatal(err)
	}

	go schedule(self)
	bot.Block()
}

func initLog() (*os.File, error) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	fd, err := os.OpenFile("wechat.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	log.SetOutput(fd)
	return fd, nil
}
