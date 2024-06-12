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
	DeviceId  = "4fedc7908cdc0454a5732dfbd105349ba9aeda18365b9e9470067b83b73372f7"
	UserAgent = "TencentDaojucheng=v4.8.0.0&appSource=android&appVersion=156&ch=10003&sDeviceID=4fedc7908cdc0454a5732dfbd105349ba9aeda18365b9e9470067b83b73372f7&firmwareVersion=9&phoneBrand=xiaomi&phoneVersion=MI+6X&displayMetrics=1080 * 2030&cpu=AArch64 Processor rev 2 (aarch64)&net=wifi&sVersionName=v4.8.0.0&plNo=253"
)
