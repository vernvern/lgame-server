package api

type ApiGetGameResponse struct {
	// game id
	GameId uint

    // 用户
    User string

	// 倒数多少秒
	Countdown int

	// 持续时间
	Duration int

	// 开始时间
	Start int64
}

type ApiUploadResultRequest struct {
    Times  int  `json:"times"`
    Userid string `json:"user"`
    GameId uint  `json:"gameid"`
}


type ApiLoginRequest struct {
    Passwd  string
}

type ApiTopResponse struct {
    Users []ApiTopResponseTop `json:"top"`
}


type ApiTopResponseTop struct {
    Times  int  `json:"times"`
    Userid string `json:"user"`
    GameId uint  `json:"gameid"`
}




type ApiGetGameRequest struct {
    Countdown  int  `json:"countdown"`
    Duration int `json:"duration"`
    Passwd string `json:"passwd"`
}

type ApiStopGameRequest struct {
    Passwd string `json:"passwd"`
}



