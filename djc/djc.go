package djc

import (
	"djcTest/crypto"
	"encoding/json"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"strconv"
	"time"
)

func (c *Client) Run() (logInfo string, err error) {

	if time.Now().Day() == 1 {
		c.exchange("2601") //精华
		time.Sleep(1 * time.Second)
		c.exchange("2602") //训练卡
	}

	if time.Now().Weekday() == time.Monday {
		for i := 0; i < 3; i++ {
			c.exchange("2600") //福袋
			time.Sleep(1 * time.Second)
		}
	}

	var balance string
	//获取账户余额
	if balance, err = c.getBalance(); err != nil {
		log.Println("getBalance", err)
		err = errors.New("获取余额失败")
		return
	}
	log.Printf("当前用户余额:%s\n", balance)

	_, _ = c.sign()

	c.getSignRules()

	//执行每日抽奖
	c.lottery()

	result, err := c.djcGet("welink.usertask.swoole", "get_usertask_list", "", false)
	if err != nil {
		return
	}
	var taskList TaskList
	err = json.Unmarshal(result, &taskList)
	if err != nil {
		err = errors.New("获取任务列表解析json失败:" + err.Error())
		log.Println(err)
		return
	}

	for _, day := range taskList.Data.List.Day {
		if day.SBtnDesc == "去完成" {
			//尝试完成打开
			switch day.STask {
			case "点击打卡3个活动":
				for i := 0; i < 3; i++ {
					c.doTodayTask("app.task.report", "activity_card")
				}
				c.receiveTask(fmt.Sprintf("%v", day.IruleId))
				break
			case "打卡活动中心":
				c.doTodayTask("app.task.report", "activity_center")
				c.receiveTask(fmt.Sprintf("%v", day.IruleId))
				break
			default:
				log.Printf("已跳过%s暂时无法自动完成", day.STask)
				break
			}
		} else if day.SBtnDesc == "领取奖励" {
			//领取兑换成功后的任务
			c.receiveTask(fmt.Sprintf("%v", day.IruleId))
		} else {
			log.Printf("%s已完成\n", day.STask)
		}
	}

	log.Println("=============尝试完成限时任务=============")
	for _, limit := range taskList.Data.List.LimitTime {
		if fmt.Sprintf("%v", limit.IStatus) == "0" {
			c.doLimitTask(limit.IruleId)
		} else {
			log.Printf("限时任务%s已完成\n", limit.STask)
		}

	}

	log.Println("=============尝试完成宝箱任务=============")

	result, err = c.djcGet("welink.usertask.swoole", "get_usertask_list", "", false)
	if err != nil {
		return
	}

	if jsoniter.Get(result, "data", "chest_list", "100001", "iCurrentNum").ToInt() >= 40 {
		if jsoniter.Get(result, "data", "chest_list", "100001", "iReceive").ToString() == "1" {
			log.Println("今日活跃度银宝箱已经领取过了")
		} else {
			log.Println("已满足领取活跃度银宝箱条件，尝试领取银宝箱")
			c.receiveTask("100001")
		}
	}

	if jsoniter.Get(result, "data", "chest_list", "100002", "iCurrentNum").ToInt() >= 70 {
		if jsoniter.Get(result, "data", "chest_list", "100002", "iReceive").ToString() == "1" {
			log.Println("今日活跃度金宝箱已经领取过了")
		} else {
			log.Println("已满足领取活跃度金宝箱条件，尝试领取银宝箱")
			c.receiveTask("100002")
		}
	}

	return
}

func (c *Client) getBalance() (balance string, err error) {
	timestamp := time.Now().UnixNano() / 1e6
	sign, err := crypto.GetEncrypt(fmt.Sprintf("%s+%s+%v+153", c.OpenId, DeviceId, timestamp))
	url := fmt.Sprintf("https://djcapp.game.qq.com/daoju/igw/main/?_service=app.bean.balance&iAppId=1001&_app_id=1001&sDeviceID=%s&djcRequestId=%s-%d-%s&appVersion=156&p_tk=%s&osVersion=Android-28&ch=10003&sVersionName=v4.8.0.0&appSource=android&sDjcSign=%s",
		DeviceId, DeviceId, timestamp, CreateRandomString(3), c.Ptk, sign)
	result, err := c.httpGet(url)
	balance = jsoniter.Get(result, "data", "balance").ToString()
	return
}

