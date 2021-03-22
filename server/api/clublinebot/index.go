package clublinebot

import (
	clublinebotLogic "heroku-line-bot/logic/clublinebot"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// Index Line Bot
// @Tags Line Bot
// @Summary Line Bot Handler
// @Accept json
// @Produce  json
// @Param param body reqs.Index true "參數"
// @Success 200
// @Router /club-line-bot [post]
func Index(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	if err := clublinebotLogic.Bot.Handle(string(jsonData)); err != nil {
		return
	}
	c.JSON(200, nil)
}
