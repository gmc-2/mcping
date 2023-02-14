package main

import (
	"fmt"
	"mcping/core"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var pingcolor string
	pingtype := "java"
	if len(os.Args) <= 1 {
		fmt.Printf("usage: %s <host> (type)\n<>: required\n(): optional\nallowed types: java (default) and bedrock\n", filepath.Base(os.Args[0]))
		os.Exit(0)
	}
	host := os.Args[1]
	providerhost := host
	if len(os.Args) > 2 {
		pingtype = os.Args[2]
	}
	if strings.Contains(providerhost, ":") {
		split := strings.Split(providerhost, ":")
		providerhost = split[0]
	}
	hosting, country, region := core.GetHostingProvider(providerhost)
	if pingtype == "java" {
		for {
			start := time.Now()
			maxPlayers, onlinePlayers, ping := core.JavaPing(host)
			if ping < 100 {
				pingcolor = "0;255;0m"
			} else if 100 <= ping && ping < 300 {
				pingcolor = "255;255;0m"
			} else if 300 <= ping {
				pingcolor = "255;0;0m"
			}
			fmt.Printf("ping to \x1b[38;2;102;0;204m%s\x1b[0m is \x1b[38;2;%s%.2fms\x1b[0m, players: \x1b[38;2;255;128;0m%s/%s\x1b[0m, hosting provider: \x1b[38;2;0;255;255m%s\x1b[0m, hosting region: \x1b[38;2;255;87;51m%s, %s\x1b[0m\n", host, pingcolor, ping, onlinePlayers, maxPlayers, hosting, region, country)
			time.Sleep(time.Duration(((1000 - time.Since(start).Milliseconds()) * 1000000)))
		}
	} else if pingtype == "bedrock" {
		for {
			start := time.Now()
			maxPlayers, onlinePlayers, ping := core.BedrockPing(host)
			if ping < 100 {
				pingcolor = "0;255;0m"
			} else if 100 <= ping && ping < 300 {
				pingcolor = "255;255;0m"
			} else if 300 <= ping {
				pingcolor = "255;0;0m"
			}
			fmt.Printf("ping to \x1b[38;2;102;0;204m%s\x1b[0m is \x1b[38;2;%s%.2fms\x1b[0m, players: \x1b[38;2;255;128;0m%s/%s\x1b[0m, hosting provider: \x1b[38;2;0;255;255m%s\x1b[0m, hosting region: \x1b[38;2;255;87;51m%s, %s\x1b[0m\n", host, pingcolor, ping, onlinePlayers, maxPlayers, hosting, region, country)
			time.Sleep(time.Duration(((1000 - time.Since(start).Milliseconds()) * 1000000)))
		}
	}
}
