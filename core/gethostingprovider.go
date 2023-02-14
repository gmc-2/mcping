package core

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

func GetHostingProvider(host string) (hosting string, country string, region string) {
	jew,err:=net.LookupHost(host)
	if err != nil {
		fmt.Printf("unable to look up host: %s\n",host)
		os.Exit(0)
	}
	ip := jew[0]
	resp,err:= http.Get("http://ip-api.com/json/"+ip)
	body,_ := ioutil.ReadAll(resp.Body)
	hosting = gjson.Get(string(body),"isp").String()
	country = gjson.Get(string(body),"country").String()
	region = gjson.Get(string(body),"regionName").String()
	return
}