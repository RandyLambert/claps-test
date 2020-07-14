package dao

import (
	"claps-test/util"
	"claps-test/model"
)

func GetBot(botId string)(bot *model.Bot,err error){
	bot = &model.Bot{}
	err = util.DB.Debug().Where("id=?",botId).Find(bot).Error
	return
}