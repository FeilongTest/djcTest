package djc

type SignRules struct {
	Data []struct {
		IFlowId int    `json:"iFlowId"`
		IDays   string `json:"iDays"`
		IAmount string `json:"iAmount"`
		IRuleId int    `json:"iRuleId"`
		ICanUse int    `json:"iCanUse"`
	} `json:"data"`
	EventId string `json:"event_id"`
	Msg     string `json:"msg"`
	Ret     string `json:"ret"`
	Span    string `json:"span"`
	TraceId string `json:"trace_id"`
	Ts      string `json:"ts"`
}

type TaskList struct {
	Ret  string `json:"ret"`
	Msg  string `json:"msg"`
	IRet string `json:"iRet"`
	SMsg string `json:"sMsg"`
	Data struct {
		List struct {
			Day []struct {
				IruleId       interface{} `json:"iruleId"`
				IladderStatus string      `json:"iladderStatus"`
				ISort         string      `json:"iSort"`
				Sladderup     string      `json:"sladderup"`
				Sladderdown   string      `json:"sladderdown"`
				ItopType      string      `json:"itopType"`
				STask         string      `json:"sTask"`
				STaskDesc     string      `json:"sTaskDesc"`
				STaskDirec    string      `json:"sTaskDirec"`
				SPic          string      `json:"sPic"`
				SBtnDesc      string      `json:"sBtnDesc"`
				SBtnType      string      `json:"sBtnType"`
				SBtnHref      string      `json:"sBtnHref"`
				SReward       []struct {
					RewardId   string `json:"rewardId"`
					RewardName string `json:"rewardName"`
					RewardIcon string `json:"rewardIcon"`
					Num        string `json:"num"`
					Giftid     int    `json:"giftid"`
				} `json:"sReward"`
				ICompleteNum string      `json:"iCompleteNum"`
				Iroleinfo    string      `json:"iroleinfo"`
				DtStart      string      `json:"dtStart"`
				DtEnd        string      `json:"dtEnd"`
				ITaskType    string      `json:"iTaskType"`
				SExt3        string      `json:"sExt3"`
				IReceive     interface{} `json:"iReceive"`
				DtComplete   string      `json:"dtComplete"`
				DtReceive    string      `json:"dtReceive"`
				DtDeadline   string      `json:"dtDeadline"`
				DtCommit     string      `json:"dtCommit"`
				IsNew        int         `json:"isNew"`
				TaskStatus   int         `json:"task_status"`
			} `json:"day"`
			LimitTime []struct {
				IruleId       string      `json:"iruleId"`
				IladderStatus string      `json:"iladderStatus"`
				ISort         string      `json:"iSort"`
				IStatus       interface{} `json:"iStatus"`
				Sladderup     string      `json:"sladderup"`
				Sladderdown   string      `json:"sladderdown"`
				ItopType      string      `json:"itopType"`
				STask         string      `json:"sTask"`
				STaskDesc     string      `json:"sTaskDesc"`
				STaskDirec    string      `json:"sTaskDirec"`
				SPic          string      `json:"sPic"`
				SBtnDesc      string      `json:"sBtnDesc"`
				SBtnType      string      `json:"sBtnType"`
				SBtnHref      string      `json:"sBtnHref"`
				SReward       []struct {
					RewardId   string `json:"rewardId"`
					RewardName string `json:"rewardName"`
					RewardIcon string `json:"rewardIcon"`
					Num        string `json:"num"`
					Giftid     int    `json:"giftid"`
				} `json:"sReward"`
				ICompleteNum string `json:"iCompleteNum"`
				Iroleinfo    string `json:"iroleinfo"`
				DtStart      string `json:"dtStart"`
				DtEnd        string `json:"dtEnd"`
				ITaskType    string `json:"iTaskType"`
				SExt3        string `json:"sExt3"`
				DtComplete   string `json:"dtComplete"`
				DtReceive    string `json:"dtReceive"`
				DtDeadline   string `json:"dtDeadline"`
				DtCommit     string `json:"dtCommit"`
				IsNew        int    `json:"isNew"`
				TaskStatus   int    `json:"task_status"`
			} `json:"limit_time"`
			Game interface{} `json:"game"`
		} `json:"list"`
		LotteryTask int `json:"lottery_task"`
	} `json:"data"`
	SSerialNum string `json:"sSerialNum"`
	RemoteAddr string `json:"remote_addr"`
	EventId    string `json:"_event_id"`
	ServerTime int    `json:"serverTime"`
	SAmsSerial string `json:"sAmsSerial"`
}
