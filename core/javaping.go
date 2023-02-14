package core

import (
	"fmt"
	"math"
	"time"

	"github.com/Tnze/go-mc/bot"
	"github.com/tidwall/gjson"
)

func JavaPing(host string) (maxPlayers string, onlinePlayers string, ping float64) {
	loop:
	resp,jew,err := bot.PingAndList(host)
	if err != nil {
		fmt.Printf("unable to ping host: %s, error: %s\n",host,err)
		time.Sleep(time.Second)
		goto loop
	}
	maxPlayers= gjson.Get(string(resp),"players.max").String()
	onlinePlayers= gjson.Get(string(resp),"players.online").String()
	ping = (math.Floor(float64(jew.Nanoseconds())/10000))/100
	return
}