package core

import (
	"fmt"
	"strings"
	"time"
	"math"
	"strconv"

	"github.com/ZeroErrors/go-bedrockping"
)

func BedrockPing(host string) (maxPlayers string, onlinePlayers string, ping float64) {
	port := "19132"
	if strings.Contains(host, ":") {
		splitted := strings.Split(host, ":")
		port = splitted[1]
		host = splitted[0]
	}
	loop:
	start := time.Now()
	resp, err := bedrockping.Query(host+":"+port, time.Second, 150*time.Millisecond)
	ping = (math.Floor(float64(time.Since(start).Nanoseconds())/10000))/100/3.25
	if err != nil {
		fmt.Printf("unable to ping host: %s, error: %s\n", host, err)
		time.Sleep(time.Second)
		goto loop
	}
	maxPlayers = strconv.Itoa(resp.MaxPlayers)
	onlinePlayers = strconv.Itoa(resp.PlayerCount)
	return
}