// 签到
func (c *Client) sign() (logInfo string, err error) {
	timestamp := time.Now().UnixNano() / 1e6
	sign, err := crypto.GetEncrypt(fmt.Sprintf("%s+%s+%v+153", c.OpenId, DeviceId, timestamp))
	url1 := "https://comm.ams.game.qq.com/ams/ame/amesvr?ameVersion=0.3&sServiceType=dj&iActivityId=11117&sServiceDepartment=djc&set_info=newterminals&&appSource=android&appVersion=156&ch=10003&sDeviceID=" + DeviceId + "&osVersion=Android-28&p_tk=" + c.Ptk + "&sVersionName=v4.8.0.0"
	result, err := c.httpPost(url1, map[string]string{
		"djcRequestId":       DeviceId + "-1711712807701-" + CreateRandomString(3),
		"appVersion":         "153",
		"sign_version":       "1.0",
		"ch":                 "10003",
		"iActivityId":        "11117",
		"sDjcSign":           sign,
		"sDeviceID":          DeviceId,
		"p_tk":               c.Ptk,
		"month":              time.Now().Format("200601"),
		"osVersion":          "Android-28",
		"iFlowId":            "96939",
		"sVersionName":       "v4.7.7.0",
		"sServiceDepartment": "djc",
		"sServiceType":       "dj",
		"appSource":          "android",
		"g_tk":               "1842395457", //逆向分析得到所有用到的g_tk均为计算后的固定值1842395457
	})
	log.Println("签到信息:" + jsoniter.Get(result, "modRet", "sMsg").ToString())
	if jsoniter.Get(result, "modRet", "sMsg").ToString() != "" {
		_ = c.getSignAmount("324410")
	}
	return
}

// 获取签到奖励
func (c *Client) getSignAmount(iFlowId string) (err error) {
	timestamp := time.Now().UnixNano() / 1e6
	sign, err := crypto.GetEncrypt(fmt.Sprintf("%s+%s+%v+153", c.OpenId, DeviceId, timestamp))
	url1 := "https://comm.ams.game.qq.com/ams/ame/amesvr?ameVersion=0.3&sServiceType=dj&iActivityId=11117&sServiceDepartment=djc&set_info=newterminals&w_ver=36&w_id=45&appSource=android&appVersion=156&ch=10003&sDeviceID=" + DeviceId + "&osVersion=Android-28&p_tk=" + c.Ptk + "&sVersionName=v4.8.0.0"
	result, err := c.httpPost(url1, map[string]string{
		"djcRequestId":       DeviceId + "-1711712807701-" + CreateRandomString(3),
		"appVersion":         "153",
		"sign_version":       "1.0",
		"ch":                 "10003",
		"iActivityId":        "11117",
		"sDjcSign":           sign,
		"sDeviceID":          DeviceId,
		"p_tk":               c.Ptk,
		"osVersion":          "Android-28",
		"iFlowId":            iFlowId,
		"sVersionName":       "v4.7.7.0",
		"sServiceDepartment": "djc",
		"sServiceType":       "dj",
		"appSource":          "android",
		"g_tk":               "1842395457", //逆向分析得到所有用到的g_tk均为计算后的固定值1842395457
	})
	log.Println("获取签到奖励信息:" + jsoniter.Get(result, "modRet", "sMsg").ToString())
	return
}

