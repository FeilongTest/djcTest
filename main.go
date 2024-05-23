package main

import (
	"djcTest/crypto"
	"djcTest/djc"
	"encoding/json"
	"github.com/levigross/grequests"
	"log"
	"os"
	"sync"
)

func main() {
	crypto.Init()
	wg := sync.WaitGroup{}

	for _, cookie := range cookies {
		wg.Add(1)
		ck := cookie
		go func() { //776562198
			c := djc.Client{
				Session: grequests.NewSession(nil),
				Cookies: ck.Cookie,
				OpenId:  ck.OpenId,
				Ptk:     ck.Ptk,
			}
			_, _ = c.Run()
			wg.Done()
		}()
	}

	wg.Wait()
	log.Println("所有全部执行完成")

}

type Cookie struct {
	Cookie string `json:"cookie"`
	Ptk    string `json:"ptk"`
	OpenId string `json:"openId"`
}

var cookies []Cookie

func init() {
	c := os.Getenv("COOKIE")
	err := json.Unmarshal([]byte(c), &cookies)
	if err != nil {
		panic(err)
	}
}
