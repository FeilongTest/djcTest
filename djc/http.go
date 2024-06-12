package djc

import (
	"djcTest/crypto"
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"strconv"
	"time"
)

func (c *Client) djcGet(service, opType, taskId string, isReceive bool) (result []byte, err error) {
	iRuledId := ""
	if taskId != "" {
		if isReceive {
			iRuledId = "&iruleId=" + taskId
		} else {
			iRuledId = "&ruleid=" + taskId
		}
	}
	timestamp := time.Now().UnixMilli() / 1e6
	sign, err := crypto.GetEncrypt(c.OpenId + DeviceId + strconv.Itoa(int(timestamp)) + "153")
	if err != nil {
		err = errors.New("djcGet计算sign出错:" + err.Error())
		return
	}
	url := fmt.Sprintf("https://djcapp.game.qq.com/daoju/igw/main/?_service=%s&optype=%s&_app_id=1001&output_format=json&iAppId=1001&_app_id=1001%s&sDeviceID=%s&djcRequestId=%s-%d-%s&appVersion=156&p_tk=%s&osVersion=Android-28&ch=10003&sVersionName=v4.8.0.0&appSource=android&sDjcSign=%s",
		service, opType, iRuledId, DeviceId, DeviceId, timestamp, CreateRandomString(3), c.Ptk, sign)
	result, err = c.httpGet(url)
	return
}

func (c *Client) httpGet(url string) (result []byte, err error) {
	resp, err := c.Session.Get(url, &grequests.RequestOptions{
		UserAgent: UserAgent,
		Headers: map[string]string{
			"Cookie": c.Cookies,
		},
	})
	result = resp.Bytes()
	return
}

func (c *Client) httpPost(url string, data map[string]string) (result []byte, err error) {
	resp, err := c.Session.Post(url, &grequests.RequestOptions{
		UserAgent: UserAgent,
		Data:      data,
		Headers: map[string]string{
			"Cookie": c.Cookies,
		},
	})
	result = resp.Bytes()
	return
}