// 获取签到总天数
func (c *Client) getSignTotalDays() (totalDays int) {
	timestamp := time.Now().UnixNano() / 1e6
	sign, _ := crypto.GetEncrypt(fmt.Sprintf("%s+%s+%v+153", c.OpenId, DeviceId, timestamp))
	url1 := "https://comm.ams.game.qq.com/ams/ame/amesvr?ameVersion=0.3&sServiceType=dj&iActivityId=11117&sServiceDepartment=djc&set_info=newterminals&w_ver=36&w_id=45&appSource=android&appVersion=156&ch=10003&sDeviceID=" + DeviceId + "&osVersion=Android-28&p_tk=" + c.Ptk + "&sVersionName=v4.8.0.0"
	result, _ := c.httpPost(url1, map[string]string{
		"djcRequestId":       DeviceId + "-1711712807701-" + CreateRandomString(3),
		"appVersion":         "153",
		"sign_version":       "1.0",
		"ch":                 "10003",
		"iActivityId":        "11117",
		"sDjcSign":           sign,
		"sDeviceID":          DeviceId,
		"p_tk":               c.Ptk,
		"month":              time.Now().Format("200601"),
		"osVersion":          "Android-28",
		"iFlowId":            "96938",
		"sVersionName":       "v4.7.7.0",
		"sServiceDepartment": "djc",
		"sServiceType":       "dj",
		"appSource":          "android",
		"g_tk":               "1842395457", //逆向分析得到所有用到的g_tk均为计算后的固定值1842395457
	})
	totalDays = jsoniter.Get(result, "modRet", "data").Size()
	return
}

// 获取签到规则
func (c *Client) getSignRules() {
	timestamp := time.Now().UnixNano() / 1e6
	sign, _ := crypto.GetEncrypt(fmt.Sprintf("%s+%s+%v+153", c.OpenId, DeviceId, timestamp))
	url1 := "https://djcapp.game.qq.com/daoju/igw/main/?_service=app.sign.rules&output_format=json&iAppId=1001&_app_id=1001&w_ver=36&w_id=45&sDeviceID=" + DeviceId + "&djcRequestId=" + DeviceId + "-" + strconv.Itoa(int(timestamp)) + "-" + CreateRandomString(3) + "&appVersion=156&p_tk=" + c.Ptk + "&osVersion=Android-28&ch=10003&sVersionName=v4.8.0.0&appSource=android&sDjcSign=" + sign
	result, err := c.httpGet(url1)
	if err != nil {
		return
	}
	var signRules SignRules
	err = jsoniter.Unmarshal(result, &signRules)
	if err != nil {
		err = errors.New("获取签到规则解析Json失败")
		return
	}
	//获取总天数
	totalDays := c.getSignTotalDays()
	log.Println("累计签到总天数:" + strconv.Itoa(totalDays))
	for _, datum := range signRules.Data {
		if datum.ICanUse == 0 {
			log.Printf("累计签到%s天的奖励已领取过", datum.IDays)
		} else {
			iDays, _ := strconv.Atoi(datum.IDays)
			if totalDays >= iDays {
				// 领取签到奖励
				_ = c.getSignAmount(strconv.Itoa(datum.IFlowId))
			} else {
				log.Printf("累计签到%s天的奖励不满足领取条件", datum.IDays)
			}
		}
	}
}

// 做今日任务
func (c *Client) doTodayTask(service, taskType string) {
	timestamp := time.Now().UnixNano() / 1e6
	sign, _ := crypto.GetEncrypt(fmt.Sprintf("%s+%s+%v+153", c.OpenId, DeviceId, timestamp))
	result, err := c.httpGet(fmt.Sprintf("https://djcapp.game.qq.com/daoju/igw/main/?_service=%s&task_type=%s&_app_id=1001&output_format=json&iAppId=1001&_app_id=1001&sDeviceID=%s&djcRequestId=%s-%d-%s&appVersion=156&p_tk=%s&osVersion=Android-28&ch=10003&sVersionName=v4.8.0.0&appSource=android&sDjcSign=%s",
		service, taskType, DeviceId, DeviceId, timestamp, CreateRandomString(3), c.Ptk, sign))
	if err != nil {
		log.Println("执行今日任务失败", err)
		return
	}
	log.Println(string(result))
	return
}

func (c *Client) doActivityCenter(taskId string) {
	if result, err := c.djcGet("welink.usertask.swoole", "report_usertask_rushtime", taskId, false); err != nil {
		if jsoniter.Get(result, "ret").ToString() == "0" {
			log.Printf("正在提交taskId=%s任务...\n", taskId)
		} else {
			log.Println("提交活动中心任务失败", err)
		}
	}

	return
}

