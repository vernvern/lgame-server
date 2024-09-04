package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	model "lgame/internal/model"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

var (
    PASSWD = "huwaitegong"
)

// @Summary 用户请求游戏
func ApiGetGame(c *gin.Context) {
	game := model.Game{}
	model.DB.Model(&model.Game{}).Where("status = 0").Order("id desc").First(&game)

	resp := Resp{Code: 0}

	if game.ID == 0 {
		resp.Code = -1
		resp.Msg = "waitting start"
	} else {
		data := ApiGetGameResponse{}
		data.GameId = game.ID
		data.Duration = game.Duration
		resp.Data = data
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary 上传游戏结果
func ApiUploadResult(c *gin.Context) {
	params := ApiUploadResultRequest{}
	resp := Resp{Code: 0}
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		resp.Code = -1
		resp.Msg = "bind data error"
		slog.Error("resp.Msg", "error", err)
		c.JSON(http.StatusOK, resp)
		return
	}

	top := model.Top{}
	top.Userid = params.Userid
	top.GameId = params.GameId
	top.Times = params.Times

	err := model.DB.Save(&top).Error
	if err != nil {
		resp.Code = -232
		resp.Msg = "db error"
		slog.Error(resp.Msg, "error", err)
	} else {
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary 查看排名
func ApiTop(c *gin.Context) {
	game := model.Game{}
	model.DB.Order("id desc").First(&game)

	tops := []model.Top{}
	model.DB.Where("game_id = ?", game.ID).Order("times desc").Limit(100).Find(&tops)

	resp := Resp{Code: 0}
	data := ApiTopResponse{}
	data.Users = []ApiTopResponseTop{}
	for i := range tops {
		data.Users = append(data.Users, ApiTopResponseTop{
			Times:  tops[i].Times,
			Userid: tops[i].Userid,
		})
	}
    resp.Data = data

	c.JSON(http.StatusOK, resp)
}

// @Summary 登录
func ApiLogin(c *gin.Context) {

}

// @Summary 创建游戏
func ApiAddGame(c *gin.Context) {
	params := ApiGetGameRequest{}
	resp := Resp{Code: 0}
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		resp.Code = -1
		resp.Msg = "bind data error"
		slog.Error("resp.Msg", "error", err)
		c.JSON(http.StatusOK, resp)
		return
	}

	if params.Passwd != PASSWD{
		resp.Code = -2
		resp.Msg = "密码错误"
	}

	game := model.Game{}
	game.Countdown = params.Countdown
	game.Duration = params.Duration
	game.Status = 0

	model.DB.Save(&game)

	c.JSON(http.StatusOK, resp)

}

// @Summary 游戏结束
func ApiStopGame(c *gin.Context) {
	params := ApiStopGameRequest{}
	resp := Resp{Code: 0}
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		resp.Code = -1
		resp.Msg = "bind data error"
		slog.Error("resp.Msg", "error", err)
		c.JSON(http.StatusOK, resp)
		return
	}

	if params.Passwd != PASSWD {
		resp.Code = -2
		resp.Msg = "密码错误"
	}

    model.DB.Model(&model.Game{}).Where("status = ?", 0).Updates(&model.Game{Status: -1})

	c.JSON(http.StatusOK, resp)


}
