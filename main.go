package main

import (
	"djcTest/crypto"
	"djcTest/djc"
	"github.com/levigross/grequests"
	"log"
	"os"
	"sync"
)

func main() {
	crypto.Init()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() { //776562198
		c := djc.Client{
			Session: grequests.NewSession(nil),
			Cookies: "djc_appSource=android; djc_appVersion=153; acctype=qc; appid=1101958653; openid=E6D10C8944C1767110F1EB5049C28A46; access_token=E34D0C221AFC38593758FD55D43621A9",
			OpenId:  "E6D10C8944C1767110F1EB5049C28A46",
			Ptk:     "148290189",
		}
		_, _ = c.Run()
		wg.Done()
	}()

	wg.Add(1)
	go func() { //467256306
		c2 := djc.Client{
			Session: grequests.NewSession(nil),
			Cookies: "djc_appSource=android; djc_appVersion=153; acctype=qc; appid=1101958653; openid=3FE2B5B01F48F41C41BD80B3DC9C5797; access_token=8EFEDD28A096396D9B91DF51C55E0B6B",
			OpenId:  "3FE2B5B01F48F41C41BD80B3DC9C5797",
			Ptk:     "492195175",
		}
		_, _ = c2.Run()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		c3 := djc.Client{
			Session: grequests.NewSession(nil),
			Cookies: "djc_appSource=android; djc_appVersion=153; acctype=qc; appid=1101958653; openid=A785E5E7C7E396D487CCEFD4ED92367D; access_token=04C64C1F49DBA23D8FEE69437362C6C9",
			OpenId:  "A785E5E7C7E396D487CCEFD4ED92367D",
			Ptk:     "1447148307",
		}
		_, _ = c3.Run()
		wg.Done()
	}()

	wg.Wait()
	log.Println("所有全部执行完成")

}

func init() {
	c := os.Getenv("BDUSS")
	log.Println("test", c)
}
