package djc

import (
	"github.com/levigross/grequests"
)

type Client struct {
	Session   *grequests.Session
	UserAgent string
	Cookies   string
	Ptk       string
	OpenId    string //从Cookie中获取
}

const (
	DeviceId  = "7911dfdf09dd81553c51b4d962f8b5f62caf433ab35a07b6f507e2a13232ebd5"
	UserAgent = "TencentDaojucheng=v4.7.7.0&appSource=android&appVersion=153&ch=10003&sDeviceID=7911dfdf09dd81553c51b4d962f8b5f62caf433ab35a07b6f507e2a13232ebd5&firmwareVersion=9&phoneBrand=Xiaomi&phoneVersion=MI+6&displayMetrics=1080 * 1920&cpu=AArch64 Processor rev 1 (aarch64)&net=wifi&sVersionName=v4.7.7.0&plNo=109"
)