func (c *Client) doTask(taskId string) {
	if _, err := c.djcGet("welink.usertask.swoole", "report_usertask_rushtime", taskId, false); err != nil {

		log.Printf("尝试提交taskId=%s任务...\n", taskId)
		/**
		if jsoniter.Get(result, "ret").ToString() == "0" {
			log.Printf("正在提交taskId=%s任务...\n", taskId)
		} else {
			log.Println("提交活动中心任务失败", err)
		}

		*/

	}

	return
}

// 领取任务
func (c *Client) receiveTask(taskId string) {
	if result, err := c.djcGet("welink.usertask.swoole", "receive_usertask", taskId, true); err == nil {
		log.Printf("taskId=%s任务领取返回信息:%s\n", taskId, jsoniter.Get(result, "sMsg").ToString())
	}
	return
}

// 做限时任务
func (c *Client) doLimitTask(taskId string) {
	c.doTask(taskId)
	time.Sleep(1 * time.Second)
	c.receiveTask(taskId)
	return
}

// 兑换
func (c *Client) exchange(iGoodsSeqId string) {
	timestamp := time.Now().UnixNano() / 1e6
	sign, _ := crypto.GetEncrypt(fmt.Sprintf("%s+%s+%v+153", c.OpenId, DeviceId, timestamp))
	url1 := "https://djcapp.game.qq.com/daoju/igw/main/?_service=app.role.bind_list&iAppId=1001&_app_id=1001&_biz_code=nba2k2&type=0&sDeviceID=" + DeviceId + "&djcRequestId=" + DeviceId + "-" + strconv.Itoa(int(timestamp)) + "-" + CreateRandomString(3) + "&appVersion=156&p_tk=" + c.Ptk + "&osVersion=Android-28&ch=10003&sVersionName=v4.8.0.0&appSource=android&sDjcSign=" + sign
	result, err := c.httpGet(url1)
	if err != nil {
		log.Println("获取角色信息失败")
		return
	}
	roleId := jsoniter.Get(result, "data", 0, "sRoleInfo", "roleCode").ToString()
	roleName := jsoniter.Get(result, "data", 0, "sRoleInfo", "roleName").ToString()

	//兑换
	url1 = "https://djcapp.game.qq.com/daoju/igw/main/?_service=buy.plug.swoole.judou&iAppId=1001&_app_id=1003&_output_fmt=1&_plug_id=9800&_from=app&iGoodsSeqId=" + iGoodsSeqId + "&iActionId=9002&iActionType=26&_biz_code=nba2k2&biz=nba2k2&iZone=30&lRoleId=" + roleId + "&rolename=" + roleName + "&p_tk=" + c.Ptk + "&_cs=2&sDeviceID=" + DeviceId + "&djcRequestId=" + DeviceId + "-" + strconv.Itoa(int(timestamp)) + "-" + CreateRandomString(3) + "&appVersion=156&p_tk=" + c.Ptk + "&osVersion=Android-28&ch=10003&sVersionName=v4.8.0.0&appSource=android&sDjcSign=" + sign
	result, err = c.httpGet(url1)
	if err != nil {
		log.Println("兑换道具失败")
		return
	}
	log.Println("兑换道具成功", string(result))
}

// 每日抽奖
func (c *Client) lottery() {
	result, err := c.djcGet("welink.usertask.swoole", "lottery_usertask", "", false)
	if err != nil {
		log.Println("执行每日抽奖失败")
		return
	}
	ret := jsoniter.Get(result, "ret").ToInt()
	if ret == -1 {
		log.Println(jsoniter.Get(result, "sMsg").ToString())
	} else if ret == 0 {
		sTask := jsoniter.Get(result, "data", "data", "sTask").ToString()
		lUin := jsoniter.Get(result, "data", "data", "lUin").ToString()
		if sTask == "幸运任务" {
			//领取
			iRuleId := jsoniter.Get(result, "data", "data", "iruleId").ToString()
			c.receiveTask(iRuleId)
		} else {
			log.Printf("%s今天抽到的任务为:%s,自动跳过\n", lUin, sTask)
		}
	} else {
		log.Println("每日抽奖返回结果异常,请检查", string(result))
	}
}
