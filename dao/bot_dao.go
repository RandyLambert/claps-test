package dao

import (
	"claps-test/util"
	"claps-test/model"
)

func GetBotById(botId string)(bot *model.Bot,err error){
	bot = &model.Bot{}
	err = util.DB.Debug().Where("id=?",botId).Find(bot).Error
	return
}

//根据projectId获取所有的机器人Id
func ListBotIdsByProjectId(projectId uint32)(botId *[]model.BotId,err error){
	botId = &[]model.BotId{}
	err = util.DB.Debug().Table("bot").Where("project_id=?",projectId).Find(botId).Error
	return
}